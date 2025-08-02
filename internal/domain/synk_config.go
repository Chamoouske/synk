package domain

type SynkConfig struct {
	Watch       string   `json:"watch"`
	Connections []Device `json:"connections"`
}
