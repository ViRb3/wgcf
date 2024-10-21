package main

import (
	"log"

	"github.com/ViRb3/wgcf/v2/cmd"
	"github.com/ViRb3/wgcf/v2/util"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
