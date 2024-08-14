package txbuilder_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/0xPolygon/cdk/dataavailability/mocks_da"
	"github.com/0xPolygon/cdk/sequencesender/seqsendertypes"
	"github.com/0xPolygon/cdk/sequencesender/txbuilder"
	"github.com/0xPolygon/cdk/sequencesender/txbuilder/mocks_txbuilder"
	"github.com/0xPolygon/cdk/state/datastream"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBananaValidiumName(t *testing.T) {
	testData := newBananaValidiumTestData(t, txbuilder.MaxBatchesForL1Disabled)
	require.NotNil(t, testData.sut)
	require.True(t, strings.Contains(testData.sut.String(), "Banana"))
	require.True(t, strings.Contains(testData.sut.String(), "Validium"))
}

func TestBananaValidiumBuildSequenceBatchesTxSequenceErrorsFromDA(t *testing.T) {
	testData := newBananaValidiumTestData(t, txbuilder.MaxBatchesForL1Disabled)
	seq, err := newSequenceBananaValidiumForTest(testData)
	require.NoError(t, err)
	ctx := context.TODO()
	testData.da.EXPECT().PostSequenceBanana(ctx, mock.Anything).Return(nil, nil).Once()

	_, err = testData.sut.BuildSequenceBatchesTx(ctx, seq)
	require.Error(t, err, "data availability message is nil")

	testData.da.EXPECT().PostSequenceBanana(ctx, mock.Anything).Return(nil, fmt.Errorf("test error"))
	_, err = testData.sut.BuildSequenceBatchesTx(ctx, seq)
	require.Error(t, err, "error posting sequences to the data availability protocol: test error")

}

func TestBananaValidiumBuildSequenceBatchesTxSequenceDAOk(t *testing.T) {
	testData := newBananaValidiumTestData(t, txbuilder.MaxBatchesForL1Disabled)
	seq, err := newSequenceBananaValidiumForTest(testData)
	require.NoError(t, err)
	ctx := context.TODO()
	daMessage := []byte{1}
	testData.da.EXPECT().PostSequenceBanana(ctx, mock.Anything).Return(daMessage, nil)
	inner := &ethtypes.LegacyTx{}
	seqBatchesTx := ethtypes.NewTx(inner)
	testData.rollupContract.EXPECT().SequenceBatchesValidium(mock.MatchedBy(func(opts *bind.TransactOpts) bool {
		return opts.NoSend == true
	}), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, daMessage).Return(seqBatchesTx, nil).Once()
	tx, err := testData.sut.BuildSequenceBatchesTx(ctx, seq)
	require.NoError(t, err)
	require.Equal(t, seqBatchesTx, tx)
}

type testDataBananaValidium struct {
	rollupContract *mocks_txbuilder.RollupBananaValidiumContractor
	getContract    *mocks_txbuilder.GlobalExitRootBananaContractor
	cond           *mocks_txbuilder.CondNewSequence
	da             *mocks_da.SequenceSenderBanana
	opts           bind.TransactOpts
	sut            *txbuilder.TxBuilderBananaValidium
}

func newBananaValidiumTestData(t *testing.T, maxBatchesForL1 uint64) *testDataBananaValidium {
	zkevmContractMock := mocks_txbuilder.NewRollupBananaValidiumContractor(t)
	gerContractMock := mocks_txbuilder.NewGlobalExitRootBananaContractor(t)
	condMock := mocks_txbuilder.NewCondNewSequence(t)
	daMock := mocks_da.NewSequenceSenderBanana(t)

	opts := bind.TransactOpts{}
	sut := txbuilder.NewTxBuilderBananaValidium(zkevmContractMock, gerContractMock, daMock, opts, maxBatchesForL1)
	require.NotNil(t, sut)
	sut.SetCondNewSeq(condMock)
	return &testDataBananaValidium{
		rollupContract: zkevmContractMock,
		getContract:    gerContractMock,
		cond:           condMock,
		da:             daMock,
		opts:           opts,
		sut:            sut,
	}
}

func newSequenceBananaValidiumForTest(testData *testDataBananaValidium) (seqsendertypes.Sequence, error) {
	l2Block := &datastream.L2Block{
		Timestamp:       1,
		BatchNumber:     1,
		L1InfotreeIndex: 3,
		Coinbase:        []byte{1, 2, 3},
		GlobalExitRoot:  []byte{4, 5, 6},
	}
	batch := testData.sut.NewBatchFromL2Block(l2Block)
	batches := []seqsendertypes.Batch{
		batch,
	}
	lastAcc := common.HexToHash("0x8aca9664752dbae36135fd0956c956fc4a370feeac67485b49bcd4b99608ae41")
	testData.rollupContract.EXPECT().LastAccInputHash(mock.Anything).Return(lastAcc, nil).Once()
	l1infoRoot := common.HexToHash("0x66ca9664752dbae36135fd0956c956fc4a370feeac67485b49bcd4b99608ae41")
	testData.getContract.EXPECT().L1InfoRootMap(mock.Anything, uint32(3)).Return(l1infoRoot, nil).Once()
	return testData.sut.NewSequence(batches, common.Address{})
}