package main

import (
	"log"

	"github.com/tunedmystic/konnect/cmd"
)

func init() {
	// https://goo.gl/nPMoCL
	log.SetFlags(0)
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
