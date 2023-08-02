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
		CollateralAccountID(collateralAccountID).
		CollateralCoin(collateralCoin).
		LoanTerm(loanTerm).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(&VIPLoanBorrowResponse{
		LoanAccountID:       "12345678",
		RequestID:           "12345678",
		LoanCoin:            "BTC",
		LoanAmount:          "100.55",
		CollateralAccountID: "123456781,123456782,123456783",
		CollateralCoin:      "BUSD,USDT,ETH",
		LoanTerm:            "30",
	}, res)
}

func (s *vipLoanServiceTestSuite) TestOngoingOrders() {
	data := []byte(`
		{
			"rows": [
				{
					"orderId": 100000001,
					"loanCoin": "BUSD",
					"totalDebt": "10000",
					"residualInterest": "10.27687923",
					"collateralAccountId": "12345678,23456789",
					"collateralCoin": "BNB,BTC,ETH",
					"totalCollateralValueAfterHaircut": "25000.27565492",
					"lockedCollateralValue": "25000.27565492",
					"currentLTV": "0.57",
					"expirationTime": 1575018510000
				}
			],
			"total": 1
		}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	orderID := "12345678"
	collateralAccountID := "123456781,123456782,123456783"
	loanCoin := "BTC"
	collateralCoin := "BUSD,USDT,ETH"
	limit := 20
	current := 30
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"orderId":             orderID,
			"collateralAccountId": collateralAccountID,
			"loanCoin":            loanCoin,
			"collateralCoin":      collateralCoin,
			"current":             current,
			"limit":               limit,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewVIPLoanOngoingOrdersService().
		OrderID(orderID).
		CollateralAccountID(collateralAccountID).
		LoanCoin(loanCoin).
		CollateralCoin(collateralCoin).
		Current(current).
		Limit(limit).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(&VIPLoanOngoingOrdersResponse{
		Rows: []VIPLoanOngoingOrder{
			{
				OrderID:                          100000001,
				LoanCoin:                         "BUSD",
				TotalDebt:                        "10000",
				ResidualInterest:                 "10.27687923",
				CollateralAccountID:              "12345678,23456789",
				CollateralCoin:                   "BNB,BTC,ETH",
				TotalCollateralValueAfterHaircut: "25000.27565492",
				LockedCollateralValue:            "25000.27565492",
				CurrentLTV:                       "0.57",
				ExpirationTime:                   1575018510000,
			},
		},
		Total: 1,
	}, res)
}
