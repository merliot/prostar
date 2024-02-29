package main

import (
	"log"
	"os"

	"github.com/merliot/ps30m"
)

func main() {
	ps30m := ps30m.New("proto", "ps30m", "proto").(*ps30m.Ps30m)
	if err := ps30m.GenerateUf2s(); err != nil {
		log.Println("Error generating UF2s:", err)
		os.Exit(1)
	}
}
