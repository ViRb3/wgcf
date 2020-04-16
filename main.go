package main

import (
	"log"
	"wgcf/cmd"
	"wgcf/util"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
