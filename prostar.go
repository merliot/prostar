// Morningstar Prostar PWM array charge controller device.
//
// modbus ref: https://www.morningstarcorp.com/wp-content/uploads/technical-doc-prostar-modbus-specification-en.pdf

package prostar

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/merliot/dean"
	"github.com/merliot/device"
	"github.com/merliot/device/modbus"
	"github.com/x448/float16"
)

const (
	regVerSw       = 0x0000
	regAdcIa       = 0x0011
	regAdcIl       = 0x0016
	regChargeState = 0x0021
	regLoadState   = 0x002E
	regVbMinDaily  = 0x0041
)

type System struct {
	SWVersion     string
	BattVoltMulti uint16
}

type Controller struct {
	Amps float32
}

type Battery struct {
	Volts       float32
	SenseVolts  float32
	SlowVolts   float32
	SlowNetAmps float32
}

type LoadInfo struct {
	Volts float32
	Amps  float32
	State uint16
	Fault uint16
}

type Array struct {
	Volts float32
	Amps  float32
	State uint16
}

type Daily struct {
	MinBattVolts float32
	MaxBattVolts float32
	ChargeAh     float32
	LoadAh       float32
}

type msgStatus struct {
	Path   string
	Status string
}

type msgSystem struct {
	Path   string
	System System
}

type msgController struct {
	Path       string
	Controller Controller
}

type msgBattery struct {
	Path    string
	Battery Battery
}

type msgLoadInfo struct {
	Path     string
	LoadInfo LoadInfo
}

type msgArray struct {
	Path  string
	Array Array
}

type msgDaily struct {
	Path  string
	Daily Daily
}

type Prostar struct {
	*device.Device
	modbus.Modbus `json:"-"`
	Status        string
	System        System
	Controller    Controller
	Battery       Battery
	LoadInfo      LoadInfo
	Array         Array
	Daily         Daily
	tty           string
}

var targets = []string{"demo", "x86-64", "rpi", "nano-rp2040"}

func New(id, model, name string) dean.Thinger {
	fmt.Println("NEW PROSTAR\r")
	return &Prostar{
		Device: device.New(id, model, name, fs, targets).(*device.Device),
		Status: "OK",
	}
}

func (p *Prostar) save(pkt *dean.Packet) {
	pkt.Unmarshal(p).Broadcast()
}

func (p *Prostar) getState(pkt *dean.Packet) {
	p.Path = "state"
	pkt.Marshal(p).Reply()
}

func (p *Prostar) Subscribers() dean.Subscribers {
	return dean.Subscribers{
		"state":             p.save,
		"get/state":         p.getState,
		"update/status":     p.save,
		"update/system":     p.save,
		"update/controller": p.save,
		"update/battery":    p.save,
		"update/load":       p.save,
		"update/array":      p.save,
		"update/daily":      p.save,
	}
}

func (p *Prostar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.API(w, r, p)
}

func swap(b []byte) uint16 {
	return (uint16(b[0]) << 8) | uint16(b[1])
}

func unswap(v uint16) []byte {
	return []byte{byte(v >> 8), byte(v)}
}

func noswap(b []byte) uint16 {
	return (uint16(b[1]) << 8) | uint16(b[0])
}

func f16(b []byte) float32 {
	return float16.Float16(swap(b)).Float32()
}

func unf16(f32 float32) []byte {
	f16 := float16.Fromfloat32(f32)
	return []byte{byte(f16 >> 8), byte(f16 & 0xFF)}
}

// bcdToDecimal converts a BCD-encoded uint16 value to decimal.
func bcdToDecimal(bcd uint16) string {
	decimal := uint16(0)
	multiplier := uint16(1)

	for bcd > 0 {
		// Extract the rightmost 4 bits (a decimal digit) from BCD
		digit := bcd & 0xF

		// Add the decimal value of the digit to the result
		decimal += digit * multiplier

		// Move to the next decimal place
		multiplier *= 10

		// Shift BCD to the right by 4 bits
		bcd >>= 4
	}

	// Convert the decimal value to a string
	return strconv.FormatUint(uint64(decimal), 10)
}

// Round to 2 decimal places
func round(num float32) float32 {
	rounded := math.Round(float64(num)*100) / 100
	return float32(rounded)
}

func (p *Prostar) readSystem(s *System) error {
	regs, err := p.ReadRegisters(regVerSw, 2)
	if err != nil {
		return err
	}
	s.SWVersion = bcdToDecimal(swap(regs[0:2]))
	s.BattVoltMulti = noswap(regs[2:4])
	return nil
}

func (p *Prostar) readDynamic(c *Controller, b *Battery, l *LoadInfo, s *Array) error {

	// FILTERED ADC

	regs, err := p.ReadRegisters(regAdcIa, 4)
	if err != nil {
		return err
	}

	s.Amps = round(f16(regs[0:2]))
	b.Volts = round(f16(regs[2:4]))
	s.Volts = round(f16(regs[4:6]))
	l.Volts = round(f16(regs[6:8]))

	regs, err = p.ReadRegisters(regAdcIl, 4)
	if err != nil {
		return err
	}

	l.Amps = round(f16(regs[0:2]))
	b.SenseVolts = round(f16(regs[2:4]))
	b.SlowVolts = round(f16(regs[4:6]))
	b.SlowNetAmps = round(f16(regs[6:8]))

	// CHARGER STATUS

	regs, err = p.ReadRegisters(regChargeState, 1)
	if err != nil {
		return err
	}

	s.State = swap(regs[0:2])

	// LOAD STATUS

	regs, err = p.ReadRegisters(regLoadState, 2)
	if err != nil {
		return err
	}

	l.State = swap(regs[0:2])
	l.Fault = swap(regs[2:4])

	return nil
}

func (p *Prostar) readDaily(d *Daily) error {

	// LOGGER

	regs, err := p.ReadRegisters(regVbMinDaily, 4)
	if err != nil {
		return err
	}

	d.MinBattVolts = round(f16(regs[0:2]))
	d.MaxBattVolts = round(f16(regs[2:4]))
	d.ChargeAh = round(f16(regs[4:6]))
	d.LoadAh = round(f16(regs[6:8]))

	return nil
}

func (p *Prostar) sendStatus(i *dean.Injector, newStatus string) {
	if p.Status == newStatus {
		return
	}

	var status = msgStatus{Path: "update/status"}
	var pkt dean.Packet

	status.Status = newStatus
	i.Inject(pkt.Marshal(status))
}

func (p *Prostar) sendSystem(i *dean.Injector) {
	var system = msgSystem{Path: "update/system"}
	var pkt dean.Packet

	// sendSystem blocks until we get a good system info read

	for {
		if err := p.readSystem(&system.System); err != nil {
			p.sendStatus(i, err.Error())
			continue
		}
		i.Inject(pkt.Marshal(system))
		break
	}

	p.sendStatus(i, "OK")
}

func (p *Prostar) sendDynamic(i *dean.Injector) {
	var controller = msgController{Path: "update/controller"}
	var battery = msgBattery{Path: "update/battery"}
	var loadInfo = msgLoadInfo{Path: "update/load"}
	var array = msgArray{Path: "update/array"}
	var pkt dean.Packet

	err := p.readDynamic(&controller.Controller, &battery.Battery,
		&loadInfo.LoadInfo, &array.Array)
	if err != nil {
		p.sendStatus(i, err.Error())
		return
	}

	// If anything has changed, send update msg(s)

	if controller.Controller != p.Controller {
		i.Inject(pkt.Marshal(controller))
	}
	if battery.Battery != p.Battery {
		i.Inject(pkt.Marshal(battery))
	}
	if loadInfo.LoadInfo != p.LoadInfo {
		i.Inject(pkt.Marshal(loadInfo))
	}
	if array.Array != p.Array {
		i.Inject(pkt.Marshal(array))
	}

	p.sendStatus(i, "OK")
}

func (p *Prostar) sendHourly(i *dean.Injector) {
	var daily = msgDaily{Path: "update/daily"}
	var pkt dean.Packet

	err := p.readDaily(&daily.Daily)
	if err != nil {
		p.sendStatus(i, err.Error())
		return
	}

	// If anything has changed, send update msg(s)

	if daily.Daily != p.Daily {
		i.Inject(pkt.Marshal(daily))
	}

	p.sendStatus(i, "OK")
}

func (p *Prostar) config() {
	tty := p.ParamFirstValue("tty")
	p.Modbus = modbus.New(newTransport(tty))
}

func (p *Prostar) Setup() {
	p.Device.Setup()
	p.config()
}

func (p *Prostar) Run(i *dean.Injector) {

	p.sendSystem(i)
	p.sendDynamic(i)
	p.sendHourly(i)

	nextHour := time.Now().Add(time.Hour)

	for {
		p.sendDynamic(i)
		time.Sleep(5 * time.Second)
		if time.Now().After(nextHour) {
			p.sendHourly(i)
			nextHour = time.Now().Add(time.Hour)
		}
	}
}
