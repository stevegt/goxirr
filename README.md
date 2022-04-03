# goxirr
[![GoDoc](https://godoc.org/github.com/maksim77/goxirr?status.svg)](https://godoc.org/github.com/maksim77/goxirr)

Goxirr is a simple implementation of a function for calculating the Internal Rate of Return for irregular cash flow (XIRR).

This fork was modified by stevegt to support more extreme conditions:

- handles case when IRR == -100%
- handles case when cashflow == 0 in any given transaction
- uses a bracketing function rather than stepwise successive
  approximation in order to handle cases where the limit of the
  residual is +/- infinity when approaching the correct IRR

There are these tunables:

- Span:  Initial successive approximation bracket size.  You may need
  to increase this if you expect +/- IRR larger than the default.
  Defaults to 999.9.
- SpanFactor:  Span is multiplied by this factor after every
  iteration.  Decreasing this will decrease computation time but will
  also limit accuracy in pathological cases.  Defaults to 0.8.
- Limit:  Maximum number of successive approximation iterations.
  Defaults to 99999.
- Epsilon:  Tolerance for IRR error.  Also used as an infinitesimal
  when cashflow == 0 in any transaction. Increasing this will decrease
  computation time as well as accuracy.  Defaults to 0.0001.

## Links
- [Wikipedia](https://en.wikipedia.org/wiki/Internal_rate_of_return)
- [Excel support](https://support.office.com/en-us/article/XIRR-function-DE1242EC-6477-445B-B11B-A303AD9ADC9D)

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/stevegt/goxirr"
)

func main() {
	goxirr.Limit = 10000000 // e.g.

	firstDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	t1 := goxirr.Transaction{
		Date: firstDate,
		Cash: -123400,
	}

	t2 := goxirr.Transaction{
		Date: firstDate.Add(time.Hour * 24 * 365 * 1),
		Cash: 36200,
	}

	t3 := goxirr.Transaction{
		Date: firstDate.Add(time.Hour * 24 * 365 * 2),
		Cash: 54800,
	}

	t4 := goxirr.Transaction{
		Date: firstDate.Add(time.Hour * 24 * 365 * 3),
		Cash: 48100,
	}

	tas := goxirr.Transactions{t1, t2, t3, t4}
	fmt.Println(goxirr.Xirr(tas))
}
```

