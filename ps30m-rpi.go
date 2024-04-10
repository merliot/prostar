//go:build rpi

package ps30m

import "errors"

type transport struct {
}

func newTransport() *transport {
	return &transport{}
}

func (t *transport) Write(buf []byte) (n int, err error) {
	return 0, errors.New("Write not implemented")
}

func (t *transport) Read(buf []byte) (n int, err error) {
	return 0, errors.New("Read not implemented")
}
