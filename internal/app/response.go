package app

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// response  - structure for writing response for requests
type response struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
}

// Response - function for response for requests
func Response(ctx *fasthttp.RequestCtx, statusCode int, description string, status bool) {
	encoder := json.NewEncoder(ctx)
	encoder.SetIndent("", "\t")
	ctx.Response.SetStatusCode(statusCode)
	err := encoder.Encode(response{
		Success:     status,
		Description: description,
	})
	if err != nil {
		ctx.Write([]byte("an error occurred while marshaling the response: " + err.Error()))
		ctx.SetStatusCode(500)
	}
}
