// +build tools

package tools

// This is a hack so that we can version our ci tools within the project
// using go.mod
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
