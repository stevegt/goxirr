/*
Package goxirr is a simple implementation of a function for calculating
the Internal Rate of Return for irregular cash flow (XIRR).
*/
package goxirr

import (
	"math"
	"time"
)

//A Transaction represent a single transaction from a series of irregular payments.
type Transaction struct {
	Date time.Time
	Cash float64
}

//Transactions represent a cash flow consisting of individual transactions
type Transactions []Transaction

var Span float64 = 999.9
var SpanFactor float64 = 0.8
var Limit int64 = 99999
var Epsilon = 0.0001

//Xirr returns the Internal Rate of Return (IRR) for an irregular series of cash flows (XIRR)
func Xirr(transactions Transactions) float64 {
	var years []float64
	for _, t := range transactions {
		years = append(years, (t.Date.Sub(transactions[0].Date).Hours()/24)/365)
	}

	limit := Limit
	span := Span
	guess := 0.0

	for limit > 0 {
		limit--

		// approach guess from both high and low brackets -- otherwise
		// we risk residual approaching infinity when IRR == -100%
		guessHi := guess + span
		guessLo := guess - span
		residualHi := getResidual(transactions, years, guessHi)
		residualLo := getResidual(transactions, years, guessLo)

		if math.IsNaN(residualHi) || math.IsNaN(residualLo) {
			span *= .99
			continue
		}

		// fmt.Println("span", span, "guess", guess, "residualHi", residualHi, "residualLo", residualLo)
		if math.Abs(residualHi) < Epsilon && math.Abs(residualLo) < Epsilon {
			break
		}

		if math.Abs(residualHi) < math.Abs(residualLo) {
			guess = (guess + guessHi) * .5
		} else {
			guess = (guess + guessLo) * .5
		}
		span *= SpanFactor

	}

	return math.Round(guess*100*100) / 100
}

func getResidual(transactions Transactions, years []float64, guess float64) (residual float64) {
	for i, t := range transactions {
		cash := t.Cash
		if cash == 0 {
			// set cash to an infinitesimal when there is 0 return or
			// 0 investment -- otherwise IRR will converge toward net
			// cash instead of a percentage
			cash = Epsilon
		}
		residual += cash / math.Pow(1+guess, years[i])
	}
	return
}
