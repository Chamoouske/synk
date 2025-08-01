package init

import (
	"fmt"
	"os"
	"synk/internal/domain"
	"synk/internal/infraestructure/pem"
)

const CommandName = "init"

type InitCommand struct {
	notifyer domain.Notifyer
}

func NewInitCommand(notifyer domain.Notifyer) *InitCommand {
	return &InitCommand{notifyer: notifyer}
}

func (c *InitCommand) Execute(args []string) error {
	err := c.createKeys()
	if err != nil {
		c.notifyer.Notify("Erro ao criar chaves: " + err.Error())
		os.Exit(1)
	}

	err = c.registerService()
	if err != nil {
		c.notifyer.Notify("Erro ao registrar serviço: " + err.Error())
		os.Exit(1)
	}
	return nil
}

func (c *InitCommand) createKeys() error {
	privateKeyPEM, publicKeyPEM, err := pem.GenerateKeys()
	if err != nil {
		return fmt.Errorf("erro ao gerar chaves: %w", err)
	}

	err = pem.SaveKeys(privateKeyPEM, publicKeyPEM)
	if err != nil {
		return fmt.Errorf("erro ao salvar chaves: %w", err)
	}

	c.notifyer.Notify("Chave Privada: " + string(privateKeyPEM))
	c.notifyer.Notify("Chave Pública: " + string(publicKeyPEM))

	return nil
}

func (c *InitCommand) registerService() error {
	return nil
}

func (c *InitCommand) UnregisterService() error {
	return nil
}
