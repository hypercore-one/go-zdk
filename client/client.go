package client

import (
	"context"

	"github.com/zenon-network/go-zenon/common/types"
	rpc "github.com/zenon-network/go-zenon/rpc/server"
)

const DefaultUrl = "ws://127.0.0.1:35998"

type client struct {
	protocolVersion uint64
	chainIdentifier uint64
	zToken          types.ZenonTokenStandard
	qToken          types.ZenonTokenStandard
	ec              *rpc.Client
}

type Option func(c *client)

func ProtocolVersion(version uint64) Option {
	return func(c *client) {
		c.protocolVersion = version
	}
}

func ChainIdentifier(chainId uint64) Option {
	return func(c *client) {
		c.chainIdentifier = chainId
	}
}

func ZToken(zts types.ZenonTokenStandard) Option {
	return func(c *client) {
		c.zToken = zts
	}
}

func QToken(zts types.ZenonTokenStandard) Option {
	return func(c *client) {
		c.qToken = zts
	}
}

func NewClient(url string, opts ...Option) (Client, error) {
	c := client{protocolVersion: 1, chainIdentifier: 1, zToken: types.ZnnTokenStandard, qToken: types.QsrTokenStandard}
	for _, opt := range opts {
		opt(&c)
	}
	r, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	c.ec = r
	return &c, nil

}

func (c *client) ProtocolVersion() uint64 {
	return c.protocolVersion
}

func (c *client) ChainIdentifier() uint64 {
	return c.chainIdentifier
}

func (c *client) ZToken() types.ZenonTokenStandard {
	return c.zToken
}

func (c *client) QToken() types.ZenonTokenStandard {
	return c.qToken
}

func (c *client) Call(result interface{}, method string, args ...interface{}) error {
	return c.ec.Call(result, method, args...)
}

func (c *client) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (Subscription, error) {
	return c.ec.Subscribe(ctx, namespace, channel, args...)
}
