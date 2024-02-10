//go:build !tinygo

package ps30m

import (
	"net/http"
)

func (p *Ps30m) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.API(w, r, p)
}
