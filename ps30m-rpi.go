//go:build rpi

package ps30m

import (
	"errors"
	"fmt"
	"time"

	"github.com/goburrow/serial"
)

type transport struct {
	serial.Port
}

var dev = "/dev/ttyUSB0"

func newTransport() *transport {
	port, err := serial.Open(&serial.Config{
		Address:  dev,
		BaudRate: 9600,
		StopBits: 2,
		Parity:   "N",
		Timeout:  time.Second,
	})
	if err != nil {
		fmt.Println("Error opening serial port", dev, err)
		return nil
	}
	return &transport{port}
}

func (t *transport) Read(buf []byte) (n int, err error) {
	if t.Port == nil {
		return 0, errors.New("Port" + dev + "not opened")
	}
	return t.Port.Read(buf)
}

func (t *transport) Write(buf []byte) (n int, err error) {
	if t.Port == nil {
		return 0, errors.New("Port " + dev + "not opened")
	}
	return t.Port.Write(buf)
}
