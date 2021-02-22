package main

import (
	opticgo "github.com/ViRb3/optic-go"
	"log"
)

const SpecFile = "openapi-spec.json"

func main() {
	if err := opticgo.Format(SpecFile); err != nil {
		log.Fatalln(err)
	}
}
