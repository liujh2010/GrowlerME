package match

import (
	"time"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
	"github.com/growler/common"
)

// Action represents the action required by the order book event.
type Action bool

const (
	// ADD means need to add an order.
	ADD Action = true
	// SUB means need to sub an order.
	SUB Action = false
)

// Event defines a standard order book events.
type Event struct {
	Action Action
	IsBuy  bool
	Pair   common.Pair
	Price  uint64
	Amount uint64
	Time   int64
}

type B struct {
	pair       common.Pair
	ask        *treemap.Map
	bid        *treemap.Map
	lastUpdate int64
	depth      int
}

func createB(pair common.Pair, depth int) *B {
	return &B{
		pair: pair,
		ask: treemap.NewWith(func(a, b interface{}) int {
			return utils.UInt64Comparator(b, a)
		}),
		bid:        treemap.NewWith(utils.UInt64Comparator),
		lastUpdate: time.Now().UnixNano(),
		depth:      depth,
	}
}

func (b *B) Pair() common.Pair {
	return b.pair
}

func (b *B) opt(isBuy, isAdd bool, p, a uint64, time int64) {
	if time <= b.lastUpdate {
		return
	}
	b.lastUpdate = time

	var m *treemap.Map

	if isBuy {
		m = b.ask
	} else {
		m = b.bid
	}

	amount, ok := m.Get(p)
	if ok {
		v := amount.(uint64)
		if isAdd {
			v += a
		} else {
			v -= a
		}
		m.Put(p, v)
	} else {
		if m.Size() >= b.depth {
			k, _ := m.Max()
			m.Remove(k)
		}
		m.Put(p, a)
	}
}

type OrderBookChangeHandler func(e *Event, book *B)

type OrderBook struct {
	in      chan *Event
	books   map[common.Pair]*B
	handler OrderBookChangeHandler
}

func CreateOrderBook(pairs []common.Pair, depths []int, inChan chan *Event, handler OrderBookChangeHandler) *OrderBook {
	b := &OrderBook{
		in:      inChan,
		books:   make(map[common.Pair]*B, len(pairs)),
		handler: handler,
	}
	for i, p := range pairs {
		b.books[p] = createB(p, depths[i])
	}
	return b
}

func (o *OrderBook) InChannel() chan<- *Event {
	return o.in
}

func (o *OrderBook) Work() {
	for e := range o.in {
		b := o.books[e.Pair]
		isAdd := false
		if e.Action == ADD {
			isAdd = true
		}
		b.opt(e.IsBuy, isAdd, e.Price, e.Amount, e.Time)
		o.handler(e, b)
	}
}

func (o *OrderBook) GetBook(pair common.Pair) *B {
	return o.books[pair]
}
