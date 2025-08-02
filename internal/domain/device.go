package domain

type Device struct {
	ID          string   `json:"id"`
	PublicKey   string   `json:"public_key"`
	PrivateKey  string   `json:"private_key"`
	Connections []string `json:"connections"`
}
