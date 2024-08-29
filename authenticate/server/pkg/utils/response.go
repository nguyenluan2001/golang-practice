package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type SResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Response(ctx *fasthttp.RequestCtx, status string, data interface{}) {
	responseData := SResponse{
		Status: status,
		Data:   data,
	}
	fmt.Println("data: ", responseData)
	content, err := json.Marshal(
		responseData,
	)
	fmt.Println("Content: ", string(content))
	if err != nil {
		log.Fatalln("Response failed")
	}
	ctx.Response.SetBody(content)
}
