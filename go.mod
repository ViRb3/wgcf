module github.com/ViRb3/wgcf

go 1.13

replace github.com/ViRb3/wgcf/openapi => ./openapi

require (
	github.com/ViRb3/optic-go v0.0.0-20210222034258-dd1353e31eb5
	github.com/ViRb3/sling/v2 v2.0.2
	github.com/ViRb3/wgcf/openapi v0.0.0-00010101000000-000000000000
	github.com/manifoldco/promptui v0.8.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.0
	github.com/spf13/viper v1.8.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/yaml.v2 v2.4.0
)
