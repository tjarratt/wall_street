package echo

type Echoer interface {
	Enable(fileDescriptors []uintptr) error
	Disable(fileDescriptors []uintptr) (int, error)
}

type realEchoer struct{}

func NewEchoer() Echoer {
	return realEchoer{}
}

func (echoer realEchoer) Enable(fileDescriptors []uintptr) error {
	return echoOn(fileDescriptors)
}

func (echoer realEchoer) Disable(fileDescriptors []uintptr) (int, error) {
	return echoOff(fileDescriptors)
}
