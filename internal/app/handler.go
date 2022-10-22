package app

import (
	"avito/internal/models"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// accrualFunds godoc
// @Summary Accrues funds to the user's balance
// @Tags user
// @Description accepts amount and user ID
// @ID accrual-funds
// @Accept  json
// @Param id_and_amount body models.AccrualFunds true "data for accrual"
// @Produce json
// @Success 200 {object} response "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /accrual [post]
// accrualFunds - method of accruing cash to the balance
func (a *App) accrualFunds(ctx *fasthttp.RequestCtx) {
	var ac models.AccrualFunds
	if err := a.parser.UnmarshalBody(ctx, &ac, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
		return
	}
	statusCode, err := a.services.AccrualFunds(ac)
	if err != nil {
		a.logger.Errorf("money transfer error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := fmt.Sprintf("funds have been successfully credited to the balance of the user with id %d", ac.UserID)
	Response(ctx, statusCode, message, true)
}

// getBalance godoc
// @Summary Returns the user's current balance
// @Tags user
// @Description accepts user id
// @ID get-balance
// @Accept  json
// @Param id body models.UserBalance true "user id"
// @Produce json
// @Success 200 {object} response "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /get_balance [post]
// getBalance - method to get the user's balance
func (a *App) getBalance(ctx *fasthttp.RequestCtx) {
	var ub models.UserBalance
	if err := a.parser.UnmarshalBody(ctx, &ub, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
		return
	}
	statusCode, err := a.services.GetBalance(&ub)
	if err != nil {
		a.logger.Errorf("balance check error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := fmt.Sprintf("user's balance with id %d = %.2f", ub.UserID, ub.Balance)
	Response(ctx, statusCode, message, true)
}

// blockFunds godoc
// @Summary Blocks user funds when ordering a service
// @Tags user
// @Description accepts user id, service id, order id, amount
// @ID block-funds
// @Accept  json
// @Param block_request body models.Order true "data for order"
// @Produce json
// @Success 201 {object} response "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /create_order [post]
// blockFunds - method of reserving funds from the main balance in a separate account
func (a *App) blockFunds(ctx *fasthttp.RequestCtx) {
	var order models.Order
	if err := a.parser.UnmarshalBody(ctx, &order, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
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

// unblockFunds godoc
// @Summary Unblocks the user's funds when the service is canceled
// @Tags user
// @Description accepts order id
// @ID unblock-funds
// @Accept  json
// @Param unblock_request body models.Unblock true "order id for unblock"
// @Produce json
// @Success 200 {object} response "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /cancel_order [post]
// unblockFunds - method of reserving money if it was not possible to apply the service
func (a *App) unblockFunds(ctx *fasthttp.RequestCtx) {
	var unblock models.Unblock
	if err := a.parser.UnmarshalBody(ctx, &unblock, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
		return
	}
	statusCode, err := a.services.UnblockFunds(unblock)
	if err != nil {
		a.logger.Errorf("unblock funds error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := fmt.Sprintf("order %d cancelled, funds refunded", unblock.OrderID)
	Response(ctx, statusCode, message, true)
}

// chargeFunds godoc
// @Summary Withdraws previously blocked funds
// @Tags order
// @Description accepts user id, service id, order id, amount
// @ID charge-funds
// @Accept  json
// @Param charge_request body models.Order true "data for charge funds"
// @Produce json
// @Success 200 {object} response "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /charge [post]
// chargeFunds - method for charging previously reserved funds
func (a *App) chargeFunds(ctx *fasthttp.RequestCtx) {
	var order models.Order
	if err := a.parser.UnmarshalBody(ctx, &order, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
		return
	}
	statusCode, err := a.services.ChargeFunds(order)
	if err != nil {
		a.logger.Errorf("charge funds error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := fmt.Sprintf("funds for the order %d have been successfully charged", order.OrderID)
	Response(ctx, statusCode, message, true)
}

// getReport godoc
// @Summary Requests a financial report on paid services for the month
// @Tags order
// @Description accepts year and month
// @ID get-report
// @Accept  json
// @Param report_request body models.Report true "data for get report"
// @Produce json
// @Success 200 {object} response "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /get_report [post]
// getReport - method provides monthly accounting report
func (a *App) getReport(ctx *fasthttp.RequestCtx) {
	var report models.Report
	if err := a.parser.UnmarshalBody(ctx, &report, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
		return
	}
	data, statusCode, err := a.services.GetReport(report)
	if err != nil {
		a.logger.Errorf("get report error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	key := fmt.Sprintf("%d%d", report.Year, report.Month)
	a.store.Lock()
	a.store.reports[key] = data
	a.store.Unlock()
	link := fmt.Sprintf("http://%s:%d/reports/?report=%s", a.config.ServiceHost, a.config.ServicePort, key)
	message := fmt.Sprintf("the requested report has been successfully generated and is available at the link: %s", link)
	Response(ctx, statusCode, message, true)
}

// downloadReport godoc
// @Summary Downloads a file with a report in CSV format
// @Tags order
// @Description accepts report key
// @ID download-file
// @Param report query string true "report key"
// @Success 200 {string} string "success"
// @Failure 404 {object} response "not found"
// @Router /reports/ [get]
// downloadReport - method provides monthly accounting report
func (a *App) downloadReport(ctx *fasthttp.RequestCtx) {
	key := string(ctx.QueryArgs().Peek("report"))
	if report, ok := a.store.IsExist(key); ok {
		ctx.SetStatusCode(200)
		ctx.SetContentType("application/CSV")
		ctx.Write(report)
		return
	}
	a.logger.Errorf("report for key %s not found", key)
	Response(ctx, 404, "report not found", false)
}

// transferFunds godoc
// @Summary Makes a transfer of funds between two users
// @Tags user
// @Description accepts sender id, receiver id, amount
// @ID transfer-funds
// @Accept  json
// @Param transfer_request body models.Transfer true "data for transfer funds"
// @Produce json
// @Success 200 {object} response "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /transfer [post]
// transferFunds - method for transferring funds between users
func (a *App) transferFunds(ctx *fasthttp.RequestCtx) {
	var t models.Transfer
	if err := a.parser.UnmarshalBody(ctx, &t, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
		return
	}
	statusCode, err := a.services.TransferFunds(t)
	if err != nil {
		a.logger.Errorf("transfer error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	message := "funds transfer completed successfully"
	Response(ctx, statusCode, message, true)
}

// getUserTransactions godoc
// @Summary Requests a list of all user transactions with comments
// @Tags transaction
// @Description accepts user id, order by data, limit and offset
// @ID get-transactions
// @Accept  json
// @Param transactions_request body models.TransactionListRequest true "data for get transactions"
// @Produce json
// @Success 200 {array} models.TransactionList "success"
// @Failure 400 {object} response "bad request"
// @Failure 500 {object} response "server error"
// @Router /transactions [post]
// getUserTransactions - method to get list of user's transactions
func (a *App) getUserTransactions(ctx *fasthttp.RequestCtx) {
	var tr models.TransactionListRequest
	if err := a.parser.UnmarshalBody(ctx, &tr, true); err != nil {
		a.logger.Errorf("data parsing error: %s", err.Error())
		Response(ctx, 400, err.Error(), false)
		return
	}
	tl, statusCode, err := a.services.GetUserTransactions(tr)
	if err != nil {
		a.logger.Errorf("transaction list getting error: %s", err.Error())
		Response(ctx, statusCode, err.Error(), false)
		return
	}
	ctx.SetStatusCode(200)
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(tl)
}
