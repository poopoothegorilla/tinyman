package tinyman

import (
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/types"
)

const (
	swapFee = 2000
)

// PrepareSwapTransactions one asset (ASA or Algo) for another with the Pool.
func PrepareSwapTransactions(
	sp types.SuggestedParams,
	validatorAppID uint64,
	asset1ID, asset2ID, liquidityAssetID uint64,
	assetInID uint64,
	assetInAmt, assetOutAmt uint64,
	swapType SwapType,
	senderAdd string) (*TransactionGroup, error) {

	poolLogicSigAccount, err := getPoolLogicSigAccount(validatorAppID, asset1ID, asset2ID)
	if err != nil {
		return nil, err
	}

	poolAddress, err := poolLogicSigAccount.Address()
	if err != nil {
		return nil, err
	}

	swapTypes := map[SwapType]string{
		FixedInput:  "fi",
		FixedOutput: "fo",
	}

	assetOutID := asset1ID
	if assetInID == asset1ID {
		assetOutID = asset2ID
	}

	// PaymentTxn
	// Pay - pay fees in Algo from Swapper to Pool
	tx1, err := future.MakePaymentTxn(senderAdd, poolAddress.String(), swapFee, []byte("fee"), "", sp)
	if err != nil {
		return nil, err
	}

	// ApplicationNoOpTxn
	// App Call - NoOp call to Validator App with args ['swap', (fixed-input) or
	// 'fo' (fixed-output)], with Swapper account
	// Argument 1 'fi' specifies that the input (sell) is fixed but the output may
	// vary with slippage
	// Argument 1 'fo' specifies that the output (buy) is fixed but the input may
	// vary with slippage
	foreignAssets := []uint64{asset1ID}
	if asset2ID != 0 {
		foreignAssets = append(foreignAssets, asset2ID)
	}
	foreignAssets = append(foreignAssets, liquidityAssetID)
	args := [][]byte{
		[]byte("swap"),
		[]byte(swapTypes[swapType]),
	}
	accounts := []string{senderAdd}
	var group types.Digest
	var lease [32]byte
	var rekeyTo types.Address

	tx2, err := future.MakeApplicationNoOpTx(validatorAppID, args, accounts, nil, foreignAssets, sp, poolAddress, nil, group, lease, rekeyTo)
	if err != nil {
		return nil, err
	}

	// AssetTransferTxn
	var tx3 types.Transaction
	if assetInID != 0 {
		// (a) AssetTransfer - Transfer of sell asset from Swapper to Pool
		// if sell asset is an ASA
		tx3, err = future.MakeAssetTransferTxn(senderAdd, poolAddress.String(), assetInAmt, nil, sp, "", assetInID)
	} else {
		// (b) Pay - Transfer of Algo from Swapper to Pool
		// if sell asset is an ALGO
		tx3, err = future.MakePaymentTxn(senderAdd, poolAddress.String(), assetInAmt, nil, "", sp)
	}
	if err != nil {
		return nil, err
	}

	// AssetTransferTxn
	var tx4 types.Transaction
	if assetOutID != 0 {
		// (a) AssetTransfer - Transfer of buy asset from Pool to Swapper
		// if buy asset is an ASA
		tx4, err = future.MakeAssetTransferTxn(poolAddress.String(), senderAdd, assetOutAmt, nil, sp, "", assetOutID)
	} else {
		// (b) Pay - Transfer of buy asset from Pool to Swapper
		// if buy asset is an ALGO
		tx4, err = future.MakePaymentTxn(poolAddress.String(), senderAdd, assetOutAmt, nil, "", sp)
	}
	if err != nil {
		return nil, err
	}

	txnGroup, err := NewTransactionGroup([]types.Transaction{tx1, tx2, tx3, tx4})
	if err != nil {
		return nil, err
	}
	if err := txnGroup.SignWithLogicSig(poolLogicSigAccount); err != nil {
		return nil, err
	}

	return txnGroup, nil
}
