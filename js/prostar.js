import { WebSocketController, ViewMode } from './common.js'

export function run(prefix, url, viewMode) {
	const prostar = new Prostar(prefix, url, viewMode)
}

class Prostar extends WebSocketController {

	open() {
		super.open()

		if (this.state.DeployParams === "") {
			return
		}

		this.show()
	}

	show() {
		this.showStatus()
		this.showSystem()
		this.showController()
		this.showBattery()
		this.showLoadInfo()
		this.showSolar()
		this.showDaily()
		this.showHistorical()
	}

	showStatus() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var status = document.getElementById("status")
			status.value = ""
			status.value += "Status:                      " + this.state.Status
			break;
		}
	}

	showSystem() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("system")
			ta.value = ""
			ta.value += "Software Version:            " + this.state.System.SWVersion + "\r\n"
			ta.value += "Batt Voltage Multiplier:     " + this.state.System.BattVoltMulti
			break;
		}
	}

	showController() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("controller")
			ta.value = ""
			ta.value += "* Current (A):               " + this.state.Controller.Amps
			break;
		}
	}

	showBattery() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("battery")
			ta.value = ""
			ta.value += "* Voltage (V):               " + this.state.Battery.Volts + "\r\n"
			ta.value += "* Current (A):               " + this.state.Battery.Amps + "\r\n"
			ta.value += "* Sense Voltage (V):         " + this.state.Battery.SenseVolts + "\r\n"
			ta.value += "* Slow Filter Voltage (V):   " + this.state.Battery.SlowVolts + "\r\n"
			ta.value += "* Slow Filter Current (A):   " + this.state.Battery.SlowAmps
			break;
		case ViewMode.ViewTile:
			document.getElementById("battery-volts").innerText = this.state.Battery.Volts.toFixed(2)
			document.getElementById("battery-amps").innerText = this.state.Battery.Amps.toFixed(2)
			break;
		}
	}

	showLoadInfo() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("load")
			ta.value = ""
			ta.value += "* Voltage (V):               " + this.state.LoadInfo.Volts + "\r\n"
			ta.value += "* Current (A):               " + this.state.LoadInfo.Amps
			break;
		case ViewMode.ViewTile:
			var volts = document.getElementById("load-volts")
			var amps = document.getElementById("load-amps")
			volts.innerText = this.state.LoadInfo.Volts.toFixed(2)
			amps.innerText = this.state.LoadInfo.Amps.toFixed(2)
			if (this.state.LoadInfo.Amps === 0) {
				volts.style.background = "tomato"
				amps.style.background = "tomato"
			}
			break;
		}
	}

	showSolar() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("solar")
			ta.value = ""
			ta.value += "* Voltage (V):               " + this.state.Solar.Volts + "\r\n"
			ta.value += "* Current (A):               " + this.state.Solar.Amps
			break;
		case ViewMode.ViewTile:
			document.getElementById("solar-volts").innerText = this.state.Solar.Volts.toFixed(2)
			document.getElementById("solar-amps").innerText = this.state.Solar.Amps.toFixed(2)
			break;
		}
	}

	showDaily() {
		/*
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("daily")
			ta.value = ""
			break;
		}
		*/
	}

	showHistorical() {
		/*
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("historical")
			ta.value = ""
			break;
		}
		*/
	}

	handle(msg) {
		switch(msg.Path) {
		case "update/status":
			this.state.Status = msg.Status
			this.showStatus()
			break
		case "update/system":
			this.state.System = msg.System
			this.showSystem()
			break
		case "update/controller":
			this.state.Controller = msg.Controller
			this.showController()
			break
		case "update/battery":
			this.state.Battery = msg.Battery
			this.showBattery()
			break
		case "update/load":
			this.state.LoadInfo = msg.LoadInfo
			this.showLoadInfo()
			break
		case "update/solar":
			this.state.Solar = msg.Solar
			this.showSolar()
			break
		case "update/daily":
			this.state.Daily = msg.Daily
			this.showDaily()
			break
		case "update/historical":
			this.state.Historical = msg.Historical
			this.showHistorical()
			break
		}
	}
}
