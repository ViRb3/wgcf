//go:generate openapi-generator generate -i openapi-spec.yml -g go -o openapi --additional-properties=disallowAdditionalPropertiesIfNotPresent=false

package main

import (
	"log"

	"github.com/ViRb3/wgcf/v2/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
