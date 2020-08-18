package common

import (
	"testing"
)

func TestDefaultOrderLimitTake(t *testing.T) {
	tests := []struct {
		target, match, resTarget, resMatch DefaultOrder
	}{
		// 1st group equal
		{
			DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 0,
				Turnover:        44788022096532,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         true,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 0,
				Turnover:        44788022096532,
			},
		},

		// 2nd group target price and amount are high
		{
			DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           88432643,
				amount:          13456,
				RemainingAmount: 13456,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 440042,
				Turnover:        1189949644208,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         true,
				isCancel:        false,
				price:           88432643,
				amount:          13456,
				RemainingAmount: 0,
				Turnover:        1189949644208,
			},
		},

		// 3rd group target price and amount are low
		{
			DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           78691587,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          9841237,
				RemainingAmount: 9841237,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         true,
				isCancel:        false,
				price:           78691587,
				amount:          453498,
				RemainingAmount: 0,
				Turnover:        44788022096532,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          9841237,
				RemainingAmount: 9387739,
				Turnover:        44788022096532,
			},
		},

		// 4th group target price is high and amount is low
		{
			DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 53498,
				Turnover:        39504493600000,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           88761234,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 0,
				Turnover:        44253042096532,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           88761234,
				amount:          453498,
				RemainingAmount: 400000,
				Turnover:        4748548496532,
			},
		},

		// 5th group target price is low and amount is high
		{
			DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 375348,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 80850,
				Turnover:        38190261907632,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 0,
				Turnover:        35244682138782,
			},
		},

		// 6th group target is the wrong order
		{
			DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          0,
				RemainingAmount: 0,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           98761234,
				amount:          0,
				RemainingAmount: 0,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			},
		},

		// 7th group match is the wrong order
		{
			DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          0,
				RemainingAmount: 0,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           98761234,
				amount:          0,
				RemainingAmount: 0,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			},
		},
	}

	index := 1
	for _, test := range tests {
		test.target.Take(&test.match, test.match.price)

		// price
		if test.target.price != test.resTarget.price {
			t.Errorf("the %v case: target price error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.price, test.target.price, test.target, test.match)
		}

		if test.match.price != test.resMatch.price {
			t.Errorf("the %v case: match price error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.price, test.match.price, test.target, test.match)
		}

		// amount
		if test.target.amount != test.resTarget.amount {
			t.Errorf("the %v case: target amount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.amount, test.target.amount, test.target, test.match)
		}

		if test.match.amount != test.resMatch.amount {
			t.Errorf("the %v case: match amount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.amount, test.match.amount, test.target, test.match)
		}

		// RemainingAmount
		if test.target.RemainingAmount != test.resTarget.RemainingAmount {
			t.Errorf("the %v case: target RemainingAmount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.RemainingAmount, test.target.RemainingAmount, test.target, test.match)
		}

		if test.match.RemainingAmount != test.resMatch.RemainingAmount {
			t.Errorf("the %v case: match RemainingAmount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.RemainingAmount, test.match.RemainingAmount, test.target, test.match)
		}

		// Turnover
		if test.target.Turnover != test.resTarget.Turnover {
			t.Errorf("the %v case: target Turnover error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.Turnover, test.target.Turnover, test.target, test.match)
		}

		if test.match.Turnover != test.resMatch.Turnover {
			t.Errorf("the %v case: match Turnover error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.Turnover, test.match.Turnover, test.target, test.match)
		}

		// isClose
		if test.target.isClose != test.resTarget.isClose {
			t.Errorf("the %v case: target isClose error want %t got %t \ntarget:%v\nmatch:%v", index, test.resTarget.isClose, test.target.isClose, test.target, test.match)
		}

		if test.match.isClose != test.resMatch.isClose {
			t.Errorf("the %v case: match isClose error want %t got %t \ntarget:%v\nmatch:%v", index, test.resMatch.isClose, test.match.isClose, test.target, test.match)
		}
		index++
	}
}

func TestDefaultOrderMarketTake(t *testing.T) {
	tests := []struct {
		target, match, resTarget, resMatch DefaultOrder
	}{
		/*
			buyer market
		*/
		// 1st group tarket is market buyer and it's turnover equal match turnover
		{
			DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          44788022096532,
				RemainingAmount: 44788022096532,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           0,
				amount:          44788022096532,
				RemainingAmount: 0,
				Turnover:        44788022096532,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         true,
				isCancel:        false,
				price:           98761234,
				amount:          453498,
				RemainingAmount: 0,
				Turnover:        44788022096532,
			},
		},

		// 2nd group tarket is market buyer and it's turnover less than match turnover
		{
			DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          1189949604201,
				RemainingAmount: 1189949604201,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           88432643,
				amount:          13456,
				RemainingAmount: 13456,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           0,
				amount:          1189949604201,
				RemainingAmount: 0, // Actual remainder 88392636
				Turnover:        1189861211565,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           88432643,
				amount:          13456,
				RemainingAmount: 1,
				Turnover:        1189861211565,
			},
		},

		// 3rd group tarket is market buyer and it's turnover more than match turnover
		{
			DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          984933110293158,
				RemainingAmount: 984933110293158, //9972871 131,634‬
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           98761234,
				amount:          9841237,
				RemainingAmount: 9841237,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          984933110293158,
				RemainingAmount: 13000400086700,
				Turnover:        971932710206458,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         true,
				isCancel:        false,
				price:           98761234,
				amount:          9841237,
				RemainingAmount: 0,
				Turnover:        971932710206458,
			},
		},

		// 4th group target turnover balance cannot be traded
		{
			DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           88761234,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         false,
				isBuyer:         true,
				isClose:         true,
				isCancel:        true,
				price:           0,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           88761234,
				amount:          453498,
				RemainingAmount: 453498,
				Turnover:        0,
			},
		},

		/*
			seller market
		*/
		// 5th group target is normal order
		{
			DefaultOrder{
				isLimit:         false,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          453498,
				RemainingAmount: 375348,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			}, DefaultOrder{
				isLimit:         false,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          453498,
				RemainingAmount: 80850,
				Turnover:        38190261907632,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         true,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 0,
				Turnover:        35244682138782,
			},
		},

		// 6th group target is the wrong order
		{
			DefaultOrder{
				isLimit:         false,
				isBuyer:         false,
				isClose:         false,
				isCancel:        false,
				price:           0,
				amount:          0,
				RemainingAmount: 0,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			}, DefaultOrder{
				isLimit:         false,
				isBuyer:         false,
				isClose:         true,
				isCancel:        false, // maybe cancel?
				price:           0,
				amount:          0,
				RemainingAmount: 0,
				Turnover:        7718190437100,
			}, DefaultOrder{
				isLimit:         true,
				isBuyer:         true,
				isClose:         false,
				isCancel:        false,
				price:           103471234,
				amount:          340623,
				RemainingAmount: 294498,
				Turnover:        4772610668250,
			},
		},
	}

	index := 1
	for _, test := range tests {
		test.target.Take(&test.match, test.match.price)

		// price
		if test.target.price != test.resTarget.price {
			t.Errorf("the %v case: target price error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.price, test.target.price, test.target, test.match)
		}

		if test.match.price != test.resMatch.price {
			t.Errorf("the %v case: match price error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.price, test.match.price, test.target, test.match)
		}

		// amount
		if test.target.amount != test.resTarget.amount {
			t.Errorf("the %v case: target amount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.amount, test.target.amount, test.target, test.match)
		}

		if test.match.amount != test.resMatch.amount {
			t.Errorf("the %v case: match amount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.amount, test.match.amount, test.target, test.match)
		}

		// RemainingAmount
		if test.target.RemainingAmount != test.resTarget.RemainingAmount {
			t.Errorf("the %v case: target RemainingAmount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.RemainingAmount, test.target.RemainingAmount, test.target, test.match)
		}

		if test.match.RemainingAmount != test.resMatch.RemainingAmount {
			t.Errorf("the %v case: match RemainingAmount error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.RemainingAmount, test.match.RemainingAmount, test.target, test.match)
		}

		// Turnover
		if test.target.Turnover != test.resTarget.Turnover {
			t.Errorf("the %v case: target Turnover error want %d got %d \ntarget:%v\nmatch:%v", index, test.resTarget.Turnover, test.target.Turnover, test.target, test.match)
		}

		if test.match.Turnover != test.resMatch.Turnover {
			t.Errorf("the %v case: match Turnover error want %d got %d \ntarget:%v\nmatch:%v", index, test.resMatch.Turnover, test.match.Turnover, test.target, test.match)
		}

		// isClose
		if test.target.isClose != test.resTarget.isClose {
			t.Errorf("the %v case: target isClose error want %t got %t \ntarget:%v\nmatch:%v", index, test.resTarget.isClose, test.target.isClose, test.target, test.match)
		}

		if test.match.isClose != test.resMatch.isClose {
			t.Errorf("the %v case: match isClose error want %t got %t \ntarget:%v\nmatch:%v", index, test.resMatch.isClose, test.match.isClose, test.target, test.match)
		}
		index++
	}
}

func BenchmarkTakeWorstCase(b *testing.B) {
	target := &DefaultOrder{
		isLimit:         false,
		isBuyer:         true,
		isClose:         false,
		isCancel:        false,
		price:           0,
		amount:          44788022096532,
		RemainingAmount: 44788022096532,
		Turnover:        0,
	}

	match := &DefaultOrder{
		isLimit:         true,
		isBuyer:         false,
		isClose:         false,
		isCancel:        false,
		price:           98761234,
		amount:          453498,
		RemainingAmount: 453498,
		Turnover:        0,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target.Take(match, 98761234)

		// reset order

		target.RemainingAmount = 44788022096532
		target.Turnover = 0
		target.isClose = false
		match.RemainingAmount = 453498
		match.Turnover = 0
		match.isClose = false
	}
}

func BenchmarkTakeNormalCase(b *testing.B) {
	target := &DefaultOrder{
		isLimit:         true,
		isBuyer:         true,
		isClose:         false,
		isCancel:        false,
		price:           98761234,
		amount:          453498,
		RemainingAmount: 453498,
		Turnover:        0,
	}

	match := &DefaultOrder{
		isLimit:         true,
		isBuyer:         false,
		isClose:         false,
		isCancel:        false,
		price:           88432643,
		amount:          13456,
		RemainingAmount: 13456,
		Turnover:        0,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target.Take(match, 98761234)

		// reset order

		target.RemainingAmount = 453498
		target.Turnover = 0
		target.isClose = false
		match.RemainingAmount = 13456
		match.Turnover = 0
		match.isClose = false
	}
}
