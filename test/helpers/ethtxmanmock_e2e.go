package helpers

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xPolygon/cdk/log"
	"github.com/0xPolygonHermez/zkevm-ethtx-manager/ethtxmanager"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/mock"
)

func NewEthTxManMock(
	t *testing.T,
	client *simulated.Backend,
	auth *bind.TransactOpts,
) *EthTxManagerMock {
	ethTxMock := NewEthTxManagerMock(t)
	ethTxMock.On("Add", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			ctx := context.Background()
			nonce, err := client.Client().PendingNonceAt(ctx, auth.From)
			if err != nil {
				log.Error(err)
				return
			}
			gas, err := client.Client().EstimateGas(ctx, ethereum.CallMsg{
				From:  auth.From,
				To:    args.Get(1).(*common.Address),
				Value: big.NewInt(0),
				Data:  args.Get(4).([]byte),
			})
			if err != nil {
				log.Error(err)
				res, err := client.Client().CallContract(ctx, ethereum.CallMsg{
					From:  auth.From,
					To:    args.Get(1).(*common.Address),
					Value: big.NewInt(0),
					Data:  args.Get(4).([]byte),
				}, nil)
				log.Debugf("contract call: %s", res)
				if err != nil {
					log.Errorf("%+v", err)
				}
				return
			}
			price, err := client.Client().SuggestGasPrice(ctx)
			if err != nil {
				log.Error(err)
			}
			tx := types.NewTx(&types.LegacyTx{
				To:       args.Get(1).(*common.Address),
				Nonce:    nonce,
				Value:    big.NewInt(0),
				Data:     args.Get(4).([]byte),
				Gas:      gas,
				GasPrice: price,
			})
			tx.Gas()
			signedTx, err := auth.Signer(auth.From, tx)
			if err != nil {
				log.Error(err)
				return
			}
			err = client.Client().SendTransaction(ctx, signedTx)
			if err != nil {
				log.Error(err)
				return
			}
			client.Commit()
		}).
		Return(common.Hash{}, nil)
	// res, err := c.ethTxMan.Result(ctx, id)
	ethTxMock.On("Result", mock.Anything, mock.Anything).
		Return(ethtxmanager.MonitoredTxResult{Status: ethtxmanager.MonitoredTxStatusMined}, nil)

	return ethTxMock
}