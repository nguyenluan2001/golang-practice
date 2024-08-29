package util

import "github.com/valyala/fasthttp"

func ApplyCORS(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "*")
}
