package tinyman

const (
	burnFee = 4000
)

// Burn Pool liquidity assets in exchange for removing assets from the Pool.
// func Burn(
// 	validatorAppID uint64,
// 	asset1ID, asset2ID, liquidityAssetID uint64,
// 	asset1Amt, asset2Amt, liquidityAssetAmt uint64,
// 	senderAdd string,
// 	sp types.SuggestedParams) ([]byte, error) {
// fmt.Println("Burning")
//
// poolLogicSigAccount := getLogicSigAccount(validatorAppID, asset1ID, asset2ID)
// poolAddress, err := poolLogicSigAccount.Address()
// if err != nil {
// 	return nil, err
// }
//
// // PaymentTxn
// // Pay - pay fees in Algo from Pooler to Pool
// tx1, err := future.MakePaymentTxn(senderAdd, poolAddress.String(), burnFee, []byte("fee"), false, sp)
// if err != nil {
// 	return nil, err
// }
//
// // ApplicationNoOpTxn
// // App Call - NoOp call to Validator App with args ['burn'], with Pooler
// // account
// foreignAssets := []uint64{asset1ID}
// if asset2ID != 0 {
// 	foreignAssets = append(foreignAssets, asset2ID)
// }
// foreignAssets = append(foreignAssets, liquidityAssetID)
// args := [][]byte{
// 	[]byte("burn"),
// }
// accounts := []string{senderAdd}
//
// tx2, err := future.MakeApplicationNoOpTx(validatorAppID, args, accounts, nil, foreignAssets, sp, poolAddress.String(), nil, nil, nil, nil)
// if err != nil {
// 	return nil, err
// }
//
// // AssetTransferTxn
// // AssetTransfer - Transfer of asset 1 from Pool to Pooler
// tx3, err := future.MakeAssetTransferTxn(poolAddress.String(), senderAddr, asset1Amt, nil, sp, false, asset1ID)
// if err != nil {
// 	return nil, err
// }
//
// // AssetTransferTxn
// var tx4 types.Transaction
// if asset2ID != 0 {
// 	// (a) AssetTransfer - Transfer of asset 2 from Pool to Pooler
// 	// - If asset 2 is an ASA
// 	tx4, err = future.MakeAssetTransferTxn(poolAddress.String(), senderAddr, asset2Amt, nil, sp, false, asset2ID)
// } else {
// 	// (b) Pay - Transfer of Algo from Pool to Pooler
// 	// - If asset 2 is Algo
// 	tx4, err = future.MakePaymentTxn(poolAddress.String(), senderAddr, asset2Amt, nil, false, sp)
// }
// if err != nil {
// 	return nil, err
// }
//
// // AssetTransferTxn
// // AssetTransfer - Transfer of liquidity token asset from Pooler to Pool
// tx5, err := future.MakeAssetTransferTxn(senderAdd, poolAddress.String(), liquidityAssetAmt, nil, sp, nil, liquidityAssetID)
// if err != nil {
// 	return nil, err
// }
//
// txns := []types.Transaction{tx1, tx2, tx3, tx4, tx5}
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
