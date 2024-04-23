package main

import (
	"log"
	"os"

	"github.com/merliot/prostar"
)

//go:generate go run main.go
func main() {
	prostar := prostar.New("proto", "prostar", "proto").(*prostar.Prostar)
	if err := prostar.GenerateUf2s("../.."); err != nil {
		log.Println("Error generating UF2s:", err)
		os.Exit(1)
	}
}
