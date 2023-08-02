package binance

import (
	"context"
	"net/http"
	"strings"
)

// VIPLoanRepayService submits a repayment request for VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#vip-loan-repay-trade
type VIPLoanRepayService struct {
	c       *Client
	orderID string
	amount  string
}

func (s *VIPLoanRepayService) OrderID(v string) *VIPLoanRepayService {
	s.orderID = v
	return s
}

func (s *VIPLoanRepayService) Amount(v string) *VIPLoanRepayService {
	s.amount = v
	return s
}

// Do sends the request.
func (s *VIPLoanRepayService) Do(ctx context.Context) (*VIPLoanRepayResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/vip/repay",
		secType:  secTypeSigned,
	}
	r.setParam("orderId", s.orderID)
	r.setParam("amount", s.amount)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &VIPLoanRepayResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type VIPLoanRepayResponse struct {
	LoanCoin           string      `json:"loanCoin"`
	RepayAmount        string      `json:"repayAmount"`
	RemainingPrincipal string      `json:"remainingPrincipal"`
	RemainingInterest  string      `json:"remainingInterest"`
	CollateralCoin     string      `json:"collateralCoin"`
	CurrentLTV         string      `json:"currentLTV"`
	RepayStatus        RepayStatus `json:"repayStatus"` // Repaid, Repaying, Failed
}

// VIPLoanBorrowService submits a borrowing request for VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#vip-loan-borrow-trade
type VIPLoanBorrowService struct {
	c                    *Client
	loanAccountID        string
	loanCoin             string
	loanAmount           string
	collateralAccountIds []string
	collateralCoins      []string
	loanTerm             int
}

func (s *VIPLoanBorrowService) LoanAccountID(v string) *VIPLoanBorrowService {
	s.loanAccountID = v
	return s
}

func (s *VIPLoanBorrowService) LoanCoin(v string) *VIPLoanBorrowService {
	s.loanCoin = v
	return s
}

func (s *VIPLoanBorrowService) LoanAmount(v string) *VIPLoanBorrowService {
	s.loanAmount = v
	return s
}

func (s *VIPLoanBorrowService) CollateralAccountID(v string) *VIPLoanBorrowService {
	s.collateralAccountIds = append(s.collateralAccountIds, v)
	return s
}

func (s *VIPLoanBorrowService) CollateralCoin(v string) *VIPLoanBorrowService {
	s.collateralCoins = append(s.collateralCoins, v)
	return s
}

// LoanTerm is the number of days for the loan. 30/60 days.
func (s *VIPLoanBorrowService) LoanTerm(v int) *VIPLoanBorrowService {
	s.loanTerm = v
	return s
}

// Do sends the request.
func (s *VIPLoanBorrowService) Do(ctx context.Context) (*VIPLoanBorrowResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/vip/borrow",
		secType:  secTypeSigned,
	}
	r.setParam("loanAccountId", s.loanAccountID)
	r.setParam("loanCoin", s.loanCoin)
	r.setParam("loanAmount", s.loanAmount)
	r.setParam("collateralAccountId", strings.Join(s.collateralAccountIds, ","))
	r.setParam("collateralCoin", strings.Join(s.collateralCoins, ","))
	r.setParam("loanTerm", s.loanTerm)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &VIPLoanBorrowResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type VIPLoanBorrowResponse struct {
	LoanAccountID       string `json:"loanAccountId"`
	RequestID           string `json:"requestId"`
	LoanCoin            string `json:"loanCoin"`
	LoanAmount          string `json:"loanAmount"`
	CollateralAccountID string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
	LoanTerm            string `json:"loanTerm"`
}

// VIPLoanOngoingOrdersService query ongoing orders of VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#get-vip-loan-ongoing-orders-user_data
type VIPLoanOngoingOrdersService struct {
	c                   *Client
	orderID             *string
	collateralAccountID *string
	loanCoin            *string
	collateralCoin      *string
	current             *int
	limit               *int
}

func (s *VIPLoanOngoingOrdersService) OrderID(v string) *VIPLoanOngoingOrdersService {
	s.orderID = &v
	return s
}

func (s *VIPLoanOngoingOrdersService) CollateralAccountID(v string) *VIPLoanOngoingOrdersService {
	s.collateralAccountID = &v
	return s
}

func (s *VIPLoanOngoingOrdersService) LoanCoin(v string) *VIPLoanOngoingOrdersService {
	s.loanCoin = &v
	return s
}

func (s *VIPLoanOngoingOrdersService) CollateralCoin(v string) *VIPLoanOngoingOrdersService {
	s.collateralCoin = &v
	return s
}

func (s *VIPLoanOngoingOrdersService) Current(v int) *VIPLoanOngoingOrdersService {
	s.current = &v
	return s
}

func (s *VIPLoanOngoingOrdersService) Limit(v int) *VIPLoanOngoingOrdersService {
	s.limit = &v
	return s
}

// Do sends the request.
func (s *VIPLoanOngoingOrdersService) Do(ctx context.Context) (*VIPLoanOngoingOrdersResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/ongoing/orders",
		secType:  secTypeSigned,
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.collateralAccountID != nil {
		r.setParam("collateralAccountId", *s.collateralAccountID)
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.collateralCoin != nil {
		r.setParam("collateralCoin", *s.collateralCoin)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &VIPLoanOngoingOrdersResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type VIPLoanOngoingOrdersResponse struct {
	Rows  []VIPLoanOngoingOrder `json:"rows"`
	Total int                   `json:"total"`
}

type VIPLoanOngoingOrder struct {
	OrderID                          int    `json:"orderId"`
	LoanCoin                         string `json:"loanCoin"`
	TotalDebt                        string `json:"totalDebt"`
	ResidualInterest                 string `json:"residualInterest"`
	CollateralAccountID              string `json:"collateralAccountId"`
	CollateralCoin                   string `json:"collateralCoin"`
	TotalCollateralValueAfterHaircut string `json:"totalCollateralValueAfterHaircut"`
	LockedCollateralValue            string `json:"lockedCollateralValue"`
	CurrentLTV                       string `json:"currentLTV"`
	ExpirationTime                   int64  `json:"expirationTime"`
}

// VIPLoanRepaymentHistoryService query repayment history of VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#get-vip-loan-repayment-history-user_data
type VIPLoanRepaymentHistoryService struct {
	c         *Client
	orderID   *string
	loanCoin  *string
	startTime *int64
	endTime   *int64
	current   *int
	limit     *int
}

func (s *VIPLoanRepaymentHistoryService) OrderID(v string) *VIPLoanRepaymentHistoryService {
	s.orderID = &v
	return s
}

func (s *VIPLoanRepaymentHistoryService) LoanCoin(v string) *VIPLoanRepaymentHistoryService {
	s.loanCoin = &v
	return s
}

func (s *VIPLoanRepaymentHistoryService) StartTime(v int64) *VIPLoanRepaymentHistoryService {
	s.startTime = &v
	return s
}

func (s *VIPLoanRepaymentHistoryService) EndTime(v int64) *VIPLoanRepaymentHistoryService {
	s.endTime = &v
	return s
}

func (s *VIPLoanRepaymentHistoryService) Current(v int) *VIPLoanRepaymentHistoryService {
	s.current = &v
	return s
}

func (s *VIPLoanRepaymentHistoryService) Limit(v int) *VIPLoanRepaymentHistoryService {
	s.limit = &v
	return s
}

// Do sends the request.
func (s *VIPLoanRepaymentHistoryService) Do(ctx context.Context) (*VIPLoanRepayHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/repay/history",
		secType:  secTypeSigned,
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &VIPLoanRepayHistoryResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type VIPLoanRepayHistoryResponse struct {
	Rows  []VIPLoanRepayHistory `json:"rows"`
	Total int                   `json:"total"`
}

type VIPLoanRepayHistory struct {
	LoanCoin       string `json:"loanCoin"`
	RepayAmount    string `json:"repayAmount"`
	CollateralCoin string `json:"collateralCoin"`
	RepayStatus    string `json:"repayStatus"`
	RepayTime      string `json:"repayTime"`
	OrderID        string `json:"orderId"`
}

// VIPLoanRenewService submits a renewal request for VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#vip-loan-renew-trade
type VIPLoanRenewService struct {
	c        *Client
	orderID  string
	loanTerm int
}

func (s *VIPLoanRenewService) OrderID(v string) *VIPLoanRenewService {
	s.orderID = v
	return s
}

func (s *VIPLoanRenewService) LoanTerm(v int) *VIPLoanRenewService {
	s.loanTerm = v
	return s
}

// Do sends the request.
func (s *VIPLoanRenewService) Do(ctx context.Context) (*VIPLoanRenewResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/loan/vip/renew",
		secType:  secTypeSigned,
	}
	r.setParam("orderId", s.orderID)
	r.setParam("loanTerm", s.loanTerm)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &VIPLoanRenewResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type VIPLoanRenewResponse struct {
	LoanAccountID       string `json:"loanAccountId"`
	LoanCoin            string `json:"loanCoin"`
	LoanAmount          string `json:"loanAmount"`
	CollateralAccountID string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
	LoanTerm            string `json:"loanTerm"`
}

// VIPLoanCollateralService query collateral accounts of VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#check-locked-value-of-vip-collateral-account-user_data
type VIPLoanCollateralService struct {
	c                   *Client
	orderID             *string
	collateralAccountID *string
}

func (s *VIPLoanCollateralService) OrderID(v string) *VIPLoanCollateralService {
	s.orderID = &v
	return s
}

func (s *VIPLoanCollateralService) CollateralAccountID(v string) *VIPLoanCollateralService {
	s.collateralAccountID = &v
	return s
}

// Do sends the request.
func (s *VIPLoanCollateralService) Do(ctx context.Context) (*VIPLoanCollateralResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/collateral/account",
		secType:  secTypeSigned,
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.collateralAccountID != nil {
		r.setParam("collateralAccountId", *s.collateralAccountID)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &VIPLoanCollateralResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type VIPLoanCollateralResponse struct {
	Rows  []VIPLoanCollateral `json:"rows"`
	Total int                 `json:"total"`
}

type VIPLoanCollateral struct {
	CollateralAccountID string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
	CollateralValue     string `json:"collateralValue"`
}

// VIPLoanLoanableService query loanable coins of VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#get-loanable-assets-data-user_data
type VIPLoanLoanableService struct {
	c        *Client
	loanCoin *string
	vipLevel *int
}

func (s *VIPLoanLoanableService) LoanCoin(v string) *VIPLoanLoanableService {
	s.loanCoin = &v
	return s
}

func (s *VIPLoanLoanableService) VipLevel(v int) *VIPLoanLoanableService {
	s.vipLevel = &v
	return s
}

// Do sends the request.
func (s *VIPLoanLoanableService) Do(ctx context.Context) (*VIPLoanLoanableResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/loan/vip/loanable/data",
		secType:  secTypeSigned,
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.vipLevel != nil {
		r.setParam("vipLevel", *s.vipLevel)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &VIPLoanLoanableResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type VIPLoanLoanableResponse struct {
	Rows  []VIPLoanLoanableCoin `json:"rows"`
	Total int                   `json:"total"`
}

type VIPLoanLoanableCoin struct {
	LoanCoin             string `json:"loanCoin"`
	DailyInterestRate30  string `json:"_30dDailyInterestRate"`
	YearlyInterestRate30 string `json:"_30dYearlyInterestRate"`
	DailyInterestRate60  string `json:"_60dDailyInterestRate"`
	YearlyInterestRate60 string `json:"_60dYearlyInterestRate"`
	MinLimit             string `json:"minLimit"`
	MaxLimit             string `json:"maxLimit"`
	VipLevel             int    `json:"vipLevel"`
}
