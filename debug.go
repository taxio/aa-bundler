package main

import "github.com/ethereum/go-ethereum/common"

type Debug struct{}

func (d *Debug) ClearState() (string, error) {
	panic("Not implemented")
}

func (d *Debug) DumpMempool(entryPoint common.Address) ([]UserOperation, error) {
	panic("Not implemented")
}

func (d *Debug) SendBundleNow() (common.Hash, error) {
	panic("Not implemented")
}

func (d *Debug) SetBundlingMode(mode BundlingMode) (string, error) {
	panic("Not implemented")
}

type Reputation struct {
	Address     common.Address  `json:"address"`
	OpsSeen     uint            `json:"opsSeen"`
	OpsIncluded uint            `json:"opsIncluded"`
	Status      PaymasterStatus `json:"status"`
}

func (d *Debug) SetReputation(reputations []Reputation) (string, error) {
	panic("Not implemented")
}

func (d *Debug) DumpReputation() ([]Reputation, error) {
	panic("Not implemented")
}
