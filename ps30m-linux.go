//go:build !tinygo

package ps30m

import "embed"

//go:embed css go.mod *.go html images js template
var fs embed.FS
