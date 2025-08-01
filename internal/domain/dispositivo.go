package domain

type Dispositivo struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	IP    string `json:"ip"`
	Porta int    `json:"porta"`
}
