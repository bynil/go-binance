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

// VIPLoanOngoingOrdersService submits a borrowing request for VIP loan.
//
// See https://binance-docs.github.io/apidocs/spot/en/#vip-user_data
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
