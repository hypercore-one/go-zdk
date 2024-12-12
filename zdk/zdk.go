package zdk

import (
	"github.com/hypercore-one/go-zdk/api"
	"github.com/hypercore-one/go-zdk/client"
)

type Zdk struct {
	client.Client
	Ledger    api.LedgerApi
	Stats     api.StatsApi
	Embedded  api.EmbeddedApi
	Subscribe api.SubscribeApi
}

func NewZdk(client client.Client) *Zdk {
	return &Zdk{
		Client:    client,
		Ledger:    api.NewLedgerApi(client),
		Stats:     api.NewStatsApi(client),
		Embedded:  api.NewEmbeddedApi(client),
		Subscribe: api.NewSubscribeApi(client),
	}
}
