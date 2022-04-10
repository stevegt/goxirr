/*
Package goxirr is a simple implementation of a function for calculating
the Internal Rate of Return for irregular cash flow (XIRR).
*/
package goxirr

import (
	"math"
	"sort"
	"time"

	. "github.com/stevegt/goadapt"
)

//A Transaction represent a single transaction from a series of irregular payments.
type Transaction struct {
	Date time.Time
	Cash float64
}

//Transactions represent a cash flow consisting of individual transactions
type Transactions []Transaction

var Min float64 = -1000
var Max float64 = 1000
var Epsilon = 0.0001

type Best struct {
	guess    float64
	residual float64
}

type Bestset []Best

func (b Bestset) add(guess, residual float64) Bestset {
	if math.IsNaN(guess) || math.IsNaN(residual) {
		return b
	}
	b = append(b, Best{guess: guess, residual: residual})
	sort.Slice(b, func(x, y int) bool { return math.Abs(b[x].residual) < math.Abs(b[y].residual) })
	if len(b) > 2 {
		b = b[:2]
	}
	return b
}

func (b Bestset) min() Best {
	if len(b) == 0 {
		return Best{guess: math.NaN(), residual: math.MaxFloat64}
	}
	return b[0]
}

func (b Bestset) max() Best {
	return b[len(b)-1]
}

func (b Bestset) residual() float64 {
	return b.min().residual
}

func (b Bestset) absResidual() float64 {
	return math.Abs(b.residual())
}

func (b Bestset) guess() float64 {
	return b.min().guess
}

//Xirr returns the Internal Rate of Return (IRR) for an irregular series of cash flows (XIRR)
func Xirr(transactions Transactions) float64 {
	var years []float64
	for _, t := range transactions {
		years = append(years, (t.Date.Sub(transactions[0].Date).Hours()/24)/365)
	}

	min := Min * .01
	max := Max * .01
	bestGuess := Epsilon
	steps := 10
	var best Bestset

	for {
		span := max - min
		step := span / float64(steps)
		Debug("min %f max %f span %f steps %d step %f\n", min, max, span, steps, step)

		prev := best.min()
		for i := 0; i <= steps; i++ {
			guess := min + float64(i)*step
			residual := getResidual(transactions, years, guess)
			Debug("min", min, "max", max, "i", i, "step", step, "guess", guess, "residual", residual)

			best = best.add(guess, residual)
			Debug("absResidual", best.absResidual(), "best", best)

			if best.absResidual() == 0 {
				break
			}
		}
		Debug("best", best.guess(), best)

		if best.absResidual() < Epsilon {
			bestGuess = best.guess()
			Debug("found", bestGuess)
			break
		}

		if span < Epsilon {
			bestGuess = (max + min) * .5
			Debug("estimating", bestGuess)
			break
		}

		if best.absResidual() == math.Abs(prev.residual) {
			steps *= 10
			Debug("increased steps", steps)
		} else if math.Abs(best.guess()-prev.guess) < Epsilon*.01 {
			bestGuess = best.guess()
			Debug("converging", bestGuess)
			break
		}

		min = (min+best.min().guess)/2.0 - Epsilon*.01
		max = (max+best.max().guess)/2.0 + Epsilon*.01
	}

	return math.Round(bestGuess*100*100) / 100.0
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
