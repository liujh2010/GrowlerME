package match

import (
	"errors"

	"github.com/growler/common"
	"github.com/growler/queue"
)

// Matcher is a match module for trading pairs.
// It holds the buying and selling queues and through the channel for external input and output order.
// The specific settlement logic can be customization.
type Matcher struct {
	id               int
	enableOfferPrice bool
	lastPrices       map[common.Pair]uint64
	asks             map[common.Pair]queue.IBuyerQueue
	bids             map[common.Pair]queue.ISellerQueue
	in               chan common.IOrder
	outO             chan<- common.IOrder
}

// CreateMatcher returns *Matcher and it use to create a new Matcher.
func CreateMatcher(id int, askQueues map[common.Pair]queue.IBuyerQueue, bidQueues map[common.Pair]queue.ISellerQueue,
	bufferSize int, outO chan<- common.IOrder, enableOfferPrice bool) *Matcher {

	var lastPrices map[common.Pair]uint64 = nil
	if enableOfferPrice {
		lastPrices = make(map[common.Pair]uint64)
	}

	m := &Matcher{
		id:               id,
		enableOfferPrice: enableOfferPrice,
		lastPrices:       lastPrices,
		asks:             askQueues,
		bids:             bidQueues,
		in:               make(chan common.IOrder, bufferSize),
		outO:             outO,
	}

	return m
}

// ID returns the unique identifier of each Matcher.
func (m *Matcher) ID() int {
	return m.id
}

// InChannel returns the order come in channel.
func (m *Matcher) InChannel() chan<- common.IOrder {
	return m.in
}

//Exist checks whether the order is in the order queue.
func (m *Matcher) Exist(o common.IOrder) bool {
	if o.IsBuyer() {
		return m.asks[o.Pair()].Exist(o)
	}
	return m.bids[o.Pair()].Exist(o)
}

// LastPrice returns the last price in the engine.
func (m *Matcher) LastPrice(pair common.Pair) (uint64, error) {
	if !m.enableOfferPrice {
		return 0, errors.New("This feature is not enabled. To enable this feature, turn on enableOfferLimitPrice when creating Matcher")
	}

	return m.lastPrices[pair], nil
}

// Work lets Matcher start receiving orders in the in channel and matching.
func (m *Matcher) Work() {
	for order := range m.in {
		switch order.Action() {
		case common.MATCH:
			m.matching(order)
		case common.CANCEL:
			m.cancel(order)
		}
	}
}

// matching operation queue and call matching method for settlement.
func (m *Matcher) matching(new common.IOrder) {
	ask := m.asks[new.Pair()]
	bid := m.bids[new.Pair()]

	switch new.IsBuyer() {
	case true:
		for {
			if new.IsLimit() && new.Price() < bid.LowestPrice() {
				ask.Push(new)
				break
			}
			old, ok := bid.Lowest()
			if !ok {
				if new.IsLimit() {
					ask.Push(new)
				} else {
					new.CloseOrder()
					m.outO <- new
				}
				break
			}

			new.Take(old, m.offerPrice(new, old))

			if old.IsClose() {
				bid.Deal()
				m.outO <- old
			}

			if new.IsClose() {
				m.outO <- new
				break
			}
		}
	case false:
		for {
			if new.IsLimit() && new.Price() > ask.HighestPrice() {
				bid.Push(new)
				break
			}
			old, ok := ask.Highest()
			if !ok {
				if new.IsLimit() {
					bid.Push(new)
				} else {
					new.CloseOrder()
					m.outO <- new
				}
				break
			}

			new.Take(old, m.offerPrice(new, old))

			if old.IsClose() {
				ask.Deal()
				m.outO <- old
			}

			if new.IsClose() {
				m.outO <- new
				break
			}
		}
	}
}

func (m *Matcher) cancel(o common.IOrder) {
	if o.IsLimit() {
		if o.IsBuyer() {
			m.asks[o.Pair()].Del(o)
		} else {
			m.bids[o.Pair()].Del(o)
		}
		o.CancelOrder()
		m.outO <- o
	}
}

func (m *Matcher) offerPrice(new common.IOrder, old common.IOrder) uint64 {

	if m.enableOfferPrice {
		pair := new.Pair()
		lp := m.lastPrices[pair]

		var cp, bp, sp uint64

		if new.IsLimit() && old.IsLimit() {
			// Limit
			if new.IsBuyer() {
				bp = new.Price()
				sp = old.Price()
			} else {
				bp = old.Price()
				sp = new.Price()
			}

			if sp > lp {
				cp = sp
			} else if lp < bp {
				cp = lp
			} else {
				cp = bp
			}

		} else {
			// Market
			if new.IsLimit() {
				cp = new.Price()
			} else {
				cp = old.Price()
			}
		}

		m.lastPrices[pair] = cp
		return cp
	}

	return old.Price()
}
