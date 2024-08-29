package app

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type AppPtr struct {
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (appPtr *AppPtr) Response(ctx *fasthttp.RequestCtx, status int, message string, data interface{}) interface{} {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	result, err := json.Marshal(response)
	if err != nil {
		return err
	}
	ctx.SetBody(result)
	return nil
}
