package app

import (
	"avito/internal/models"
	"github.com/valyala/fasthttp"
)

// accrualFunds - method of accruing cash to the balance
func (a *App) accrualFunds(ctx *fasthttp.RequestCtx) {
	var ac models.AccrualCash
	if err := a.parser.UnmarshalBody(ctx, &ac, false); err != nil {
		Response(ctx, 500, err.Error(), false)
		return
	}
	statusCode, err := a.services.AccrualFunds(ac)
	if err != nil {
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	Response(ctx, statusCode, "funds have been successfully credited to the user's balance", true)
}
