//go:generate openapi-generator generate -i openapi-spec.yml -g go -o openapi --git-user-id ViRb3 --git-repo-id wgcf/v2 --additional-properties=disallowAdditionalPropertiesIfNotPresent=false,withGoMod=false,isGoSubmodule=true

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
