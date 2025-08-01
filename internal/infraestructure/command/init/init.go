package init

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"synk/internal/domain"
)

const CommandName = "init"

type InitCommand struct {
	notifyer domain.Notifyer
}

func NewInitCommand(notifyer domain.Notifyer) *InitCommand {
	return &InitCommand{notifyer: notifyer}
}

func (c *InitCommand) Execute(args []string) error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		c.notifyer.Notify("Erro ao gerar a chave privada ECDSA: " + err.Error())
		os.Exit(1)
	}

	publicKey := &privateKey.PublicKey
	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)

	if err != nil {
		c.notifyer.Notify("Erro ao serializar a chave privada: " + err.Error())
		os.Exit(1)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		c.notifyer.Notify("Erro ao serializar a chave pública: " + err.Error())
		os.Exit(1)
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	c.notifyer.Notify("Chave Privada: " + string(privateKeyPEM))
	c.notifyer.Notify("Chave Pública: " + string(publicKeyPEM))

	err = os.WriteFile("ecdsa_private.pem", privateKeyPEM, 0600)
	if err != nil {
		c.notifyer.Notify("Erro ao salvar a chave privada ECDSA: " + err.Error())
	}
	err = os.WriteFile("ecdsa_public.pem", publicKeyPEM, 0644)
	if err != nil {
		c.notifyer.Notify("Erro ao salvar a chave pública ECDSA: " + err.Error())
	}
	return nil
}
