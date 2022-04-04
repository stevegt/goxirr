package goxirr

import (
	"fmt"
	"testing"
	"time"
)

func ExampleXirr() {
	firstDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	t1 := Transaction{
		Date: firstDate,
		Cash: -100,
	}
	t2 := Transaction{
		Date: firstDate.Add(time.Hour * 24 * 365),
		Cash: 112,
	}

	tas := Transactions{t1, t2}
	fmt.Println(Xirr(tas))
	// Output: 12
}

func TestXirr(t *testing.T) {
	type args struct {
		transactions []Transaction
	}

	var case1, case2, case3, case4, case5, case6, case7 args
	case1.transactions = append(case1.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 200,
	})

	case2.transactions = append(case2.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 100,
	})

	case3.transactions = append(case3.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 112,
	})

	case4.transactions = append(case4.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 50,
	})

	case5.transactions = append(case5.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 0,
	})

	case6.transactions = append(case6.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -50,
	})

	case7.transactions = append(case7.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2019, time.January, 12, 0, 0, 0, 0, time.UTC),
		Cash: 101,
	})

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "100%", args: case1, want: 100},
		{name: "0%", args: case2, want: 0.0},
		{name: "12%", args: case3, want: 12},
		{name: "-50%", args: case4, want: -50},
		{name: "-100%", args: case5, want: -100},
		{name: "-150%", args: case6, want: -150},
		{name: "39.12%", args: case7, want: 39.12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xirr(tt.args.transactions); got != tt.want {
				t.Errorf("Xirr() = %v, want %v", got, tt.want)
			}
		})
	}
}
