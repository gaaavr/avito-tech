// Package parser is used to parse data into json format and validate data.
package parser

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
)

// Parser - structure for working with data parsing and validation
type Parser struct {
	validator *validator.Validate
}

// NewParser - constructor function for Parser
func NewParser() *Parser {
	return &Parser{
		validator: validator.New(),
	}
}

// UnmarshalBody -  function converts the data from the request body to json format
// and, if necessary, validates the received data
func (p *Parser) UnmarshalBody(ctx *fasthttp.RequestCtx, data interface{}, validate bool) error {
	if err := json.Unmarshal(ctx.Request.Body(), &data); err != nil {
		return err
	}
	if validate {
		if err := p.validator.Struct(data); err != nil {
			return fmt.Errorf("invalid data for request: %s", err.Error())
		}
	}
	return nil
}
