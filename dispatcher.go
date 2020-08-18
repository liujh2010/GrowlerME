package growler

import (
	"github.com/growler/common"
	"github.com/growler/match"
	"github.com/growler/processer"
)

// MatcherSelector returns an integer of type int, which is the same as Matcher's ID.
// Dispatcher uses this method to select the corresponding Matcher matching order.
type MatcherSelector func(order common.IOrder) int

// Dispatcher managers the flow of data.
type Dispatcher struct {
	isWork        bool
	preprocessor  processer.IProcesser
	postProcessor processer.IProcesser
	adaptor       common.IEventAdaptor
	in            chan interface{}
	outOnHand     chan common.IOrder
	out           chan interface{}
	selector      MatcherSelector
	matchers      map[int]*match.Matcher
}

// Work makes a Dispatcher work
func (d *Dispatcher) Work() {
	d.isWork = true

	// start Matcher.
	for _, matcher := range d.matchers {
		go matcher.Work()
	}

	// After placing the order through the preprocessor and select the correct Matcher to matching.
	go func(d *Dispatcher) {
		for event := range d.in {
			if d.isWork {
				d.preprocessor.Do(event)
				order := d.adaptor.EventToIOrder(event)
				matcherID := d.selector(order)
				d.matchers[matcherID].InChannel() <- order
			}
		}
	}(d)

	// Pass the matched order through the post processor.
	go func(d *Dispatcher) {
		for order := range d.outOnHand {
			event := d.adaptor.IOrderToEvent(order)
			d.postProcessor.Do(event)
			d.out <- event
		}
	}(d)
}

func (d *Dispatcher) configing(config *Config) {
	if len(config.Matchers) == 0 {
		panic("Engine requires at least one Matcher!")
	}

	d.in = make(chan interface{}, config.Dispatcher.OrderComeInChannelBufferSize)
	d.out = make(chan interface{}, config.Dispatcher.OrderComeOutChannelBufferSize)
	d.outOnHand = make(chan common.IOrder, config.Dispatcher.OrderOnHandChannelBufferSize)
	d.matchers = make(map[int]*match.Matcher)
	d.preprocessor = &processer.Processer{}
	d.postProcessor = &processer.Processer{}
	d.selector = config.Dispatcher.Selector
	d.adaptor = config.Dispatcher.Adaptor
	d.isWork = false

	for _, m := range config.Dispatcher.PrePrecesser.Middlewares {
		d.preprocessor.Use(m)
	}

	for _, m := range config.Dispatcher.PostPrecesser.Middlewares {
		d.postProcessor.Use(m)
	}

	for _, conf := range config.Matchers {
		matcher := match.CreateMatcher(conf.ID, conf.BuyerQueues, conf.SellerQueues,
			conf.ChannelBufferSize, d.outOnHand, conf.EnableOfferPrice)
		d.matchers[matcher.ID()] = matcher
	}
}

// CreateDispatcher returns a Dispatcher ready to work
func CreateDispatcher(config *Config) *Dispatcher {
	d := Dispatcher{}
	d.configing(config)
	return &d
}
