module github.com/cocov-ci/go-plugins/golangci-lint

go 1.19

require (
	github.com/cocov-ci/go-plugin-kit v0.1.15
	github.com/cocov-ci/go-plugins/common v0.0.0-20230310180557-30d393e7269d
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.8.2
	go.uber.org/zap v1.24.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/levigross/grequests v0.0.0-20221222020224-9eee758d18d5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/cocov-ci/go-plugins/common => ../common
