package pem

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateKeys() ([]byte, []byte, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao gerar a chave privada ECDSA: %w", err)
	}
	publicKey := &privateKey.PublicKey
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao serializar a chave privada: %w", err)
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao serializar a chave pública: %w", err)
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return privateKeyPEM, publicKeyPEM, nil
}

func SaveKeys(privateKeyPEM, publicKeyPEM []byte) error {
	err := os.WriteFile("ecdsa_private.pem", privateKeyPEM, 0600)
	if err != nil {
		return fmt.Errorf("Erro ao salvar a chave privada ECDSA: %w", err)
	}
	err = os.WriteFile("ecdsa_public.pem", publicKeyPEM, 0644)
	if err != nil {
		return fmt.Errorf("Erro ao salvar a chave pública ECDSA: %w", err)
	}
	return nil
}
