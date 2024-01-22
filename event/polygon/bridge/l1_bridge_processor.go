package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/polygon/utils"
)

func L1Deposit(polygonBridge *abi.Polygonzkevmbridge, event event.ContractEvent) (*worker.L1ToL2, error) {
	rlpLog := event.RLPLog
	d, unpackErr := polygonBridge.ParseBridgeEvent(*rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}
	l1ToL2 := worker.L1ToL2{
		QueueIndex:        nil,
		L1TransactionHash: rlpLog.TxHash,
		L2TransactionHash: common.Hash{},
		L1TxOrigin:        common.Address{},
		Status:            0,
		FromAddress:       d.DestinationAddress,
		ToAddress:         d.DestinationAddress,
		L1TokenAddress:    d.OriginAddress,
		L2TokenAddress:    common.Address{},
		ETHAmount:         big.NewInt(0),
		GasLimit:          big.NewInt(0),
		Timestamp:         int64(event.Timestamp),
		AssetType:         int64(common2.ETH),
		TokenAmounts:      "0",
		MessageHash:       common.Hash{},
	}
	if d.OriginAddress.String() == utils.L1_ETH {
		l1ToL2.ETHAmount = d.Amount
		l1ToL2.AssetType = int64(common2.ETH)
	} else {
		l1ToL2.TokenAmounts = d.Amount.String()
		l1ToL2.AssetType = int64(common2.ERC20)
	}

	return &l1ToL2, nil
}

func L1WithdrawClaimed(chainId string, polygonBridge *abi.Polygonzkevmbridge, event event.ContractEvent, db *database.DB) (*worker.L1ToL2, error) {

	rlpLog := event.RLPLog
	c, unpackErr := polygonBridge.ParseClaimEvent(*rlpLog)
	if unpackErr != nil {
		return nil, unpackErr
	}

	l2ToL1 := worker.L2ToL1{
		L1FinalizeTxHash: event.TransactionHash,
		Status:           1,
		ClaimedIndex:     c.GlobalIndex.Int64(),
	}

	err := db.L2ToL1.UpdateL2ToL1ClaimedStatus(chainId, l2ToL1)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
