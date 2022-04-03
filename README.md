# goxirr
[![GoDoc](https://godoc.org/github.com/maksim77/goxirr?status.svg)](https://godoc.org/github.com/maksim77/goxirr)

Goxirr is a simple implementation of a function for calculating the Internal Rate of Return for irregular cash flow (XIRR).

This version modified by stevegt to export three tunables to support
more extreme conditions:

- Guess: initial IRR guess; defaults to .1
- Step: initial successive approximation step size; defaults to .1
- Limit: maximum number of successive approximation iterations; defaults to 100000

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

