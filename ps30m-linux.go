//go:build !tinygo

package ps30m

import (
	"net/http"
)

func (p *Ps30m) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.API(w, r, p)
}

func (p *Ps30m) DescHtml() []byte {
	desc, _ := fs.ReadFile("html/desc.html")
	return desc
}

func (p *Ps30m) SupportedTargets() string {
	return p.Targets.FullNames()
}
