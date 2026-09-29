package schsmkms

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	openapi "github.com/CardContact/sc-hsm-cloud-service-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.step.sm/crypto/kms/apiv1"
)

func generateSelfSignedCertAndKey(t *testing.T) (certPEM, keyPEM []byte) {
	t.Helper()
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)
	require.NoError(t, err)

	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	privDER, err := x509.MarshalECPrivateKey(priv)
	require.NoError(t, err)
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privDER})

	return certPEM, keyPEM
}

func encodePublicKeyToBase64(t *testing.T, pub crypto.PublicKey) string {
	t.Helper()
	der, err := x509.MarshalPKIXPublicKey(pub)
	require.NoError(t, err)
	return base64.StdEncoding.EncodeToString(der)
}

func TestRegistration(t *testing.T) {
	fn, ok := apiv1.LoadKeyManagerNewFunc(apiv1.Type("schsmkms"))
	require.True(t, ok)
	require.NotNil(t, fn)

	// Missing uri/url should fail
	_, err := fn(context.Background(), apiv1.Options{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "'uri' must contain the url of the sc-hsm-cloud-service")

	// Valid uri should succeed
	km, err := fn(context.Background(), apiv1.Options{
		URI: "schsmkms:?url=http://localhost:8443&hsm-id=test-hsm",
	})
	require.NoError(t, err)
	require.NotNil(t, km)
	assert.IsType(t, &SmartCardHSMKMS{}, km)
}

func TestNew(t *testing.T) {
	tmpDir := t.TempDir()

	caCertPEM, _ := generateSelfSignedCertAndKey(t)
	caCertPath := filepath.Join(tmpDir, "ca.crt")
	require.NoError(t, os.WriteFile(caCertPath, caCertPEM, 0600))

	clientCertPEM, clientKeyPEM := generateSelfSignedCertAndKey(t)
	clientCertPath := filepath.Join(tmpDir, "client.crt")
	clientKeyPath := filepath.Join(tmpDir, "client.key")
	require.NoError(t, os.WriteFile(clientCertPath, clientCertPEM, 0600))
	require.NoError(t, os.WriteFile(clientKeyPath, clientKeyPEM, 0600))

	invalidCertPath := filepath.Join(tmpDir, "invalid.crt")
	require.NoError(t, os.WriteFile(invalidCertPath, []byte("not-a-pem-cert"), 0600))

	tests := []struct {
		name    string
		opts    apiv1.Options
		wantErr bool
		errSub  string
	}{
		{
			name: "ok basic",
			opts: apiv1.Options{
				URI: "schsmkms:?url=https://localhost:8443/se/api&hsm-id=hsm1",
			},
			wantErr: false,
		},
		{
			name: "ok with ca cert and client cert/key",
			opts: apiv1.Options{
				URI: "schsmkms:?url=" + url.QueryEscape("https://localhost:8443/se/api") +
					"&hsm-id=hsm1" +
					"&ca-cert=" + url.QueryEscape(caCertPath) +
					"&cert=" + url.QueryEscape(clientCertPath) +
					"&key=" + url.QueryEscape(clientKeyPath),
			},
			wantErr: false,
		},
		{
			name: "fail missing service url",
			opts: apiv1.Options{
				URI: "schsmkms:?hsm-id=hsm1",
			},
			wantErr: true,
			errSub:  "'uri' must contain the url of the sc-hsm-cloud-service",
		},
		{
			name: "fail empty uri",
			opts: apiv1.Options{
				URI: "",
			},
			wantErr: true,
			errSub:  "'uri' must contain the url of the sc-hsm-cloud-service",
		},
		{
			name: "fail ca cert not found",
			opts: apiv1.Options{
				URI: "schsmkms:?url=https://localhost:8443&ca-cert=" + url.QueryEscape(filepath.Join(tmpDir, "nonexistent.crt")),
			},
			wantErr: true,
			errSub:  "Loading ca certificate failed",
		},
		{
			name: "fail client cert load error",
			opts: apiv1.Options{
				URI: "schsmkms:?url=https://localhost:8443" +
					"&cert=" + url.QueryEscape(invalidCertPath) +
					"&key=" + url.QueryEscape(clientKeyPath),
			},
			wantErr: true,
			errSub:  "Loading client key pair failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kms, err := New(context.Background(), tt.opts)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errSub != "" {
					assert.Contains(t, err.Error(), tt.errSub)
				}
				assert.Nil(t, kms)
			} else {
				require.NoError(t, err)
				require.NotNil(t, kms)
				assert.Equal(t, tt.opts.URI != "" && kms.hsmId == "hsm1", true)
			}
		})
	}
}

func TestSmartCardHSMKMS_Close(t *testing.T) {
	kms, err := New(context.Background(), apiv1.Options{
		URI: "schsmkms:?url=http://localhost:8443",
	})
	require.NoError(t, err)
	assert.NoError(t, kms.Close())
}

func TestSmartCardHSMKMS_CreateKey(t *testing.T) {
	kms, err := New(context.Background(), apiv1.Options{
		URI: "schsmkms:?url=http://localhost:8443",
	})
	require.NoError(t, err)

	resp, err := kms.CreateKey(&apiv1.CreateKeyRequest{
		Name: "test-key",
	})
	require.NoError(t, err)
	assert.Equal(t, &apiv1.CreateKeyResponse{}, resp)
}

func TestSmartCardHSMKMS_GetPublicKey(t *testing.T) {
	ecKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	ecPubB64 := encodePublicKeyToBase64(t, &ecKey.PublicKey)

	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	rsaPubB64 := encodePublicKeyToBase64(t, &rsaKey.PublicKey)

	invalidBase64 := "not-base-64-%%%"
	invalidDER := base64.StdEncoding.EncodeToString([]byte("invalid der bytes"))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/hsms/valid-hsm/keys":
			keys := []openapi.Key{
				{
					Id:     "key-id-1",
					Label:  "key-label-1",
					Type:   openapi.EC,
					Pubkey: &ecPubB64,
				},
				{
					Id:     "key-id-2",
					Label:  "key-label-2",
					Type:   openapi.RSA,
					Pubkey: &rsaPubB64,
				},
				{
					Id:     "key-id-invalid-b64",
					Label:  "invalid-b64",
					Type:   openapi.EC,
					Pubkey: &invalidBase64,
				},
				{
					Id:     "key-id-invalid-der",
					Label:  "invalid-der",
					Type:   openapi.EC,
					Pubkey: &invalidDER,
				},
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(keys)
		case "/hsms/not-found-hsm/keys":
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message": "HSM not found"}`))
		case "/hsms/error-hsm/keys":
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message": "Internal Server Error"}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	tests := []struct {
		name    string
		hsmID   string
		keyName string
		wantPub crypto.PublicKey
		wantErr bool
		errSub  string
	}{
		{
			name:    "ok match by label (EC)",
			hsmID:   "valid-hsm",
			keyName: "key-label-1",
			wantPub: &ecKey.PublicKey,
			wantErr: false,
		},
		{
			name:    "ok match by id (RSA)",
			hsmID:   "valid-hsm",
			keyName: "key-id-2",
			wantPub: &rsaKey.PublicKey,
			wantErr: false,
		},
		{
			name:    "fail hsm not found (404)",
			hsmID:   "not-found-hsm",
			keyName: "key-label-1",
			wantErr: true,
			errSub:  "schsmkms: HSM not-found-hsm not found",
		},
		{
			name:    "fail retrieving key list error (500)",
			hsmID:   "error-hsm",
			keyName: "key-label-1",
			wantErr: true,
			errSub:  "schsmkms: Retrieving a key list failed",
		},
		{
			name:    "fail key not found",
			hsmID:   "valid-hsm",
			keyName: "unknown-key",
			wantErr: true,
			errSub:  "schsmkms: no key found for name unknown-key",
		},
		{
			name:    "fail decoding public key",
			hsmID:   "valid-hsm",
			keyName: "invalid-b64",
			wantErr: true,
			errSub:  "schsmkms: Decoding public key failed",
		},
		{
			name:    "fail parsing public key",
			hsmID:   "valid-hsm",
			keyName: "invalid-der",
			wantErr: true,
			errSub:  "schsmkms: Parsing public key failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kms, err := New(context.Background(), apiv1.Options{
				URI: "schsmkms:?url=" + url.QueryEscape(server.URL) + "&hsm-id=" + tt.hsmID,
			})
			require.NoError(t, err)

			pub, err := kms.GetPublicKey(&apiv1.GetPublicKeyRequest{
				Name: tt.keyName,
			})
			if tt.wantErr {
				require.Error(t, err)
				if tt.errSub != "" {
					assert.Contains(t, err.Error(), tt.errSub)
				}
				assert.Nil(t, pub)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantPub, pub)
			}
		})
	}
}

func TestSmartCardHSMKMS_CreateSigner(t *testing.T) {
	ecKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	ecPubB64 := encodePublicKeyToBase64(t, &ecKey.PublicKey)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/keys/valid-key":
			key := openapi.Key{
				Id:     "valid-key",
				Label:  "my-key",
				Type:   openapi.EC,
				Pubkey: &ecPubB64,
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(key)
		case "/keys/not-found-key":
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message": "key not found"}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	kms, err := New(context.Background(), apiv1.Options{
		URI: "schsmkms:?url=" + url.QueryEscape(server.URL),
	})
	require.NoError(t, err)

	t.Run("fail empty signing key", func(t *testing.T) {
		signer, err := kms.CreateSigner(&apiv1.CreateSignerRequest{
			SigningKey: "",
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "schsmkms: signingKey must not be empty")
		assert.Nil(t, signer)
	})

	t.Run("ok create signer", func(t *testing.T) {
		signer, err := kms.CreateSigner(&apiv1.CreateSignerRequest{
			SigningKey: "valid-key",
		})
		require.NoError(t, err)
		require.NotNil(t, signer)
		assert.Equal(t, &ecKey.PublicKey, signer.Public())
	})

	t.Run("fail create signer key not found", func(t *testing.T) {
		signer, err := kms.CreateSigner(&apiv1.CreateSignerRequest{
			SigningKey: "not-found-key",
		})
		require.Error(t, err)
		assert.Nil(t, signer)
	})
}
