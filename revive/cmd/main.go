package main

import (
	"github.com/cocov-ci/go-plugin-kit/cocov"

	"github.com/cocov-ci/go-plugins/revive/plugin"
)

func main() { cocov.Run(plugin.Run) }
