package growler

import (
	"github.com/growler/common"
	"github.com/growler/processer"
	"github.com/growler/queue"
)

/**********************************************************/
/**********************ENGINE CONFIG***********************/
/**********************************************************/

func GetConfig() *Config {
	return &Config{
		Matchers: []MatcherConfig{
			{
				ID:                0,
				EnableOfferPrice:  false,
				Pairs:             []common.Pair{"BTC/USDT", "ETH/USDT"},
				ChannelBufferSize: 2000,
				BuyerQueues:       nil,
				SellerQueues:      nil,
			},
			{
				ID:                1,
				EnableOfferPrice:  true,
				Pairs:             []common.Pair{"BSV/USDT", "BCH/USDT"},
				ChannelBufferSize: 1000,
				BuyerQueues:       nil,
				SellerQueues:      nil,
			},
		},
		Dispatcher: DispatcherConfig{
			OrderComeInChannelBufferSize:  5000,
			OrderComeOutChannelBufferSize: 2000,
			OrderOnHandChannelBufferSize:  2000,
			Selector:                      nil,
			PrePrecesser: PrePrecesserConfig{
				Middlewares: []processer.HandlerFunc{},
			},
			PostPrecesser: PostPrecesserConfig{
				Middlewares: []processer.HandlerFunc{},
			},
			Adaptor: nil,
		},
	}
}

type Config struct {
	Matchers   []MatcherConfig
	Dispatcher DispatcherConfig
}

type MatcherConfig struct {
	ID                int
	EnableOfferPrice  bool
	Pairs             []common.Pair
	ChannelBufferSize int
	BuyerQueues       map[common.Pair]queue.IBuyerQueue
	SellerQueues      map[common.Pair]queue.ISellerQueue
}

type PrePrecesserConfig struct {
	Middlewares []processer.HandlerFunc
}

type PostPrecesserConfig struct {
	Middlewares []processer.HandlerFunc
}

type DispatcherConfig struct {
	OrderComeInChannelBufferSize  int
	OrderComeOutChannelBufferSize int
	OrderOnHandChannelBufferSize  int
	Selector                      MatcherSelector
	PrePrecesser                  PrePrecesserConfig
	PostPrecesser                 PostPrecesserConfig
	Adaptor                       common.IEventAdaptor
}
