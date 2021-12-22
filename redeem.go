package tinyman

const (
	redeemFee = 2000
)

// Redeem claims back 'change' due to slippage in Mint/Burn/Swap process.
// func Redeem(
// 	validatorAppID uint64,
// 	asset1ID, asset2ID, liquidityAssetID uint64,
// 	assetID uint64,
// 	assetAmt uint64,
// 	senderAdd string,
// 	sp types.SuggestedParams) ([]byte, error) {
// fmt.Println("Redeeming")
//
// poolLogicSigAccount := getLogicSigAccount(validatorAppID, asset1ID, asset2ID)
// poolAddress, err := poolLogicSigAccount.Address()
// if err != nil {
// 	return nil, err
// }
//
// // PaymentTxn
// // Pay - pay fees in Algo from PoolerSwapper to Pool
// tx1, err := future.MakePaymentTxn(senderAdd, poolAddress.String(), redeemFee, []byte("fee"), false, sp)
// if err != nil {
// 	return nil, err
// }
//
// // ApplicationNoOpTxn
// // App Call - NoOp call to Validator App with args ['redeem'], with
// // Pooler/Swapper account
// foreignAssets := []uint64{asset1ID}
// if asset2ID != 0 {
// 	foreignAssets = append(foreignAssets, asset2ID)
// }
// foreignAssets = append(foreignAssets, liquidityAssetID)
// args := [][]byte{
// 	[]byte("redeem"),
// }
// accounts := []string{senderAdd}
//
// tx2, err := future.MakeApplicationNoOpTx(validatorAppID, args, accounts, nil, foreignAssets, sp, poolAddress.String(), nil, nil, nil, nil)
// if err != nil {
// 	return nil, err
// }
//
// // AssetTransferTxn
// var tx3 types.Transaction
// if assetID != 0 {
// 	// (a) AssetTransfer - Transfer of asset from Pool to Pooler/Swapper
// 	// if asset is an ASA
// 	tx3, err = future.MakeAssetTransferTxn(poolAddress.String(), senderAdd, assetAmt, nil, sp, false, assetID)
// } else {
// 	// (b) Pay - Transfer of Algo from Pool to Pooler/Swapper
// 	// if asset is ALGO
// 	tx3, err = future.MakePaymentTxn(poolAddress.String(), senderAdd, assetAmt, nil, false, sp)
// }
// if err != nil {
// 	return nil, err
// }
//
// txns := []types.Transaction{tx1, tx2, tx3}
// gid, err := crypto.ComputeGroupID(txns)
// if err != nil {
// 	return nil, err
// }
// for i, tx := range txns {
// 	txns[i].Group = gid
// }
//
// sg := []byte{}
// if err := poolLogicSigAccount.signTxns(sg, txns); err != nil {
// 	return nil, err
// }
//
// return sg, nil
// }
