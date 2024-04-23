//go:build !tinygo

package prostar

import "embed"

//go:embed css go.mod *.go html images js template
var fs embed.FS
