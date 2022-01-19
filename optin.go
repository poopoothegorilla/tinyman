package tinyman

import (
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/types"
)

// PrepareAppOptInTxns ...
func PrepareAppOptInTxns(validatorAppID uint64, address string, sp types.SuggestedParams) ([]types.Transaction, error) {
	add, err := types.DecodeAddress(address)
	if err != nil {
		return nil, err
	}

	tx, err := future.MakeApplicationOptInTx(validatorAppID, nil, nil, nil, nil, sp, add, nil, types.Digest{}, [32]byte{}, types.Address{})
	if err != nil {
		return nil, err
	}

	txns := []types.Transaction{tx}
	gid, err := crypto.ComputeGroupID(txns)
	if err != nil {
		return nil, err
	}
	for i, _ := range txns {
		txns[i].Group = gid
	}

	return txns, nil
}

// PrepareAssetOptInTxns ...
func PrepareAssetOptIn(assetID uint64, address string, sp types.SuggestedParams) ([]types.Transaction, error) {
	tx, err := future.MakeAssetAcceptanceTxn(address, nil, sp, assetID)
	if err != nil {
		return nil, err
	}

	txns := []types.Transaction{tx}
	gid, err := crypto.ComputeGroupID(txns)
	if err != nil {
		return nil, err
	}
	for i, tx := range txns {
		txns[i].Group = gid
	}

	return txns, nil
}
