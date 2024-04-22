//go:build x86_64

package ps30m

import (
	"time"

	"github.com/merliot/device/modbus"
	"github.com/tarm/serial"
)

var dev = "/dev/ttyUSB0"

type transport struct {
	*serial.Port
}

func newTransport() *transport {
	port, _ := serial.OpenPort(&serial.Config{
		Name:        dev,
		Baud:        9600,
		StopBits:    2,
		Parity:      serial.ParityNone,
		ReadTimeout: time.Second,
	})
	return &transport{port}
}

func (t *transport) Read(buf []byte) (int, error) {
	if t.Port == nil {
		return 0, modbus.ErrPortNotOpen
	}
	return t.Port.Read(buf)
}

func (t *transport) Write(buf []byte) (int, error) {
	if t.Port == nil {
		return 0, modbus.ErrPortNotOpen
	}
	return t.Port.Write(buf)
}
