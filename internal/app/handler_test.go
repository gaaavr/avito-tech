package app

import (
	"avito/configs"
	"avito/internal/models"
	"avito/internal/parser"
	"avito/internal/service"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func getAppMoc() *App {
	return &App{
		services: &service.Service{
			User:        mockUserService{},
			Order:       mockOrderService{},
			Transaction: mockTransactionService{},
		},
		parser: parser.NewParser(),
		logger: new(mockLogger),
	}
}

func TestAccrualFunds(t *testing.T) {
	mockApp := getAppMoc()
	tableTest := []struct {
		testName           string
		data               models.AccrualFunds
		expectedStatusCode int
	}{
		{
			"invalid amount",
			models.AccrualFunds{
				Amount: 0,
				UserID: 1,
			},
			400,
		},
		{
			"invalid user id",
			models.AccrualFunds{
				Amount: 100,
				UserID: -3,
			},
			400,
		},
		{
			"valid data",
			models.AccrualFunds{
				Amount: 100,
				UserID: 1,
			},
			200,
		},
		{
			"internal error",
			models.AccrualFunds{
				Amount: 100,
				UserID: 2,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.accrualFunds(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestGetBalance(t *testing.T) {
	mockApp := getAppMoc()
	tableTest := []struct {
		testName           string
		data               models.UserBalance
		expectedStatusCode int
	}{
		{
			"invalid user id",
			models.UserBalance{
				UserID: -1,
			},
			400,
		},
		{
			"invalid user id",
			models.UserBalance{
				UserID: 0,
			},
			400,
		},
		{
			"valid data",
			models.UserBalance{
				UserID: 1,
			},
			200,
		},
		{
			"internal error",
			models.UserBalance{
				UserID: 2,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.getBalance(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestBlockFunds(t *testing.T) {
	mockApp := getAppMoc()
	tableTest := []struct {
		testName           string
		data               models.Order
		expectedStatusCode int
	}{
		{
			"invalid user id",
			models.Order{
				UserID:    -1,
				OrderID:   1,
				ServiceID: 1,
				Amount:    100,
			},
			400,
		},
		{
			"invalid order id",
			models.Order{
				UserID:    1,
				OrderID:   -1,
				ServiceID: 1,
				Amount:    100,
			},
			400,
		},
		{
			"invalid service id",
			models.Order{
				UserID:    1,
				OrderID:   1,
				ServiceID: 0,
				Amount:    100,
			},
			400,
		},
		{
			"invalid amount",
			models.Order{
				UserID:    1,
				OrderID:   -1,
				ServiceID: 1,
				Amount:    -100,
			},
			400,
		},
		{
			"valid data",
			models.Order{
				UserID:    1,
				OrderID:   1,
				ServiceID: 1,
				Amount:    100,
			},
			200,
		},
		{
			"internal error",
			models.Order{
				UserID:    2,
				OrderID:   1,
				ServiceID: 1,
				Amount:    100,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.blockFunds(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestUnblockFunds(t *testing.T) {
	mockApp := getAppMoc()
	tableTest := []struct {
		testName           string
		data               models.Unblock
		expectedStatusCode int
	}{
		{
			"invalid user id",
			models.Unblock{
				UserID:  -1,
				OrderID: 1,
			},
			400,
		},
		{
			"invalid order id",
			models.Unblock{
				UserID:  1,
				OrderID: -1,
			},
			400,
		},
		{
			"valid data",
			models.Unblock{
				UserID:  1,
				OrderID: 1,
			},
			200,
		},
		{
			"internal error",
			models.Unblock{
				UserID:  2,
				OrderID: 1,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.unblockFunds(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestChargeFunds(t *testing.T) {
	mockApp := getAppMoc()
	tableTest := []struct {
		testName           string
		data               models.Order
		expectedStatusCode int
	}{
		{
			"invalid user id",
			models.Order{
				UserID:    -1,
				OrderID:   1,
				ServiceID: 1,
				Amount:    100,
			},
			400,
		},
		{
			"invalid order id",
			models.Order{
				UserID:    1,
				OrderID:   -1,
				ServiceID: 1,
				Amount:    100,
			},
			400,
		},
		{
			"invalid service id",
			models.Order{
				UserID:    1,
				OrderID:   1,
				ServiceID: 0,
				Amount:    100,
			},
			400,
		},
		{
			"invalid amount",
			models.Order{
				UserID:    1,
				OrderID:   -1,
				ServiceID: 1,
				Amount:    -100,
			},
			400,
		},
		{
			"valid data",
			models.Order{
				UserID:    1,
				OrderID:   1,
				ServiceID: 1,
				Amount:    100,
			},
			200,
		},
		{
			"internal error",
			models.Order{
				UserID:    2,
				OrderID:   1,
				ServiceID: 1,
				Amount:    100,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.chargeFunds(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestGetReport(t *testing.T) {
	mockApp := getAppMoc()
	mockApp.store = &reportStore{
		reports: make(map[string][]byte),
	}
	mockApp.config = &configs.Common{}
	tableTest := []struct {
		testName           string
		data               models.Report
		expectedStatusCode int
	}{
		{
			"invalid year",
			models.Report{
				Year:  -1,
				Month: 1,
			},
			400,
		},
		{
			"invalid month",
			models.Report{
				Year:  2022,
				Month: 13,
			},
			400,
		},
		{
			"valid data",
			models.Report{
				Year:  2022,
				Month: 10,
			},
			200,
		},
		{
			"internal error",
			models.Report{
				Year:  2008,
				Month: 10,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.getReport(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestDownloadReport(t *testing.T) {
	mockApp := getAppMoc()
	mockApp.store = &reportStore{
		reports: map[string][]byte{
			"test": nil,
		},
	}
	tableTest := []struct {
		testName           string
		key                string
		expectedStatusCode int
	}{
		{
			"invalid report key",
			"foo",
			404,
		},
		{
			"valid report key",
			"test",
			200,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		ctx.QueryArgs().Set("report", testCase.key)
		mockApp.downloadReport(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestTransferFunds(t *testing.T) {
	mockApp := getAppMoc()
	tableTest := []struct {
		testName           string
		data               models.Transfer
		expectedStatusCode int
	}{
		{
			"valid data",
			models.Transfer{
				SenderID:   1,
				ReceiverID: 2,
				Amount:     50,
			},
			200,
		},
		{
			"invalid sender id",
			models.Transfer{
				SenderID:   -1,
				ReceiverID: 2,
				Amount:     50,
			},
			400,
		},
		{
			"invalid receiver id",
			models.Transfer{
				SenderID:   1,
				ReceiverID: -2,
				Amount:     50,
			},
			400,
		},
		{
			"invalid amount",
			models.Transfer{
				SenderID:   1,
				ReceiverID: 2,
				Amount:     -50,
			},
			400,
		},
		{
			"internal error",
			models.Transfer{
				SenderID:   2,
				ReceiverID: 2,
				Amount:     50,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.transferFunds(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestGetUserTransactions(t *testing.T) {
	mockApp := getAppMoc()
	tableTest := []struct {
		testName           string
		data               models.TransactionListRequest
		expectedStatusCode int
	}{
		{
			"valid data",
			models.TransactionListRequest{
				UserID: 1,
				Limit:  2,
				Offset: 50,
			},
			200,
		},
		{
			"invalid user id",
			models.TransactionListRequest{
				UserID: -1,
				Limit:  2,
				Offset: 50,
			},
			400,
		},
		{
			"invalid limit value",
			models.TransactionListRequest{
				UserID: 1,
				Limit:  -1,
				Offset: 50,
			},
			400,
		},
		{
			"invalid offset value",
			models.TransactionListRequest{
				UserID: 1,
				Limit:  0,
				Offset: -1,
			},
			400,
		},
		{
			"internal error",
			models.TransactionListRequest{
				UserID: 2,
				Limit:  1,
				Offset: 50,
			},
			500,
		},
	}
	for _, testCase := range tableTest {
		ctx := new(fasthttp.RequestCtx)
		data, _ := json.Marshal(testCase.data)
		ctx.Request.SetBody(data)
		mockApp.getUserTransactions(ctx)
		assert.Equal(t, testCase.expectedStatusCode, ctx.Response.StatusCode(), testCase.testName)
	}
}

func TestMiddleware(t *testing.T) {
	mockApp := getAppMoc()
	testCase := struct {
		AllowHeadersKey   string
		AllowHeadersValue string
		AllowOriginKey    string
		AllowOriginValue  string
		AllowMethodsKey   string
		AllowMethodsValue string
	}{
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-Secret, Access-Control-Allow-Origin, Access-Control-Allow-Headers, X-CSRF-Token, Authorization-Token",
		"Access-Control-Allow-Origin",
		"*",
		"Access-Control-Allow-Methods",
		"POST, GET, OPTIONS",
	}
	ctx := new(fasthttp.RequestCtx)
	h := mockApp.LogRequests(mockApp.getUserTransactions)
	h(ctx)
	assert.Equal(t, testCase.AllowHeadersValue, string(ctx.Response.Header.Peek(testCase.AllowHeadersKey)))
	assert.Equal(t, testCase.AllowOriginValue, string(ctx.Response.Header.Peek(testCase.AllowOriginKey)))
	assert.Equal(t, testCase.AllowMethodsValue, string(ctx.Response.Header.Peek(testCase.AllowMethodsKey)))
}

type mockUserService struct{}
type mockOrderService struct{}
type mockTransactionService struct{}

func (ms mockUserService) AccrualFunds(ac models.AccrualFunds) (code int, err error) {
	if ac.UserID == 2 {
		return 500, fmt.Errorf("internal error")
	}
	return 200, nil
}
func (ms mockUserService) GetBalance(ub *models.UserBalance) (code int, err error) {
	if ub.UserID == 2 {
		return 500, fmt.Errorf("internal error")
	}
	return 200, nil
}
func (ms mockUserService) BlockFunds(order models.Order) (code int, err error) {
	if order.UserID == 2 {
		return 500, fmt.Errorf("internal error")
	}
	return 200, nil
}
func (ms mockUserService) TransferFunds(t models.Transfer) (code int, err error) {
	if t.SenderID == 2 {
		return 500, fmt.Errorf("internal error")
	}
	return 200, nil
}
func (ms mockUserService) UnblockFunds(unblock models.Unblock) (code int, err error) {
	if unblock.UserID == 2 {
		return 500, fmt.Errorf("internal error")
	}
	return 200, nil
}
func (ms mockOrderService) ChargeFunds(order models.Order) (code int, err error) {
	if order.UserID == 2 {
		return 500, fmt.Errorf("internal error")
	}
	return 200, nil
}
func (ms mockOrderService) GetReport(report models.Report) (data []byte, code int, err error) {
	if report.Year == 2008 {
		return nil, 500, fmt.Errorf("internal error")
	}
	return []byte("ok"), 200, nil
}
func (ms mockTransactionService) GetUserTransactions(tr models.TransactionListRequest) (tl []models.TransactionList, code int, err error) {
	if tr.UserID == 2 {
		return nil, 500, fmt.Errorf("internal error")
	}
	return nil, 200, nil
}

type mockLogger struct{}

func (ml *mockLogger) Errorf(format string, args ...interface{}) {}
func (ml *mockLogger) Fatalf(format string, args ...interface{}) {}
func (ml *mockLogger) Fatal(args ...interface{})                 {}
func (ml *mockLogger) Infof(format string, args ...interface{})  {}
func (ml *mockLogger) Info(args ...interface{})                  {}
func (ml *mockLogger) Warnf(format string, args ...interface{})  {}
func (ml *mockLogger) Debugf(format string, args ...interface{}) {}
func (ml *mockLogger) Debug(args ...interface{})                 {}
func (ml *mockLogger) Panicf(format string, args ...interface{}) {}
func (ml *mockLogger) Printf(format string, args ...interface{}) {}
