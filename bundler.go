package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Bundler struct {
	chainId common.Address
}

func (b *Bundler) SendUserOperation(userOp UserOperation, entryPoint common.Address) (common.Hash, error) {
	panic("Not implemented")
}

type EstimateUserOperationGasResult struct {
	PreVerificationGas big.Int `json:"preVerificationGas"`
	VerificationGas    big.Int `json:"verificationGas"`
	CallGasLimit       big.Int `json:"callGasLimit"`
}

func (b *Bundler) EstimateUserOperationGas(userOp UserOperation, entryPoint common.Address) (EstimateUserOperationGasResult, error) {
	panic("Not implemented")
}

type UserOperationByHashResult struct {
	UserOperation

	EntryPoint      common.Address `json:"entryPoint"`
	TransactionHash common.Hash    `json:"transactionHash"`
	BlockNumber     int            `json:"blockNumber"`
	BlockHash       common.Hash    `json:"blockHash"`
}

func (b *Bundler) GetUserOperationByHash(userOpHash common.Hash) (*UserOperationByHashResult, error) {
	panic("Not implemented")
}

type UserOperationReceipt struct {
	UserOpHash    common.Hash     `json:"userOpHash"`
	Sender        common.Address  `json:"sender"`
	Nonce         big.Int         `json:"nonce"`
	Paymaster     *common.Address `json:"paymaster"`
	ActualGasCost big.Int         `json:"actualGasCost"`
	ActualGasUsed big.Int         `json:"actualGasUsed"`
	Success       bool            `json:"success"`
	Reason        *string         `json:"reason"`
	Logs          []string        `json:"logs"`
	Receipt       types.Receipt   `json:"receipt"`
}

func (b *Bundler) GetUserOperationReceipt(userOpHash common.Hash) (*UserOperationReceipt, error) {
	panic("Not implemented")
}

func (b *Bundler) SupportedEntryPoints() ([]common.Address, error) {
	panic("Not implemented")
}

func (b *Bundler) ChainId() (common.Address, error) {
	return b.chainId, nil
}
