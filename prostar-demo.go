//go:build !x86_64 && !tinygo && !rpi

package prostar

type transport struct {
	start uint16
	words uint16
}

func newTransport(tty string) *transport {
	return &transport{}
}

func (t *transport) Write(buf []byte) (n int, err error) {
	// get start and words from Modbus request
	t.start = (uint16(buf[2]) << 8) | uint16(buf[3])
	t.words = (uint16(buf[4]) << 8) | uint16(buf[5])
	return n, nil
}

func (t *transport) Read(buf []byte) (n int, err error) {
	// simluate a Modbus request read on the device
	res := buf[3:]
	switch t.start {
	case regVerSw:
	case regAdcIa:
		// TODO make this more dynamic using a little bit of random
		copy(res[0:2], unf16(1.3))  // solar.amps
		copy(res[2:4], unf16(14.1)) // battery.volts
		copy(res[4:6], unf16(11.3)) // solar.volts
		copy(res[6:8], unf16(0))    // load.volts
	case regAdcIl:
		// TODO make this more dynamic using a little bit of random
		copy(res[0:2], unf16(3.3)) // load.amps
		copy(res[2:4], unf16(0))   // battery.sensevolts
		copy(res[4:6], unf16(0))   // battery.slowvolts
		copy(res[6:8], unf16(0))   // battery.slowamps
	}
	return int(5 + t.words*2), nil
}
