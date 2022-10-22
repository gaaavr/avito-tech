package app

import "github.com/fasthttp/router"

// Routing - method for registering all handlers in the router
func (a *App) Routing() *router.Router {
	router := router.New()
	router.POST("/accrual", a.LogRequests(a.accrualFunds))
	router.POST("/get_balance", a.LogRequests(a.getBalance))
	router.POST("/create_order", a.LogRequests(a.blockFunds))
	router.POST("/charge", a.LogRequests(a.chargeFunds))
	router.POST("/get_report", a.LogRequests(a.getReport))
	router.GET("/reports", a.LogRequests(a.downloadReport))
	router.POST("/transfer", a.LogRequests(a.transferFunds))
	router.POST("/transactions", a.LogRequests(a.getUserTransactions))
	router.POST("/cancel_order", a.LogRequests(a.unblockFunds))
	return router
}
