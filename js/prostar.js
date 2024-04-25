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
		this.showArray()
		this.showDaily()
	}

	showStatus() {
		/*
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var status = document.getElementById("status")
			status.value = ""
			status.value += "Status:                      " + this.state.Status
			break;
		}
		*/
	}

	showSystem() {
		/*
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("system")
			ta.value = ""
			ta.value += "Software Version:            " + this.state.System.SWVersion + "\r\n"
			ta.value += "Batt Voltage Multiplier:     " + this.state.System.BattVoltMulti
			break;
		}
		*/
	}

	showController() {
		/*
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			var ta = document.getElementById("controller")
			ta.value = ""
			ta.value += "* Current (A):               " + this.state.Controller.Amps
			break;
		}
		*/
	}

	showBattery() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			document.getElementById("net-battery-current").innerText = this.state.Battery.SlowNetAmps
			document.getElementById("battery-terminal-voltage").innerText = this.state.Battery.Volts
			break;
		case ViewMode.ViewTile:
			document.getElementById("battery-volts").innerText = this.state.Battery.Volts
			document.getElementById("battery-amps").innerText = this.state.Battery.SlowNetAmps
			break;
		}
	}

	loadState(state) {
		switch (state) {
		case 0: return "START";
		case 1: return "LOAD ON";
		case 2: return "LVD WARNING";
		case 3: return "LVD";
		case 4: return "FAULT";
		case 5: return "DISCONNECT";
		case 6: return "LOAD OFF";
		case 7: return "OVERRIDE";
		default: return "??";
		}
	}

	showLoadInfo() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			document.getElementById("daily-load").innerText = this.state.Daily.LoadAh
			document.getElementById("load-state").innerText = this.loadState(this.state.LoadInfo.State)
			document.getElementById("load-voltage").innerText = this.state.LoadInfo.Volts
			document.getElementById("load-current").innerText = this.state.LoadInfo.Amps
			break;
		case ViewMode.ViewTile:
			var volts = document.getElementById("load-volts")
			var amps = document.getElementById("load-amps")
			volts.innerText = this.state.LoadInfo.Volts
			amps.innerText = this.state.LoadInfo.Amps
			if (this.state.LoadInfo.Amps === 0) {
				volts.style.background = "tomato"
				amps.style.background = "tomato"
			}
			break;
		}
	}

	chargeState(state) {
		switch (state) {
		case 0: return "START";
		case 1: return "NIGHT CHECK";
		case 2: return "DISCONNECT";
		case 3: return "NIGHT";
		case 4: return "FAULT";
		case 5: return "BULK";
		case 6: return "ABSORPTION";
		case 7: return "FLOAT";
		case 8: return "EQUALIZE";
		default: return "??";
		}
	}

	showArray() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			document.getElementById("array-voltage").innerText = this.state.Array.Volts
			document.getElementById("charge-power").innerText = this.state.Array.Volts * this.state.Array.Amps
			document.getElementById("charge-state").innerText = this.chargeState(this.state.Array.State)
			break;
		case ViewMode.ViewTile:
			document.getElementById("solar-volts").innerText = this.state.Array.Volts
			document.getElementById("solar-amps").innerText = this.state.Array.Amps
			break;
		}
	}

	showDaily() {
		switch (this.viewMode) {
		case ViewMode.ViewFull:
			document.getElementById("daily-system-charge").innerText = this.state.Daily.ChargeAh
			break;
		}
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
			this.state.Array = msg.Array
			this.showArray()
			break
		case "update/daily":
			this.state.Daily = msg.Daily
			this.showDaily()
			break
		}
	}
}
