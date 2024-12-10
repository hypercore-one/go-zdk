package utils

import (
	"github.com/hypercore-one/go-zdk/client"
	"github.com/hypercore-one/go-zdk/zdk"
)

func NewClientFromMeta(url string) (client.Client, error) {
	c, err := client.NewClient(url)
	if err != nil {
		return nil, err
	}
	z := zdk.NewZdk(c)
	meta, err := z.Ledger.Meta()
	if err != nil {
		// TODO if ledger.meta method doesn't exist, just default to mainnet config
		return nil, err
	}
	nc, err := client.NewClient(url, client.ChainIdentifier(meta.ChainId), client.ZToken(meta.ZToken), client.QToken(meta.QToken))
	if err != nil {
		return nil, err
	}
	return nc, nil
}
