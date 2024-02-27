// go run ./cmd
// go run -tags prime ./cmd
// tinygo flash -target xxx ./cmd

package main

import (
	"github.com/merliot/dean"
	"github.com/merliot/device/runner"
	"github.com/merliot/ps30m"
)

var (
	id           = dean.GetEnv("ID", "ps30m01")
	name         = dean.GetEnv("NAME", "Morningstar ps30m")
	deployParams = dean.GetEnv("DEPLOY_PARAMS", "")
	wsScheme     = dean.GetEnv("WS_SCHEME", "ws://")
	port         = dean.GetEnv("PORT", "8000")
	portPrime    = dean.GetEnv("PORT_PRIME", "8001")
	user         = dean.GetEnv("USER", "")
	passwd       = dean.GetEnv("PASSWD", "")
	dialURLs     = dean.GetEnv("DIAL_URLS", "")
	ssids        = dean.GetEnv("WIFI_SSIDS", "")
	passphrases  = dean.GetEnv("WIFI_PASSPHRASES", "")
)

func main() {
	ps30m := ps30m.New(id, "ps30m", name).(*ps30m.Ps30m)
	ps30m.SetDeployParams(deployParams)
	ps30m.SetWifiAuth(ssids, passphrases)
	ps30m.SetDialURLs(dialURLs)
	ps30m.SetWsScheme(wsScheme)
	runner.Run(ps30m.Device, port, portPrime, user, passwd, dialURLs, wsScheme)
}
