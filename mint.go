package tinyman

// "github.com/algorand/go-algorand-sdk/transaction"
// "github.com/algorand/go-algorand-sdk/client/v2/algod"
// 		"github.com/algorand/go-algorand-sdk/mnemonic"

const (
	mintFee = 2000
)

// Mint Pool assets in exchange for transferring assets to the Pool account.
// func Mint(
// 	validatorAppID uint64,
// 	asset1ID, asset2ID, liquidityAssetID uint64,
// 	asset1Amt, asset2Amt, liquidityAssetAmt uint64,
// 	senderAdd string,
// 	sp types.SuggestedParams) ([]byte, error) {
// fmt.Println("Minting")
//
// poolLogicSigAccount := getLogicSigAccount(validatorAppID, asset1ID, asset2ID)
// poolAddress, err := poolLogicSigAccount.Address()
// if err != nil {
// 	return nil, err
// }
//
// // PaymentTxn
// // Pay - pay fees in Algo from Pooler to Pool
// tx1, err := future.MakePaymentTxn(senderAdd, poolAddress.String(), mintFee, []byte("fee"), false, sp)
// if err != nil {
// 	return nil, err
// }
//
// // ApplicationNoOpTxn
// // App Call - NoOp call to Validator App with args ['mint'], with Pooler
// // account
// foreignAssets := []uint64{asset1ID}
// if asset2ID != 0 {
// 	foreignAssets = append(foreignAssets, asset2ID)
// }
// foreignAssets = append(foreignAssets, liquidityAssetID)
// args := [][]byte{
// 	[]byte("mint"),
// }
// accounts := []string{senderAdd}
//
// tx2, err := future.MakeApplicationNoOpTx(validatorAppID, args, accounts, nil, foreignAssets, sp, poolAddress.String(), nil, nil, nil, nil)
// if err != nil {
// 	return nil, err
// }
//
// // AssetTransferTxn
// // AssetTransfer - Transfer of asset 1 from Pooler to Pool
// tx3, err := future.MakeAssetTransferTxn(senderAdd, poolAddress.String(), asset1Amt, nil, sp, false, asset1ID)
// if err != nil {
// 	return nil, err
// }
//
// // AssetTransferTxn
// var tx4 types.Transaction
// if asset2ID != 0 {
// 	// (a) AssetTransfer - Transfer of asset 2 from Pooler to Pool
// 	// - If asset 2 is an ASA
// 	tx4, err = future.MakeAssetTransferTxn(senderAdd, poolAddress.String(), asset2Amt, nil, sp, false, asset2ID)
// } else {
// 	// (b) Pay - Transfer of Algo from Pooler to Pool
// 	// - If asset 2 is Algo
// 	tx4, err = future.MakePaymentTxn(senderAdd, poolAddress.String(), asset2Amt, nil, false, sp)
// }
// if err != nil {
// 	return nil, err
// }
//
// // AssetTransferTxn
// // AssetTransfer - Transfer of liquidity token asset from Pool to Pooler
// // - Amount is minimum expected amount of liquidity token allowing for slippage
// tx5, err := future.MakeAssetTransferTxn(poolAddress.String(), senderAdd, liquidityAssetAmt, nil, sp, nil, liquidityAssetID)
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
