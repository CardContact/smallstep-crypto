
package schsmkms

import (
	"context"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/url"
	"net/http"
	"os"
	"go.step.sm/crypto/kms/apiv1"
	"go.step.sm/crypto/kms/schsmkms/client"
)

type SmartCardHSMKMS struct {
	client *client.APIClient
	hsmId string
}

func New(_ context.Context, opts apiv1.Options) (*SmartCardHSMKMS, error) {
	var serviceURL string
	var hsmId string
	var caCertPath string
	var clientCertPath string
	var clientKeyPath string

	if opts.URI != "" {
		u, err := url.Parse(opts.URI)
		if err == nil {
			serviceURL = u.Query().Get("url")
			hsmId = u.Query().Get("hsm-id")
			caCertPath = u.Query().Get("ca-cert")
			clientCertPath = u.Query().Get("cert")
			clientKeyPath = u.Query().Get("key")
		}
	}

	if serviceURL == "" {
		return nil, fmt.Errorf("schsmkms: 'uri' must contain the url of the sc-hsm-cloud-service e.g. '?url=https://localhost:8443/se/api'")
	}

	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations {
		{
			URL: serviceURL,
			Description: "URL of the SmartCard-HSM Cloud Service",
		},
	}

	tlsConfig := &tls.Config {
		MinVersion: tls.VersionTLS13,
	}

	if caCertPath != "" {
		caCert, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, fmt.Errorf("schsmkms: Loading ca certificate failed: %w", err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
	}

	if clientCertPath != "" && clientKeyPath != "" {
		cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
		if err != nil {
			return nil, fmt.Errorf("schsmkms: Loading client key pair failed: %w", err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	transport := &http.Transport {
		TLSClientConfig: tlsConfig,
	}
	httpClient := &http.Client {
		Transport: transport,
	}

	cfg.HTTPClient = httpClient

	apiClient := client.NewAPIClient(cfg)

	return &SmartCardHSMKMS{apiClient, hsmId}, nil
}

func init() {
	apiv1.Register(apiv1.Type("schsmkms"), func(ctx context.Context, opts apiv1.Options) (apiv1.KeyManager, error) {
		return New(ctx, opts)
	})
}

func (k *SmartCardHSMKMS) Close() error {
	return nil
}

func (k *SmartCardHSMKMS) GetPublicKey(req *apiv1.GetPublicKeyRequest) (crypto.PublicKey, error) {
	ctx := context.Background()

	keys, httpRsp, err := k.client.DefaultAPI.GetKeysForHSM(ctx, k.hsmId).Execute()

	if httpRsp.StatusCode == 404 {
		return nil, fmt.Errorf("schsmkms: HSM %v not found", k.hsmId)
	}

	if err != nil {
		return nil, fmt.Errorf("schsmkms: Retrieving a key list failed: %w", err)
	}
	defer httpRsp.Body.Close()

	var requestedKey *client.Key
	for _, key := range keys {
		if key.Label == req.Name || key.Id == req.Name {
			requestedKey = &key
		}
	}

	if requestedKey == nil {
		return nil, fmt.Errorf("schsmkms: no key found for name " + req.Name)
	}

	derBytes, err := base64.StdEncoding.DecodeString(*requestedKey.Pubkey)
	if err != nil {
		return nil, fmt.Errorf("schsmkms: Decoding public key failed: %w", err)
	}

	pubKey, err := x509.ParsePKIXPublicKey(derBytes)

	if err != nil {
		return nil, fmt.Errorf("schsmkms: Parsing public key failed: %w", err)
	}

	return pubKey, nil
}

func (k *SmartCardHSMKMS) CreateKey(req *apiv1.CreateKeyRequest) (*apiv1.CreateKeyResponse, error) {
	fmt.Println("[SmartCardHSMKMS-DEBUG] not implemented")
	return &apiv1.CreateKeyResponse{}, nil
}

func (k *SmartCardHSMKMS) CreateSigner(req *apiv1.CreateSignerRequest) (crypto.Signer, error) {
	return nil, nil
}
