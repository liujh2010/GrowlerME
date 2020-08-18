package queue

import "github.com/growler/common"

// IQueue is basic Order container interface for matcher.
type IQueue interface {
	// Pop returns a *common.IOrder and delete it in the queue.
	Pop() (common.IOrder, bool)
	// Push Adds a common.IOrder to the queue.
	Push(common.IOrder)
	// Exist returns true if common.IOrder exist in the queue.
	Exist(common.IOrder) bool
	// Del delete final status common.IOrder if common.IOrder exist in the queue.
	Del(common.IOrder) bool
}

// IOrderQueue defines the functions required by the order queue.
type IOrderQueue interface {
	IQueue
	// Deal deletes the latest Pop order in the queue.
	Deal()
}

// IBuyerQueue is buyer Order container interface for matcher.
type IBuyerQueue interface {
	IOrderQueue
	// Highest returns a common.IOrder with the highest price but does not delete it in the queue.
	Highest() (common.IOrder, bool)
	// HighestPrice returns the highest price in the IOrder queue.
	HighestPrice() uint64
}

// ISellerQueue is seller Order container interface for matcher.
type ISellerQueue interface {
	IOrderQueue
	// Lowest returns a common.IOrder with the Lowest price but does not delete it in the queue.
	Lowest() (common.IOrder, bool)
	// LowestPrice returns the lowest price in the IOrder queue.
	LowestPrice() uint64
}
