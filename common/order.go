package common

import "time"

/*******************************************/
/********Order interface Definition*********/
/*******************************************/

// Pair is to identify each trading pair.
type Pair string

// OrderAction identifies the behavior of an order in the engine that will be executed.
// # MATCH identifies this order to match.
// # CANCEL indicates that this order will be cancelled.
type OrderAction int8

const (
	// MATCH identifies this order to match.
	MATCH OrderAction = 0
	// CANCEL indicates that this order will be cancelled.
	CANCEL OrderAction = 1
	// // ROBOT_MATCH identifies that this order will be robot-specific matching.
	// ROBOT_MATCH OrderAction = 2
	// // ROBOT_CANCEL indicates that this order will be robot-specific cancelled.
	// ROBOT_CANCEL OrderAction = 3
)

// IOrder defines the interface that must be implemented to match orders in the engine.
type IOrder interface {
	Take(match IOrder, strikePrice uint64)
	Pair() Pair
	Price() uint64
	Time() uint64
	CloseOrder()
	CancelOrder()
	IsBuyer() bool
	IsLimit() bool
	IsCancel() bool
	IsClose() bool
	Action() OrderAction
}

/*******************************************/
/*********Default Order Definition**********/
/*******************************************/

// DefaultOrder save all information for a single trade order.
type DefaultOrder struct {
	ID              uint64
	pair            Pair
	isBuyer         bool
	isLimit         bool
	isClose         bool
	isCancel        bool
	action          OrderAction
	price           uint64
	amount          uint64
	RemainingAmount uint64
	Turnover        uint64
}

// CreateDefaultOrder returns new DefaultOrder point.
func CreateDefaultOrder(id uint64, pair Pair, isBuyer, isLimit bool, action OrderAction, price, amount uint64) *DefaultOrder {
	return &DefaultOrder{
		ID:              id,
		pair:            pair,
		isBuyer:         isBuyer,
		isLimit:         isLimit,
		isClose:         false,
		isCancel:        false,
		action:          action,
		price:           price,
		amount:          amount,
		RemainingAmount: amount,
		Turnover:        0,
	}
}

// Take trades the called order with the matched order.
func (o *DefaultOrder) Take(match IOrder, strikePrice uint64) {
	m := match.(*DefaultOrder)

	if !o.isLimit && o.isBuyer {
		o.takeBuyerMarketOrder(m)
		return
	}

	var turnover uint64
	var tradeAmount uint64
	if m.RemainingAmount > o.RemainingAmount {
		turnover = strikePrice * o.RemainingAmount
		tradeAmount = o.RemainingAmount

		m.RemainingAmount -= o.RemainingAmount
		o.RemainingAmount = 0

		o.isClose = true
	} else if m.RemainingAmount < o.RemainingAmount {
		turnover = strikePrice * m.RemainingAmount
		tradeAmount = m.RemainingAmount

		o.RemainingAmount -= m.RemainingAmount
		m.RemainingAmount = 0

		m.isClose = true
	} else {
		turnover = strikePrice * m.RemainingAmount
		tradeAmount = m.RemainingAmount

		o.RemainingAmount = 0
		m.RemainingAmount = 0

		o.isClose = true
		m.isClose = true
	}

	o.Turnover += turnover
	m.Turnover += turnover

	res := new(MatchingRes)
	res.Pair = o.pair
	res.Price = m.price
	res.Amount = tradeAmount
	res.Turnover = turnover
	res.IsBuyer = o.isBuyer
	if o.isBuyer {
		res.BuyerOrderID = o.ID
		res.SellerOrderID = m.ID
	} else {
		res.BuyerOrderID = m.ID
		res.SellerOrderID = o.ID
	}
	res.Time = time.Now().Unix()
}

func (o *DefaultOrder) takeBuyerMarketOrder(match *DefaultOrder) (*MatchingRes, bool) {
	realAmount := o.RemainingAmount / match.price
	if realAmount == 0 {
		o.isClose = true
		o.isCancel = true
		return nil, false
	}

	var turnover uint64
	var tradeAmount uint64
	if realAmount > match.RemainingAmount {
		turnover = match.price * match.RemainingAmount
		tradeAmount = match.RemainingAmount

		match.RemainingAmount = 0
		match.isClose = true

		o.RemainingAmount -= turnover
	} else if realAmount < match.RemainingAmount {
		turnover = match.price * realAmount
		tradeAmount = realAmount

		match.RemainingAmount -= realAmount

		o.RemainingAmount = 0
		o.isClose = true
	} else {
		turnover = match.price * realAmount
		tradeAmount = realAmount

		match.RemainingAmount = 0
		match.isClose = true

		o.RemainingAmount = 0
		o.isClose = true
	}

	match.Turnover += turnover
	o.Turnover += turnover

	res := new(MatchingRes)
	res.Pair = o.pair
	res.Price = match.price
	res.Amount = tradeAmount
	res.Turnover = turnover
	res.IsBuyer = o.isBuyer
	if o.isBuyer {
		res.BuyerOrderID = o.ID
		res.SellerOrderID = match.ID
	} else {
		res.BuyerOrderID = match.ID
		res.SellerOrderID = o.ID
	}
	res.Time = time.Now().Unix()
	return res, true
}

func (o *DefaultOrder) Pair() Pair {
	return o.pair
}

func (o *DefaultOrder) Price() uint64 {
	return o.price
}

func (o *DefaultOrder) Time() uint64 {
	return o.ID
}

func (o *DefaultOrder) CloseOrder() {
	o.isClose = true
}

func (o *DefaultOrder) CancelOrder() {
	o.isCancel = true
}

func (o *DefaultOrder) IsBuyer() bool {
	return o.isBuyer
}

func (o *DefaultOrder) IsLimit() bool {
	return o.isLimit
}

func (o *DefaultOrder) IsCancel() bool {
	return o.isCancel
}

func (o *DefaultOrder) IsClose() bool {
	return o.isClose
}

func (o *DefaultOrder) Action() OrderAction {
	return o.action
}

/*******************************************/
/*****Default Mtaching Result Definition****/
/*******************************************/

// MatchingRes save the result of each match.
type MatchingRes struct {
	Pair          Pair
	Price         uint64
	Amount        uint64
	Turnover      uint64
	IsBuyer       bool
	BuyerOrderID  uint64
	SellerOrderID uint64
	Time          int64
}
