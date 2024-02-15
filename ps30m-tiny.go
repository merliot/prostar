//go:build tinygo

package ps30m

import (
	"machine"

	"github.com/merliot/device/uart"
)

func newTransport() uart.Uart {
	u := uart.New()
	u.SetFormat(8, 2, machine.ParityNone) // 8N2
	return u
}
