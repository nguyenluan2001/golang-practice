package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	app "github.com/nguyenluan2001/golang-practice/todolist/server"
	_ "github.com/nguyenluan2001/golang-practice/todolist/server/statik" // TODO: Replace with the absolute import path
	"github.com/valyala/fasthttp"
)

const (
	BUFFER_MAX_SIZE int = 50 * 1024 * 1024 * 1024
	HEADER_MAX_SIZE int = 100 * 1024
)

func serverHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)

}
func initRouter(appPtr *app.AppPtr) *router.Router {
	r := router.New()
	r.ANY("/", appPtr.Homepage)
	apiGroup := r.Group("/api")
	apiGroup.GET("/test", appPtr.Test)
	apiGroup.GET("/todos", appPtr.GetTodoListApi)
	apiGroup.POST("/todo/create", appPtr.CreateTodoApi)
	apiGroup.PUT("/todo/{uid}", appPtr.UpdateTodoApi)
	apiGroup.DELETE("/todo/{uid}", appPtr.DeleteTodoApi)
	return r
}
func test() (instance int) {
	instance = 15
	return
}
func main() {
	fmt.Println("start main")
	fmt.Println(test())
	appPtr := app.AppPtr{}
	appRouter := initRouter(&appPtr)
	funcHandler := appRouter.Handler
	log.Fatal(fasthttp.ListenAndServe(":8081", funcHandler))
}
