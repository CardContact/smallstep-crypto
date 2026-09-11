package schsmkms

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/CardContact/sc-hsm-cloud-service-go-client"
)

type Signer struct {
	client    *openapi.APIClient
	keyID     string
	publicKey crypto.PublicKey
}

func NewSigner(client *openapi.APIClient, keyID string) (*Signer, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	key, httpRsp, err := client.DefaultAPI.GetKey(ctx, keyID).Execute()

	if httpRsp.StatusCode == 404 {
		return nil, fmt.Errorf("schsmkms: Key %v not found", keyID)
	}

	if err != nil {
		return nil, fmt.Errorf("schsmkms: GetKey failed: %w", err)
	}
	defer httpRsp.Body.Close()

	derBytes, err := base64.StdEncoding.DecodeString(*key.Pubkey)
	if err != nil {
		return nil, fmt.Errorf("schsmkms: Decoding public key failed: %w", err)
	}

	pubKey, err := x509.ParsePKIXPublicKey(derBytes)

	if err != nil {
		return nil, fmt.Errorf("schsmkms: Parsing public key failed: %w", err)
	}

	signer := &Signer{
		client:    client,
		keyID:     keyID,
		publicKey: pubKey,
	}

	return signer, nil
}

func (s *Signer) Public() crypto.PublicKey {
	return s.publicKey
}

func (s *Signer) Sign(_ io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	var algo openapi.KeyAlgorithm
	if _, ok := s.publicKey.(*rsa.PublicKey); ok {
		//if _, ok := opts.(*rsa.PSSOptions); ok {
		//	algo = client.RSA_PSS
		//} else {
		//	algo = client.RSA_PKCS1
		//}
		algo = openapi.RSA_PSS
	} else {
		algo = openapi.ECDSA
		//switch h := opts.HashFunc(); h {
		//case crypto.SHA256:
		//	algo = client.ECDSA_SHA256
		//case crypto.SHA384:
		//	algo = client.ECDSA_SHA384
		//case crypto.SHA512:
		//	algo = client.ECDSA_SHA512
		//default:
		//	return nil, fmt.Errorf("unsupported hash function %v for EC key", h)
		//}
	}

	hash := hex.EncodeToString(digest)
	signatureInput := openapi.NewSignatureInput(hash, algo)
	//fmt.Println("[SmartCardHSMKMS-Signer-DEBUG] hash " + hash)
	//fmt.Println("[SmartCardHSMKMS-Signer-DEBUG] algo " + algo)

	signatureRsp, httpRsp, err := s.client.DefaultAPI.SignHash(ctx, s.keyID).SignatureInput(*signatureInput).Execute()

	if httpRsp.StatusCode != 200 {
		return nil, fmt.Errorf("schsmkms: SignHash failed with status code: %v", httpRsp.StatusCode)
	}
	if err != nil {
		return nil, fmt.Errorf("schsmkms: SignHash failed: %w", err)
	}

	signature, err := hex.DecodeString(signatureRsp.GetSignature())
	return signature, err
}

func defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}
