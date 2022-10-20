package app

import (
	"avito/internal/models"
	"fmt"
	"github.com/valyala/fasthttp"
)

// accrualFunds - method of accruing cash to the balance
func (a *App) accrualFunds(ctx *fasthttp.RequestCtx) {
	var ac models.AccrualCash
	if err := a.parser.UnmarshalBody(ctx, &ac, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 500, err.Error(), false)
		return
	}
	statusCode, err := a.services.AccrualFunds(ac)
	if err != nil {
		a.logger.Errorf("money transfer error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := fmt.Sprintf("funds have been successfully credited to the balance of the user with id %d", ac.ID)
	Response(ctx, statusCode, message, true)
}

// getBalance - method to get the user's balance
func (a *App) getBalance(ctx *fasthttp.RequestCtx) {
	var ub models.UserBalance
	if err := a.parser.UnmarshalBody(ctx, &ub, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 500, err.Error(), false)
		return
	}
	statusCode, err := a.services.GetBalance(&ub)
	if err != nil {
		a.logger.Errorf("balance check error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := fmt.Sprintf("user's balance with id %d = %.2f", ub.ID, ub.Balance)
	Response(ctx, statusCode, message, true)
}

// blockFunds - method of reserving funds from the main balance in a separate account
func (a *App) blockFunds(ctx *fasthttp.RequestCtx) {
	var order models.Order
	if err := a.parser.UnmarshalBody(ctx, &order, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 500, err.Error(), false)
		return
	}
	statusCode, err := a.services.BlockFunds(order)
	if err != nil {
		a.logger.Errorf("block funds error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := fmt.Sprintf("order %d successfully created, funds reserved", order.OrderID)
	Response(ctx, statusCode, message, true)
}
