//go:build !tinygo

package ps30m

import (
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/merliot/device"
)

//go:embed css html images js template
var fs embed.FS

type osStruct struct {
	templates *template.Template
}

func (p *Ps30m) osNew() {
	p.CompositeFs.AddFS(fs)
	p.templates = p.CompositeFs.ParseFS("template/*")
}

func (p *Ps30m) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch strings.TrimPrefix(r.URL.Path, "/") {
	case "state":
		device.ShowState(p.templates, w, p)
	default:
		p.API(p.templates, w, r)
	}
}

func (p *Ps30m) DescHtml() []byte {
	desc, _ := fs.ReadFile("html/desc.html")
	return desc
}

func (p *Ps30m) SupportedTargets() string {
	return p.Targets.FullNames()
}
