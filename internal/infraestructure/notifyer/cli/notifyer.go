package cli

import "fmt"

type CliNotifyer struct{}

func (c *CliNotifyer) Notify(message string) error {
	fmt.Println(message)
	return nil
}
