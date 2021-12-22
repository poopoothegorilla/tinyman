package tinyman

const (
	tinymanPoolUnitName = "TM1POOL"
	tinymanURL          = "https://tinyman.org"
)

// Bootstrap setups a Pool for a pair of assets. The Pool account should be a
// LogicSig contract account.
// func Bootstrap(
// 	validatorAppID uint64,
// 	asset1ID, asset2ID uint64,
// 	asset1UnitName, asset2UnitName string,
// 	senderAdd string,
// 	sp types.SuggestedParams) ([]byte, error) {
// fmt.Println("Bootstrap")
//
// poolLogicSigAccount := getLogicSigAccount(validatorAppID, asset1ID, asset2ID)
// poolAddress, err := poolLogicSigAccount.Address()
// if err != nil {
// 	return nil, err
// }
//
// if asset1ID < asset2ID {
// 	return fmt.Errorf("bootstrap: asset_1_id(%v) < asset_2_id(%v)", asset1ID, asset2ID)
// }
//
// if asset2ID == 0 {
// 	asset2UnitName = "ALGO"
// }
//
// // PaymentTxn
// // Pay - pay fees from Pooler to Pool
// bootstrapFee := 961000
// if asset2ID > 0 {
// 	bootstrapFee = 860000
// }
// tx1, err := future.MakePaymentTxn(senderAdd, poolAddress.String(), bootstrapFee, []byte("fee"), false, sp)
// if err != nil {
// 	return nil, err
// }
//
// // ApplicationOptInTxn
// // App Call - OptIn call to Validator App with args ['bootstrap', asset1ID,
// // asset2ID]
// args := [][]byte{
// 	[]byte("bootstrap"),
// 	// asset1ID
// 	// asset2ID
// }
// foreignAssets := []uint64{asset1ID}
// if asset2ID != 0 {
// 	foreignAssets = append(foreignAssets, asset2ID)
// }
// tx2, err := future.MakeApplicationOptInTx(validatorAppID, args, nil, nil, foreignAssets, sp, poolAddress.String(), nil, nil, nil, nil)
// if err != nil {
// 	return nil, err
// }
//
// // AssetCreateTxn
// // AssetConfig - create asset for liquidity token
// tx3, err := future.MakeAssetCreateTxn(
// 	poolAddress.String(),
// 	nil,
// 	sp,
// 	total, 6, false, manager, reserve, freeze, clawback, tinymanPoolUnitName,
// 	fmt.Sprintf("Tinyman Pool %s-%s", asset1UnitName, asset2UnitName),
// 	tinymanURL, "")
//
// // AssetOptInTxn
// // Asset OptIn - Pool opt in to Asset 1
// tx4, err := future.MakeAssetAcceptanceTxn(poolAddress.String(), nil, sp, asset1ID)
// if err != nil {
// 	return nil, err
// }
//
// txns := []types.Transaction{tx1, tx2, tx3, tx4}
//
// // AssetOptInTxn
// // (Optional) Asset OptIn - Pool opt in to Asset 2
// // Only if Asset 2 is not Algo
// if asset2ID > 0 {
// 	tx, err := future.MakeAssetAcceptanceTxn(poolAddress.String(), nil, sp, asset2ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	txns = append(txns, txn)
// }
//
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
