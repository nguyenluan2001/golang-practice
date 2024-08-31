package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nguyenluan2001/golang-practice/todolist/server/model"
	"github.com/nguyenluan2001/golang-practice/todolist/server/util"
	"github.com/valyala/fasthttp"
)

func (appPtr *AppPtr) CreateTodo(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Create todo")
}

func (appPtr *AppPtr) ListTodo(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("List todo")
	util.ApplyCORS(ctx)
	jsonFile, err := os.Open("../database/todos.json")

	if err != nil {
		fmt.Println(err)
		ctx.Err()
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var todos = []model.Todo{}
	json.Unmarshal(byteValue, &todos)
	appPtr.Response(ctx, 200, "", todos)
}
