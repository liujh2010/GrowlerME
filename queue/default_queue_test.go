package queue

import (
	"fmt"
	"testing"

	"github.com/growler/common"
)

func TestBuyerDQ(t *testing.T) {
	tests := []struct{ caze, exp []*TestOrder }{
		// price case
		{
			caze: []*TestOrder{
				{
					price: 123456,
					time:  123,
				},
				{
					price: 456789,
					time:  324,
				},
				{
					price: 789123,
					time:  2354,
				},
			},
			exp: []*TestOrder{
				{
					price: 789123,
					time:  2354,
				},
				{
					price: 456789,
					time:  324,
				},
				{
					price: 123456,
					time:  123,
				},
			},
		},
		// time case
		{
			caze: []*TestOrder{
				{
					price: 123456,
					time:  324,
				},
				{
					price: 123456,
					time:  123,
				},
				{
					price: 123456,
					time:  2354,
				},
			},
			exp: []*TestOrder{
				{
					price: 123456,
					time:  123,
				},
				{
					price: 123456,
					time:  324,
				},
				{
					price: 123456,
					time:  2354,
				},
			},
		},
		// comprehensive case
		{
			caze: []*TestOrder{
				{
					price: 123456,
					time:  123,
				},
				{
					price: 5637,
					time:  234,
				},
				{
					price: 123456,
					time:  12,
				},
				{
					price: 15534351,
					time:  23424,
				},
				{
					price: 5637,
					time:  123,
				},
				{
					price: 123456,
					time:  5136435,
				},
			},
			exp: []*TestOrder{
				{
					price: 15534351,
					time:  23424,
				},
				{
					price: 123456,
					time:  12,
				},
				{
					price: 123456,
					time:  123,
				},
				{
					price: 123456,
					time:  5136435,
				},
				{
					price: 5637,
					time:  123,
				},
				{
					price: 5637,
					time:  234,
				},
			},
		},
	}

	bdq := CreateBuyerDQ()

	for i, test := range tests {
		for _, order := range test.caze {
			bdq.Push(order)
		}

		values := bdq.q.container.Values()
		fmt.Printf("No.%d case values:\n%v", i+1, values)

		index := 0
		for {
			o, ok := bdq.Highest()
			if !ok {
				break
			}

			if o.Price() != test.exp[index].price || o.Time() != uint64(test.exp[index].time) {
				t.Errorf("No.%d case: error want %v got %v in %d order", i+1, test.exp[index], o, index+1)
			}
			bdq.Deal()
			index++
		}
	}
}

func TestSellerDQ(t *testing.T) {
	tests := []struct{ caze, exp []*TestOrder }{
		// price case
		{
			caze: []*TestOrder{
				{
					price: 789123,
					time:  2354,
				},
				{
					price: 456789,
					time:  324,
				},
				{
					price: 123456,
					time:  123,
				},
			},
			exp: []*TestOrder{
				{
					price: 123456,
					time:  123,
				},
				{
					price: 456789,
					time:  324,
				},
				{
					price: 789123,
					time:  2354,
				},
			},
		},
		// time case
		{
			caze: []*TestOrder{
				{
					price: 123456,
					time:  324,
				},
				{
					price: 123456,
					time:  123,
				},
				{
					price: 123456,
					time:  2354,
				},
			},
			exp: []*TestOrder{
				{
					price: 123456,
					time:  123,
				},
				{
					price: 123456,
					time:  324,
				},
				{
					price: 123456,
					time:  2354,
				},
			},
		},
		// comprehensive case
		{
			caze: []*TestOrder{
				{
					price: 123456,
					time:  123,
				},
				{
					price: 5637,
					time:  234,
				},
				{
					price: 123456,
					time:  12,
				},
				{
					price: 15534351,
					time:  23424,
				},
				{
					price: 5637,
					time:  123,
				},
				{
					price: 123456,
					time:  5136435,
				},
			},
			exp: []*TestOrder{
				{
					price: 5637,
					time:  123,
				},
				{
					price: 5637,
					time:  234,
				},
				{
					price: 123456,
					time:  12,
				},
				{
					price: 123456,
					time:  123,
				},
				{
					price: 123456,
					time:  5136435,
				},
				{
					price: 15534351,
					time:  23424,
				},
			},
		},
	}

	sdq := CreateSellerDQ()

	for i, test := range tests {
		for _, order := range test.caze {
			sdq.Push(order)
		}

		values := sdq.q.container.Values()
		fmt.Printf("No.%d case values:\n%v", i+1, values)

		index := 0
		for {
			o, ok := sdq.Lowest()
			if !ok {
				break
			}

			if o.Price() != test.exp[index].price || o.Time() != uint64(test.exp[index].time) {
				t.Errorf("No.%d case: error want %v got %v in %d order", i+1, test.exp[index], o, index+1)
			}
			sdq.Deal()
			index++
		}
	}
}

/*
	TEST ORDER
*/

type TestOrder struct {
	price uint64
	time  int64
}

func (o *TestOrder) Take(match common.IOrder, strikePrice uint64) {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) Pair() common.Pair {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) Price() uint64 {
	return o.price
}

func (o *TestOrder) Time() uint64 {
	return uint64(o.time)
}

func (o *TestOrder) CloseOrder() {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) CancelOrder() {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) IsBuyer() bool {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) IsLimit() bool {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) IsCancel() bool {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) IsClose() bool {
	panic("not implemented") // TODO: Implement
}

func (o *TestOrder) Action() common.OrderAction {
	panic("not implemented") // TODO: Implement
}
