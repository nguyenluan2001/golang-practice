package utils

import "github.com/valyala/fasthttp"

func EnableCORS(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:5173") // change this later
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

}
