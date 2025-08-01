package domain

type Command interface {
	Execute(args []string) error
}
