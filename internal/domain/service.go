package domain

type Service interface {
	Start() error
	Stop() error
}
