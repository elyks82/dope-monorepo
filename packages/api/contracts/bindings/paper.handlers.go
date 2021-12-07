// Code generated by github.com/withtally/ethgen, DO NOT EDIT.

package bindings

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type PaperProcessor interface {
	Setup(address common.Address, eth interface {
		ethereum.ChainReader
		bind.ContractBackend
	}) error
	Initialize(ctx context.Context, start uint64, emit func(string, []interface{})) error

	ProcessApproval(ctx context.Context, e *PaperApproval, emit func(string, []interface{})) error

	ProcessDelegateChanged(ctx context.Context, e *PaperDelegateChanged, emit func(string, []interface{})) error

	ProcessDelegateVotesChanged(ctx context.Context, e *PaperDelegateVotesChanged, emit func(string, []interface{})) error

	ProcessOwnershipTransferred(ctx context.Context, e *PaperOwnershipTransferred, emit func(string, []interface{})) error

	ProcessSnapshot(ctx context.Context, e *PaperSnapshot, emit func(string, []interface{})) error

	ProcessTransfer(ctx context.Context, e *PaperTransfer, emit func(string, []interface{})) error

	mustEmbedUnimplementedPaperProcessor()
}

type UnimplementedPaperProcessor struct {
	Address  common.Address
	ABI      abi.ABI
	Contract *Paper
	Eth      interface {
		ethereum.ChainReader
		bind.ContractBackend
	}
}

func (h *UnimplementedPaperProcessor) Setup(address common.Address, eth interface {
	ethereum.ChainReader
	bind.ContractBackend
}) error {
	contract, err := NewPaper(address, eth)
	if err != nil {
		return fmt.Errorf("new Paper: %w", err)
	}

	abi, err := abi.JSON(strings.NewReader(string(PaperABI)))
	if err != nil {
		return fmt.Errorf("parsing Paper abi: %w", err)
	}

	h.Address = address
	h.ABI = abi
	h.Contract = contract
	h.Eth = eth
	return nil
}

func (h *UnimplementedPaperProcessor) ProcessElement(p interface{}) func(context.Context, types.Log, func(string, []interface{})) error {
	return func(ctx context.Context, vLog types.Log, emit func(string, []interface{})) error {
		switch vLog.Topics[0].Hex() {

		case h.ABI.Events["Approval"].ID.Hex():
			e := new(PaperApproval)
			if err := h.UnpackLog(e, "Approval", vLog); err != nil {
				return fmt.Errorf("unpacking Approval: %w", err)
			}

			e.Raw = vLog
			if err := p.(PaperProcessor).ProcessApproval(ctx, e, emit); err != nil {
				return fmt.Errorf("processing Approval: %w", err)
			}

		case h.ABI.Events["DelegateChanged"].ID.Hex():
			e := new(PaperDelegateChanged)
			if err := h.UnpackLog(e, "DelegateChanged", vLog); err != nil {
				return fmt.Errorf("unpacking DelegateChanged: %w", err)
			}

			e.Raw = vLog
			if err := p.(PaperProcessor).ProcessDelegateChanged(ctx, e, emit); err != nil {
				return fmt.Errorf("processing DelegateChanged: %w", err)
			}

		case h.ABI.Events["DelegateVotesChanged"].ID.Hex():
			e := new(PaperDelegateVotesChanged)
			if err := h.UnpackLog(e, "DelegateVotesChanged", vLog); err != nil {
				return fmt.Errorf("unpacking DelegateVotesChanged: %w", err)
			}

			e.Raw = vLog
			if err := p.(PaperProcessor).ProcessDelegateVotesChanged(ctx, e, emit); err != nil {
				return fmt.Errorf("processing DelegateVotesChanged: %w", err)
			}

		case h.ABI.Events["OwnershipTransferred"].ID.Hex():
			e := new(PaperOwnershipTransferred)
			if err := h.UnpackLog(e, "OwnershipTransferred", vLog); err != nil {
				return fmt.Errorf("unpacking OwnershipTransferred: %w", err)
			}

			e.Raw = vLog
			if err := p.(PaperProcessor).ProcessOwnershipTransferred(ctx, e, emit); err != nil {
				return fmt.Errorf("processing OwnershipTransferred: %w", err)
			}

		case h.ABI.Events["Snapshot"].ID.Hex():
			e := new(PaperSnapshot)
			if err := h.UnpackLog(e, "Snapshot", vLog); err != nil {
				return fmt.Errorf("unpacking Snapshot: %w", err)
			}

			e.Raw = vLog
			if err := p.(PaperProcessor).ProcessSnapshot(ctx, e, emit); err != nil {
				return fmt.Errorf("processing Snapshot: %w", err)
			}

		case h.ABI.Events["Transfer"].ID.Hex():
			e := new(PaperTransfer)
			if err := h.UnpackLog(e, "Transfer", vLog); err != nil {
				return fmt.Errorf("unpacking Transfer: %w", err)
			}

			e.Raw = vLog
			if err := p.(PaperProcessor).ProcessTransfer(ctx, e, emit); err != nil {
				return fmt.Errorf("processing Transfer: %w", err)
			}

		}
		return nil
	}
}

func (h *UnimplementedPaperProcessor) UnpackLog(out interface{}, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := h.ABI.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range h.ABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func (h *UnimplementedPaperProcessor) Initialize(ctx context.Context, start uint64, emit func(string, []interface{})) error {
	return nil
}

func (h *UnimplementedPaperProcessor) ProcessApproval(ctx context.Context, e *PaperApproval, emit func(string, []interface{})) error {
	return nil
}

func (h *UnimplementedPaperProcessor) ProcessDelegateChanged(ctx context.Context, e *PaperDelegateChanged, emit func(string, []interface{})) error {
	return nil
}

func (h *UnimplementedPaperProcessor) ProcessDelegateVotesChanged(ctx context.Context, e *PaperDelegateVotesChanged, emit func(string, []interface{})) error {
	return nil
}

func (h *UnimplementedPaperProcessor) ProcessOwnershipTransferred(ctx context.Context, e *PaperOwnershipTransferred, emit func(string, []interface{})) error {
	return nil
}

func (h *UnimplementedPaperProcessor) ProcessSnapshot(ctx context.Context, e *PaperSnapshot, emit func(string, []interface{})) error {
	return nil
}

func (h *UnimplementedPaperProcessor) ProcessTransfer(ctx context.Context, e *PaperTransfer, emit func(string, []interface{})) error {
	return nil
}

func (h *UnimplementedPaperProcessor) mustEmbedUnimplementedPaperProcessor() {}