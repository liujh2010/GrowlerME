package queue

import (
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
	"github.com/growler/common"
)

// DQ is the default implementation of IQueue.
type DQ struct {
	container *treeset.Set
	iterrator treeset.Iterator
}

func createDQ(comparator utils.Comparator) DQ {
	dq := DQ{container: treeset.NewWith(comparator)}
	dq.iterrator = dq.container.Iterator()
	return dq
}

// Pop returns a *common.IOrder and delete it in the queue.
func (q *DQ) Pop() (common.IOrder, bool) {
	if q.iterrator.First() {
		res := q.iterrator.Value().(common.IOrder)
		q.container.Remove(res)
		return res, true
	}
	return nil, false
}

// Push Adds a common.IOrder to the queue.
func (q *DQ) Push(order common.IOrder) {
	q.container.Add(order)
}

// Exist returns true if common.IOrder exist in the queue.
func (q *DQ) Exist(order common.IOrder) bool {
	return q.container.Contains(order)
}

// Del returns and delete final status common.IOrder if common.IOrder exist in the queue.
func (q *DQ) Del(order common.IOrder) {
	q.container.Remove(order)
}

// BuyerOrderComparator is a BuyerQueue comparator, sorted from high to low.
var BuyerOrderComparator utils.Comparator = func(a, b interface{}) int {
	l := a.(common.IOrder)
	r := b.(common.IOrder)
	lp := l.Price()
	rp := r.Price()
	lt := l.Time()
	rt := r.Time()
	c := 0

	if lp > rp {
		c = -1
	} else if lp < rp {
		c = 1
	} else if lt < rt {
		c = -1
	} else if lt > rt {
		c = 1
	}

	return c
}

// SellerOrderComparator is a SellerQueue comparator, sorted from low to high.
var SellerOrderComparator utils.Comparator = func(a, b interface{}) int {
	l := a.(common.IOrder)
	r := b.(common.IOrder)
	lp := l.Price()
	rp := r.Price()
	lt := l.Time()
	rt := r.Time()
	c := 0

	if lp > rp {
		c = 1
	} else if lp < rp {
		c = -1
	} else if lt < rt {
		c = -1
	} else if lt > rt {
		c = 1
	}

	return c
}

// BuyerDQ is the default implementation of IBuyerQueue.
type BuyerDQ struct {
	q            DQ
	highestOrder common.IOrder
	isEmpty      bool
}

// CreateBuyerDQ returns default implemented order queue *BuyerDQ.
func CreateBuyerDQ() *BuyerDQ {
	return &BuyerDQ{
		q:            createDQ(BuyerOrderComparator),
		highestOrder: nil,
		isEmpty:      true,
	}
}

// Pop returns a *common.IOrder and delete it in the queue.
func (b *BuyerDQ) Pop() (common.IOrder, bool) {
	return b.Highest()
}

// Push Adds a common.IOrder to the queue.
func (b *BuyerDQ) Push(order common.IOrder) {
	if b.isEmpty {
		b.highestOrder = order
		b.isEmpty = false
	} else if BuyerOrderComparator(b.highestOrder, order) > 0 {
		b.q.Push(b.highestOrder)
		b.highestOrder = order
	} else {
		b.q.Push(order)
	}
}

// Exist returns true if common.IOrder exist in the queue.
func (b *BuyerDQ) Exist(order common.IOrder) bool {
	if b.isEmpty {
		return false
	}

	if BuyerOrderComparator(b.highestOrder, order) == 0 {
		return true
	}

	return b.q.Exist(order)
}

// Del returns and delete final status common.IOrder if common.IOrder exist in the queue.
func (b *BuyerDQ) Del(order common.IOrder) bool {
	if b.isEmpty {
		return false
	}
	if BuyerOrderComparator(b.highestOrder, order) == 0 {
		b.Deal()
	} else {
		b.q.Del(order)
	}
	return true
}

// Deal deletes the latest Pop order in the queue.
func (b *BuyerDQ) Deal() {
	o, ok := b.q.Pop()
	if ok {
		b.highestOrder = o
	} else {
		b.highestOrder = nil
		b.isEmpty = true
	}
}

// Highest returns a common.IOrder with the highest price but does not delete it in the queue.
func (b *BuyerDQ) Highest() (common.IOrder, bool) {
	if b.isEmpty {
		return nil, false
	}
	return b.highestOrder, true
}

// HighestPrice returns the highest price in the IOrder queue.
func (b *BuyerDQ) HighestPrice() uint64 {
	if b.isEmpty {
		return 0
	}
	return b.highestOrder.Price()
}

// SellerDQ is the default implementation of ISellerQueue.
type SellerDQ struct {
	q           DQ
	lowestOrder common.IOrder
	isEmpty     bool
}

// CreateSellerDQ returns default implemented order queue *BuyerDQ.
func CreateSellerDQ() *SellerDQ {
	return &SellerDQ{
		q:           createDQ(SellerOrderComparator),
		lowestOrder: nil,
		isEmpty:     true,
	}
}

// Pop returns a *common.IOrder and delete it in the queue.
func (s *SellerDQ) Pop() (common.IOrder, bool) {
	return s.Lowest()
}

// Push Adds a common.IOrder to the queue.
func (s *SellerDQ) Push(order common.IOrder) {
	if s.isEmpty {
		s.lowestOrder = order
		s.isEmpty = false
	} else if SellerOrderComparator(s.lowestOrder, order) > 0 {
		s.q.Push(s.lowestOrder)
		s.lowestOrder = order
	} else {
		s.q.Push(order)
	}
}

// Exist returns true if common.IOrder exist in the queue.
func (s *SellerDQ) Exist(order common.IOrder) bool {
	if s.isEmpty {
		return false
	}

	if SellerOrderComparator(s.lowestOrder, order) == 0 {
		return true
	}

	return s.q.Exist(order)
}

// Del delete final status common.IOrder if common.IOrder exist in the queue.
func (s *SellerDQ) Del(order common.IOrder) bool {
	if s.isEmpty {
		return false
	}
	if BuyerOrderComparator(s.lowestOrder, order) == 0 {
		s.Deal()
	} else {
		s.q.Del(order)
	}
	return true
}

// Deal deletes the latest Pop order in the queue.
func (s *SellerDQ) Deal() {
	o, ok := s.q.Pop()
	if ok {
		s.lowestOrder = o
	} else {
		s.lowestOrder = nil
		s.isEmpty = true
	}
}

// Lowest returns a common.IOrder with the Lowest price but does not delete it in the queue.
func (s *SellerDQ) Lowest() (common.IOrder, bool) {
	if s.isEmpty {
		return nil, false
	}
	return s.lowestOrder, true
}

// LowestPrice returns the lowest price in the IOrder queue.
func (s *SellerDQ) LowestPrice() uint64 {
	if s.isEmpty {
		return 0
	}
	return s.lowestOrder.Price()
}
