package binance

import (
	"context"
	"net/http"
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
