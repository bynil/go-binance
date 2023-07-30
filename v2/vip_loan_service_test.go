package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type vipLoanServiceTestSuite struct {
	baseTestSuite
}

func TestVIPLoanService(t *testing.T) {
	suite.Run(t, new(vipLoanServiceTestSuite))
}

func (s *vipLoanServiceTestSuite) TestRepay() {
	data := []byte(`
		{
			"loanCoin": "BUSD",
			"repayAmount": "100",
			"remainingPrincipal": "100.5",
			"remainingInterest": "1.5",
			"collateralCoin": "BNB,BTC,ETH",
			"currentLTV": "0.25",
			"repayStatus": "Repaid"
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	orderID := "756783308056935434"
	amount := "100"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"orderId": orderID,
			"amount":  amount,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewVIPLoanRepayService().
		OrderID(orderID).
		Amount(amount).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(&VIPLoanRepayResponse{
		LoanCoin:           "BUSD",
		RepayAmount:        "100",
		RemainingPrincipal: "100.5",
		RemainingInterest:  "1.5",
		CollateralCoin:     "BNB,BTC,ETH",
		CurrentLTV:         "0.25",
		RepayStatus:        RepayStatusRepaid,
	}, res)
}

func (s *vipLoanServiceTestSuite) TestBorrow() {
	data := []byte(`
		{
			"loanAccountId": "12345678",
			"requestId": "12345678",
			"loanCoin": "BTC",
			"loanAmount": "100.55",
			"collateralAccountId": "123456781,123456782,123456783",
			"collateralCoin": "BUSD,USDT,ETH",
			"loanTerm": "30"
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	loanAccountID := "12345678"
	loanCoin := "BTC"
	loanAmount := "100.55"
	collateralAccountID := "123456781,123456782,123456783"
	collateralCoin := "BUSD,USDT,ETH"
	loanTerm := 30
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"loanAccountId":       loanAccountID,
			"loanCoin":            loanCoin,
			"loanAmount":          loanAmount,
			"collateralAccountId": collateralAccountID,
			"collateralCoin":      collateralCoin,
			"loanTerm":            loanTerm,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewVIPLoanBorrowService().
		LoanAccountID(loanAccountID).
		LoanCoin(loanCoin).
		LoanAmount(loanAmount).
		CollateralAccountId(collateralAccountID).
		CollateralCoin(collateralCoin).
		LoanTerm(loanTerm).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(&VIPLoanBorrowResponse{
		LoanAccountId:       "12345678",
		RequestId:           "12345678",
		LoanCoin:            "BTC",
		LoanAmount:          "100.55",
		CollateralAccountId: "123456781,123456782,123456783",
		CollateralCoin:      "BUSD,USDT,ETH",
		LoanTerm:            "30",
	}, res)
}
