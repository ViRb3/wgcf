package main

import (
	"log"

	"github.com/ViRb3/wgcf/cmd"
	"github.com/ViRb3/wgcf/util"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
