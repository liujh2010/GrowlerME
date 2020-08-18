package growler

import (
	"fmt"
	"testing"
	"time"

	"github.com/growler/common"
	"github.com/growler/processer"
	"github.com/growler/queue"
)

func TestConfiging(t *testing.T) {
	d := &Dispatcher{}
	d.configing(GetConfig())
}

func TestMatching(t *testing.T) {
	tests := []*common.DefaultOrder{
		common.CreateDefaultOrder(0, "ABC", true, true, common.MATCH, 90000000, 120000),
		common.CreateDefaultOrder(1, "ABC", true, true, common.MATCH, 90000000, 120000),
		common.CreateDefaultOrder(2, "ABC", false, true, common.MATCH, 90000000, 90000),
		common.CreateDefaultOrder(3, "ABC", true, true, common.MATCH, 90000000, 120000),
		common.CreateDefaultOrder(4, "ABC", false, true, common.MATCH, 90000000, 200000),
		common.CreateDefaultOrder(5, "ABC", true, true, common.MATCH, 90000000, 120000), // 90000
	}

	d := CreateDispatcher(getTestEngineConfig())
	d.Work()

	go func(d *Dispatcher) {
		for _, o := range tests {
			d.in <- o
		}
	}(d)

	for {
		select {
		case o := <-d.out:
			fmt.Printf("get order %v", o)
		case <-time.After(time.Second * 1):
			return
		}
	}
}

func BenchmarkAddOrder(t *testing.B) {
	d := CreateDispatcher(getTestEngineConfig())
	d.Work()

	var last *common.DefaultOrder
	for i := 0; i < t.N; i++ {
		last = common.CreateDefaultOrder(uint64(i), "ABC", true, true, common.MATCH, 90000000, 120000)
		d.in <- last
	}

	for !d.matchers[0].Exist(last) {
	}
}

func BenchAddOrder() {
	BenchmarkAddOrder(&testing.B{})
}

type TestAdaptor struct{}

func (a *TestAdaptor) EventToIOrder(e interface{}) common.IOrder {
	return e.(common.IOrder)
}

func (a *TestAdaptor) IOrderToEvent(o common.IOrder) interface{} {
	return o
}

func getTestEngineConfig() *Config {
	conf := GetConfig()
	buyqueues := make(map[common.Pair]queue.IBuyerQueue, 1)
	buyqueues["ABC"] = queue.CreateBuyerDQ()
	sellqueues := make(map[common.Pair]queue.ISellerQueue, 1)
	sellqueues["ABC"] = queue.CreateSellerDQ()
	conf.Matchers = []MatcherConfig{
		{
			ID:                0,
			EnableOfferPrice:  true,
			Pairs:             []common.Pair{"ABC"},
			ChannelBufferSize: 2000,
			BuyerQueues:       buyqueues,
			SellerQueues:      sellqueues,
		},
	}

	conf.Dispatcher.PostPrecesser.Middlewares = []processer.HandlerFunc{func(c *processer.Context) {}}
	conf.Dispatcher.PrePrecesser.Middlewares = []processer.HandlerFunc{func(c *processer.Context) {}}
	conf.Dispatcher.Selector = func(order common.IOrder) int { return 0 }
	conf.Dispatcher.Adaptor = &TestAdaptor{}
	return conf
}

func main() {
	BenchAddOrder()
}
