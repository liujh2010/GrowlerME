package match

import (
	"testing"
	"time"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/growler/common"
)

func TestOrderBookAmount(t *testing.T) {
	tests := []struct {
		es     []*Event
		amount uint64
	}{
		{
			es: []*Event{
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "1",
					Price:  1,
					Amount: 1,
					Time:   time.Now().UnixNano(),
				},
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "1",
					Price:  1,
					Amount: 1,
					Time:   time.Now().UnixNano(),
				},
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "1",
					Price:  1,
					Amount: 1,
					Time:   time.Now().UnixNano(),
				},
			},
			amount: 3,
		},
		{
			es: []*Event{
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "1",
					Price:  2,
					Amount: 5,
					Time:   time.Now().UnixNano(),
				},
				{
					Action: SUB,
					IsBuy:  true,
					Pair:   "1",
					Price:  2,
					Amount: 1,
					Time:   time.Now().UnixNano(),
				},
				{
					Action: SUB,
					IsBuy:  true,
					Pair:   "1",
					Price:  2,
					Amount: 1,
					Time:   time.Now().UnixNano(),
				},
			},
			amount: 3,
		},
	}

	in := make(chan *Event, 10)
	ob := CreateOrderBook([]common.Pair{"1"}, []int{1}, in, func(e *Event, book *B) {})
	for i, tt := range tests {
		for _, e := range tt.es {
			ob.InChannel() <- e
		}

		go ob.Work()

		time.Sleep(time.Second * 1)

		if v, ok := ob.GetBook("1").ask.Get(tt.es[0].Price); ok {
			if v.(uint64) != tt.amount {
				t.Errorf("No.%d case error want %d got %d", i+1, tt.amount, v)
			}
		}
	}
}

func TestOrderBookDepthAndSort(t *testing.T) {
	tests := []struct {
		events []*Event
		pair   common.Pair
		depth  int
		exps   []*Event
	}{
		{
			events: []*Event{
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "abc",
					Price:  1,
					Amount: 2,
					Time:   1688663554066225301,
				},
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "abc",
					Price:  3,
					Amount: 2,
					Time:   1688663554066225302,
				},
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "abc",
					Price:  4,
					Amount: 2,
					Time:   1688663554066225303,
				},
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "abc",
					Price:  4,
					Amount: 2,
					Time:   1688663554066225304,
				},
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "abc",
					Price:  2,
					Amount: 2,
					Time:   1688663554066225305,
				},
			},
			pair:  "abc",
			depth: 3,
			exps: []*Event{
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "abc",
					Price:  4,
					Amount: 4,
					Time:   0,
				},
				{
					Action: ADD,
					IsBuy:  true,
					Pair:   "abc",
					Price:  3,
					Amount: 2,
					Time:   0,
				},
				{
					Action: SUB,
					IsBuy:  true,
					Pair:   "abc",
					Price:  2,
					Amount: 2,
					Time:   0,
				},
			},
		},
		{
			events: []*Event{
				{
					Action: ADD,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  1,
					Amount: 2,
					Time:   1688663554066225306,
				},
				{
					Action: ADD,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  3,
					Amount: 2,
					Time:   1688663554066225307,
				},
				{
					Action: ADD,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  4,
					Amount: 2,
					Time:   1688663554066225308,
				},
				{
					Action: ADD,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  4,
					Amount: 2,
					Time:   1688663554066225309,
				},
				{
					Action: ADD,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  2,
					Amount: 2,
					Time:   1688663554066225310,
				},
			},
			pair:  "DEF",
			depth: 3,
			exps: []*Event{
				{
					Action: ADD,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  1,
					Amount: 2,
					Time:   0,
				},
				{
					Action: ADD,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  2,
					Amount: 2,
					Time:   0,
				},
				{
					Action: SUB,
					IsBuy:  false,
					Pair:   "DEF",
					Price:  3,
					Amount: 2,
					Time:   0,
				},
			},
		},
	}

	in := make(chan *Event, 10)

	pairs := make([]common.Pair, 0)
	depths := make([]int, 0)
	for _, tt := range tests {
		pairs = append(pairs, tt.pair)
		depths = append(depths, tt.depth)
	}

	ob := CreateOrderBook(pairs, depths, in, func(e *Event, book *B) {})

	go ob.Work()

	for i, tt := range tests {
		for _, e := range tt.events {
			ob.InChannel() <- e
		}

		time.Sleep(time.Millisecond * 300)

		b := ob.GetBook(tt.pair)
		var m *treemap.Map
		if tt.exps[0].IsBuy {
			m = b.ask
		} else {
			m = b.bid
		}
		iterator := m.Iterator()
		for _, exp := range tt.exps {
			index := 1
			iterator.Next()
			if exp.Price != iterator.Key().(uint64) {
				t.Errorf("No.%d case the %d price error want %d got %d", i+1, index, exp.Price, iterator.Key())
			}
			if exp.Amount != iterator.Value().(uint64) {
				t.Errorf("No.%d case the %d amount error want %d got %d", i+1, index, exp.Amount, iterator.Value())
			}
			index++
		}
	}
}
