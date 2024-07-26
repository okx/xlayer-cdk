package common

import (
	"encoding/binary"
	"math/big"

	"github.com/0xPolygon/cdk/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/iden3/go-iden3-crypto/keccak256"
)

// BlockNum2Bytes converts a block number to a byte slice
func BlockNum2Bytes(blockNum uint64) []byte {
	key := make([]byte, 8)
	binary.LittleEndian.PutUint64(key, blockNum)

	return key
}

// Bytes2BlockNum converts a byte slice to a block number
func Bytes2BlockNum(key []byte) uint64 {
	return binary.LittleEndian.Uint64(key)
}

func CalculateAccInputHash(
	oldAccInputHash common.Hash,
	batchData []byte,
	l1InfoRoot common.Hash,
	timestampLimit uint64,
	sequencerAddr common.Address,
	forcedBlockhashL1 common.Hash) (common.Hash, error) {
	v1 := oldAccInputHash.Bytes()
	v2 := batchData
	v3 := l1InfoRoot.Bytes()
	v4 := big.NewInt(0).SetUint64(timestampLimit).Bytes()
	v5 := sequencerAddr.Bytes()
	v6 := forcedBlockhashL1.Bytes()

	// Add 0s to make values 32 bytes long
	for len(v1) < 32 {
		v1 = append([]byte{0}, v1...)
	}
	for len(v3) < 32 {
		v3 = append([]byte{0}, v3...)
	}
	for len(v4) < 8 {
		v4 = append([]byte{0}, v4...)
	}
	for len(v5) < 20 {
		v5 = append([]byte{0}, v5...)
	}
	for len(v6) < 32 {
		v6 = append([]byte{0}, v6...)
	}

	v2 = keccak256.Hash(v2)

	log.Debugf("OldAccInputHash: %v", oldAccInputHash)
	log.Debugf("BatchHashData: %v", common.Bytes2Hex(v2))
	log.Debugf("L1InfoRoot: %v", l1InfoRoot)
	log.Debugf("TimeStampLimit: %v", timestampLimit)
	log.Debugf("Sequencer Address: %v", sequencerAddr)
	log.Debugf("Forced BlockHashL1: %v", forcedBlockhashL1)

	return common.BytesToHash(keccak256.Hash(v1, v2, v3, v4, v5, v6)), nil
}