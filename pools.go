package tinyman

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
)

const (
	FixedOutput SwapType = "fixed-output"
	FixedInput  SwapType = "fixed-input"
)

// SwapType ...
type SwapType string

// SwapQuote ...
type SwapQuote struct {
	SwapType       SwapType
	AssetAmountIn  AssetAmount
	AssetAmountOut AssetAmount
	SwapFees       float64
	Slippage       float64
}

// AmountOutWithSlippage ...
func (sq *SwapQuote) AmountOutWithSlippage() float64 {
	if sq.SwapType == FixedOutput {
		return float64(sq.AssetAmountOut.Amount)
	}

	return float64(sq.AssetAmountOut.Amount) - (float64(sq.AssetAmountOut.Amount) * sq.Slippage)
}

// AmountInWithSlippage ...
func (sq *SwapQuote) AmountInWithSlippage() float64 {
	if sq.SwapType == FixedInput {
		return float64(sq.AssetAmountIn.Amount)
	}

	return float64(sq.AssetAmountIn.Amount) - (float64(sq.AssetAmountIn.Amount) * sq.Slippage)
}

// Price ...
func (sq *SwapQuote) Price() float64 {
	return float64(sq.AssetAmountOut.Amount) / float64(sq.AssetAmountIn.Amount)
}

// PriceWithSlippage ...
func (sq *SwapQuote) PriceWithSlippage() float64 {
	return float64(sq.AmountOutWithSlippage()) / float64(sq.AmountInWithSlippage())
}

// BurnQuote ...
type BurnQuote struct {
	AmountsOut           map[uint64]AssetAmount
	LiquidityAssetAmount AssetAmount
	Slippage             float64
}

// PoolPosition ...
type PoolPosition struct {
	Asset1         AssetAmount
	Asset2         AssetAmount
	LiquidityAsset AssetAmount
	Share          float64
}

// GetPoolInfo ...
func GetPoolInfo(ctx context.Context, node *Node, validatorAppID, asset1ID, asset2ID uint64) (Pool, error) {
	poolLogicSigAccount, err := getPoolLogicSigAccount(validatorAppID, asset1ID, asset2ID)
	if err != nil {
		return Pool{}, err
	}

	poolAddress, err := poolLogicSigAccount.Address()
	if err != nil {
		return Pool{}, err
	}

	node.Take()
	accInfo, err := node.ac.AccountInformation(poolAddress.String()).Do(ctx)
	if err != nil {
		return Pool{}, err
	}
	p, err := GetPoolInfoFromAccountInfo(ctx, accInfo)
	if err != nil {
		return Pool{}, err
	}

	p.Client = NewClient(node.ac, validatorAppID, "")

	return p, nil
}

// GetPoolInfoFromAccountInfo ...
func GetPoolInfoFromAccountInfo(ctx context.Context, accInfo models.Account) (Pool, error) {
	var p Pool
	err := p.UpdateFromAccountInfo(ctx, accInfo)

	return p, err
}

// Pool ...
type Pool struct {
	Client  *Client
	Address string

	Asset1ID           uint64
	Asset2ID           uint64
	Asset1Decimals     uint64
	Asset2Decimals     uint64
	Asset1UnitName     string
	Asset2UnitName     string
	LiquidityAssetID   uint64
	LiquidityAssetName string
	Asset1Reserves     uint64
	Asset2Reserves     uint64

	IssuedLiquidity                 uint64
	UnclaimedProtocolFees           uint64
	OutstandingAsset1Amount         uint64
	OutstandingAsset2Amount         uint64
	OutstandingLiquidityAssetAmount uint64
	ValidatorAppID                  uint64
	AlgoBalance                     uint64
	Round                           uint64
}

// FetchPool ...
func FetchPool(ctx context.Context, c *Client, asset1ID, asset2ID uint64) (*Pool, error) {
	p := &Pool{
		Client: c,
	}

	p.Asset1ID = asset2ID
	p.Asset2ID = asset1ID
	if asset1ID > asset2ID {
		p.Asset1ID = asset1ID
		p.Asset2ID = asset2ID
	}

	p.Client.node.Take()
	asset1Info, err := p.Client.node.ac.GetAssetByID(p.Asset1ID).Do(ctx)
	if err != nil {
		return nil, err
	}
	p.Asset1Decimals = asset1Info.Params.Decimals
	p.Asset1UnitName = asset1Info.Params.UnitName

	p.Client.node.Take()
	asset2Info, err := p.Client.node.ac.GetAssetByID(p.Asset2ID).Do(ctx)
	if err != nil {
		return nil, err
	}
	p.Asset2Decimals = asset2Info.Params.Decimals
	p.Asset2UnitName = asset2Info.Params.UnitName

	if err := p.Refresh(ctx); err != nil {
		return nil, err
	}

	return p, nil
}

// Refresh ...
func (p *Pool) Refresh(ctx context.Context) error {
	poolLogicSigAccount, err := getPoolLogicSigAccount(p.Client.validatorAppID, p.Asset1ID, p.Asset2ID)
	if err != nil {
		return err
	}

	poolAddress, err := poolLogicSigAccount.Address()
	if err != nil {
		return err
	}

	p.Client.node.Take()
	accInfo, err := p.Client.node.ac.AccountInformation(poolAddress.String()).Do(ctx)
	if err != nil {
		return err
	}
	if err := p.UpdateFromAccountInfo(ctx, accInfo); err != nil {
		return err
	}

	return nil
}

// UpdateFromAccountInfo ...
func (p *Pool) UpdateFromAccountInfo(ctx context.Context, accInfo models.Account) error {
	if accInfo.AppsLocalState == nil || len(accInfo.AppsLocalState) == 0 {
		return fmt.Errorf("pools: no local application state")
	}
	validatorAppID := accInfo.AppsLocalState[0].Id

	idx := 0
	validatorAppState := map[string]models.TealValue{}
	for _, kv := range accInfo.AppsLocalState[idx].KeyValue {
		validatorAppState[kv.Key] = kv.Value
	}

	asset1ID := getStateInt(validatorAppState, []byte("a1"))
	asset2ID := getStateInt(validatorAppState, []byte("a2"))

	poolLogicSig, err := getPoolLogicSigAccount(validatorAppID, asset1ID, asset2ID)
	if err != nil {
		return err
	}
	poolAddress, err := poolLogicSig.Address()
	if err != nil {
		return err
	}

	if accInfo.Address != poolAddress.String() {
		return fmt.Errorf("get_pool_info: account address does not match pool address")
	}

	asset1Reserves := getStateInt(validatorAppState, []byte("s1"))
	asset2Reserves := getStateInt(validatorAppState, []byte("s2"))
	issuedLiquidity := getStateInt(validatorAppState, []byte("ilt"))
	unclaimedProtocolFees := getStateInt(validatorAppState, []byte("p"))

	liquidityAsset := accInfo.CreatedAssets[0]
	liquidityAssetID := liquidityAsset.Index

	c := make([]byte, 8)
	binary.BigEndian.PutUint64(c, asset1ID)
	temp := append([]byte("o"), c...)
	outstandingAsset1Amount := getStateInt(validatorAppState, temp)

	binary.BigEndian.PutUint64(c, asset2ID)
	temp = append([]byte("o"), c...)
	outstandingAsset2Amount := getStateInt(validatorAppState, temp)

	binary.BigEndian.PutUint64(c, liquidityAssetID)
	temp = append([]byte("o"), c...)
	outstandingLiquidityAssetAmount := getStateInt(validatorAppState, temp)

	p.Address = poolAddress.String()
	p.Asset1ID = asset1ID
	p.Asset2ID = asset2ID
	p.LiquidityAssetID = liquidityAssetID
	p.LiquidityAssetName = liquidityAsset.Params.Name
	p.Asset1Reserves = asset1Reserves
	p.Asset2Reserves = asset2Reserves
	p.IssuedLiquidity = issuedLiquidity
	p.UnclaimedProtocolFees = unclaimedProtocolFees
	p.OutstandingAsset1Amount = outstandingAsset1Amount
	p.OutstandingAsset2Amount = outstandingAsset2Amount
	p.OutstandingLiquidityAssetAmount = outstandingLiquidityAssetAmount
	p.ValidatorAppID = validatorAppID
	p.AlgoBalance = accInfo.Amount
	p.Round = accInfo.Round

	return nil
}

// Prices ...
func (p *Pool) Prices() (asset1Price, asset2Price float64) {
	asset1Price = float64(p.Asset2Reserves) / float64(p.Asset1Reserves)
	asset2Price = float64(p.Asset1Reserves) / float64(p.Asset2Reserves)

	return asset1Price, asset2Price
}

// MaximumAmount ...
func (p *Pool) MaximumAmount(assetID uint64, targetPrice float64) uint64 {
	amt := targetPrice * float64(p.Asset2Reserves)
	if assetID == p.Asset2ID {
		amt = targetPrice * float64(p.Asset1Reserves)
	}

	return uint64(amt)
}

// FixedInputSwapQuote ...
func (p *Pool) FixedInputSwapQuote(assetID, amount uint64, slippage float64) (SwapQuote, error) {
	assetIn, assetInAmount := assetID, amount
	assetOut := p.Asset1ID
	inputSupply := p.Asset2Reserves
	outputSupply := p.Asset1Reserves

	if assetIn == p.Asset1ID {
		assetOut = p.Asset2ID
		inputSupply = p.Asset1Reserves
		outputSupply = p.Asset2Reserves
	}

	if inputSupply <= 0 || outputSupply <= 0 {
		return SwapQuote{}, fmt.Errorf("pool: no liquidity")
	}

	k := float64(inputSupply) * float64(outputSupply)
	assetInAmountMinusFee := (float64(assetInAmount) * float64(997)) / float64(1000)
	swapFees := float64(assetInAmount) - assetInAmountMinusFee
	assetOutAmount := float64(outputSupply) - (float64(k) / (float64(inputSupply) + float64(assetInAmountMinusFee)))

	amountIn := AssetAmount{
		AssetID: assetIn,
		Amount:  uint64(assetInAmount),
	}
	amountOut := AssetAmount{
		AssetID: assetOut,
		Amount:  uint64(assetOutAmount),
	}

	quote := SwapQuote{
		SwapType:       FixedInput,
		AssetAmountIn:  amountIn,
		AssetAmountOut: amountOut,
		SwapFees:       swapFees,
		Slippage:       slippage,
	}

	return quote, nil
}

// ExcessAmounts ...
// func (p *Pool) ExcessAmounts(address string) map[uint64]uint64 {
// 	if address == "" {
// 		address = p.Client.Address
// 	}
// 	return p.Client.ExcessAmounts(address)
// }

// RedeemTransactions ...
// func (p *Pool) RedeemTransactions(ctx context.Context, amountOut AssetAmount, address string) []types.Transaction {
// 	if address == "" {
// 		address = p.Client.Address
// 	}
//
// 	params, err := c.ac.SuggestedParams().Do(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return Redeem(val)
// }

// BurnQuote ...
func (p *Pool) BurnQuote(liquidityAssetAmt uint64, slippage float64) BurnQuote {
	liquidityAssetAmount := AssetAmount{
		AssetID: p.LiquidityAssetID,
		Amount:  liquidityAssetAmt,
	}

	asset1Amt := (liquidityAssetAmount.Amount * p.Asset1Reserves) / p.IssuedLiquidity
	asset2Amt := (liquidityAssetAmount.Amount * p.Asset2Reserves) / p.IssuedLiquidity

	amountsOut := map[uint64]AssetAmount{
		p.Asset1ID: AssetAmount{AssetID: p.Asset1ID, Amount: asset1Amt},
		p.Asset2ID: AssetAmount{AssetID: p.Asset2ID, Amount: asset2Amt},
	}

	return BurnQuote{
		AmountsOut:           amountsOut,
		LiquidityAssetAmount: liquidityAssetAmount,
		Slippage:             slippage,
	}
}

// PoolPosition ...
func (p *Pool) PoolPosition(ctx context.Context, address string) (PoolPosition, error) {
	if address == "" {
		address = p.Client.address
	}

	p.Client.node.Take()
	acc, err := p.Client.node.ac.AccountInformation(address).Do(ctx)
	if err != nil {
		return PoolPosition{}, err
	}

	var liquidityAmt uint64
	for _, asset := range acc.Assets {
		if asset.AssetId == p.LiquidityAssetID {
			liquidityAmt = asset.Amount
			break
		}
	}

	quote := p.BurnQuote(liquidityAmt, 0.01)
	poolPosition := PoolPosition{
		Asset1:         quote.AmountsOut[p.Asset1ID],
		Asset2:         quote.AmountsOut[p.Asset2ID],
		LiquidityAsset: quote.LiquidityAssetAmount,
		Share:          (float64(liquidityAmt) / float64(p.IssuedLiquidity)),
	}

	return poolPosition, nil
}

// PrepareSwapTransactions ...
func (p *Pool) PrepareSwapTransactions(ctx context.Context, assetInID, amountIn, amountOut uint64, swapType SwapType, address string) (*TransactionGroup, error) {
	if address == "" {
		address = p.Client.address
	}

	p.Client.node.Take()
	txParams, err := p.Client.node.ac.SuggestedParams().Do(ctx)
	if err != nil {
		return nil, err
	}

	return PrepareSwapTransactions(
		txParams,
		p.ValidatorAppID,
		p.Asset1ID, p.Asset2ID, p.LiquidityAssetID,
		assetInID,
		amountIn, amountOut,
		swapType, address,
	)
}

// PrepareSwapTransactionsFromQuote ...
func (p *Pool) PrepareSwapTransactionsFromQuote(ctx context.Context, quote *SwapQuote, address string) (*TransactionGroup, error) {
	return p.PrepareSwapTransactions(
		ctx,
		quote.AssetAmountIn.AssetID,
		uint64(quote.AmountInWithSlippage()), uint64(quote.AmountOutWithSlippage()),
		quote.SwapType,
		address,
	)
}
