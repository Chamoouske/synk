package domain

type Config struct {
	Service struct {
		Name   string `ini:"name"`
		Type   string `ini:"type"`
		Port   int    `ini:"port"`
		Domain string `ini:"domain"`
	} `ini:"service"`
}
