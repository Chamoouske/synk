package domain

type Service interface {
	Start() error
	Stop() error
	GetPort() int
	AddDeviceToConnect(deviceID string) error
}
