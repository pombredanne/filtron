package api

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"

	"github.com/asciimoo/filtron/proxy"
)

type API struct {
	Proxy    *proxy.Proxy
	RuleFile string
}

func Listen(address, ruleFile string, p *proxy.Proxy) {
	log.Println("API listens on", address)
	a := &API{p, ruleFile}
	fasthttp.ListenAndServe(address, a.Handler)
}

func (a *API) Handler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	switch string(ctx.Path()) {
	case "/rules/reload":
		if err := a.Proxy.ReloadRules(a.RuleFile); err != nil {
			ctx.Error(fmt.Sprintf("{\"error\": \"%v\"}", err), 500)
			return
		}
		log.Println("Rule file reloaded")
		ctx.Write([]byte("{\"status\": \"ok\"}"))
	default:
		ctx.Error("{\"error\": \"Not found\"}", fasthttp.StatusNotFound)
	}
}
