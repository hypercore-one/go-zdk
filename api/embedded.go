package api

import (
	"github.com/hypercore-one/go-zdk/api/embedded"
	"github.com/hypercore-one/go-zdk/client"
)

type EmbeddedApi struct {
	Accelerator embedded.AcceleratorApi
	Htlc        embedded.HtlcApi
	Pillar      embedded.PillarApi
	Plasma      embedded.PlasmaApi
	Sentinel    embedded.SentinelApi
	Spork       embedded.SporkApi
	Stake       embedded.StakeApi
	Swap        embedded.SwapApi
	Token       embedded.TokenApi
}

func NewEmbeddedApi(c client.Client) EmbeddedApi {
	return EmbeddedApi{
		Accelerator: embedded.NewAcceleratorApi(c),
		Htlc:        embedded.NewHtlcApi(c),
		Pillar:      embedded.NewPillarApi(c),
		Plasma:      embedded.NewPlasmaApi(c),
		Sentinel:    embedded.NewSentinelApi(c),
		Spork:       embedded.NewSporkApi(c),
		Stake:       embedded.NewStakeApi(c),
		Swap:        embedded.NewSwapApi(c),
		Token:       embedded.NewTokenApi(c),
	}
}
