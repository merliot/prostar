//go:build tinygo

package prostar

import (
	"embed"
	"machine"

	"github.com/merliot/device/uart"
)

var fs embed.FS

func newTransport(tty string) *uart.Uart {
	u := uart.New()
	u.SetFormat(8, 2, machine.ParityNone) // 8N2
	return u
}
