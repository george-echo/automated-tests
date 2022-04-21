//go:build tools
// +build tools

package tools

import (
	_ "github.com/lileio/lile/protoc-gen-lile-server"
	_ "github.com/rakyll/gotest"
	_ "golang.org/x/tools/cmd/goimports"
)
