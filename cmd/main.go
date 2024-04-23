// go run ./cmd
// go run -tags prime ./cmd
// tinygo flash -target xxx ./cmd

package main

import (
	"github.com/merliot/dean"
	"github.com/merliot/device/runner"
	"github.com/merliot/prostar"
)

var (
	id           = dean.GetEnv("ID", "prostar01")
	name         = dean.GetEnv("NAME", "ProStar")
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
	prostar := prostar.New(id, "prostar", name).(*prostar.Prostar)
	prostar.SetDeployParams(deployParams)
	prostar.SetWifiAuth(ssids, passphrases)
	prostar.SetDialURLs(dialURLs)
	prostar.SetWsScheme(wsScheme)
	runner.Run(prostar, port, portPrime, user, passwd, dialURLs, wsScheme)
}
