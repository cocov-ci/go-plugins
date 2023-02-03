module github.com/cocov-ci/go-plugins/staticcheck

go 1.19

require (
	github.com/cocov-ci/go-plugin-kit v0.1.11
	github.com/cocov-ci/go-plugins/common v0.0.0-20230109083926-1f205d5c5b17
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.24.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/cocov-ci/go-plugins/common => ../common