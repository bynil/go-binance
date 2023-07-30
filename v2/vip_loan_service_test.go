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
