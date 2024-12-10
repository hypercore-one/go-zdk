package client

import (
	"context"

	"github.com/zenon-network/go-zenon/common/types"
)

type Client interface {
	Call(result interface{}, method string, args ...interface{}) error
	Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (Subscription, error)
	ProtocolVersion() uint64
	ChainIdentifier() uint64
	ZToken() types.ZenonTokenStandard
	QToken() types.ZenonTokenStandard
}

type Subscription interface {
	Err() <-chan error
	Unsubscribe()
}
