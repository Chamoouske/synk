package domain

type Notifyer interface {
	Notify(message string) error
}
