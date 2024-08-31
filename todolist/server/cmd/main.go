package main

import (
	"github.com/fasthttp/router"
	"github.com/nguyenluan2001/golang-practice/todolist/server/app"
	"github.com/valyala/fasthttp"
)

func test(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello golang")
}

func initRouter(appPtr *app.AppPtr) *router.Router {
	var r = router.New()
	r.GET("/api/create", appPtr.CreateTodo)
	r.GET("/api/todos", appPtr.ListTodo)
	return r
}

func main() {

	// var appRouter = initRouter()
	var appPtr = &app.AppPtr{}
	fasthttp.ListenAndServe(":8081", initRouter(appPtr).Handler)

}
