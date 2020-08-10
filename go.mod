module wgcf

go 1.13

replace (
	github.com/dghubble/sling v1.3.0 => github.com/ViRb3/sling v1.3.0-new
	wgcf/openapi => ./openapi
)

require (
	github.com/antihax/optional v1.0.0
	github.com/dghubble/sling v1.3.0
	github.com/getkin/kin-openapi v0.19.0
	github.com/manifoldco/promptui v0.7.0
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/pkg/errors v0.8.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71
	wgcf/openapi v0.0.0-00010101000000-000000000000
)
