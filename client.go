package tinyman

import (
	"context"

	"github.com/algorand/go-algorand-sdk/client/algod/models"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/types"
)

const (
	AlgoExplorerTestnet = "https://api.testnet.algoexplorer.io"
	AlgoExplorerMainnet = "https://api.algoexplorer.io"

	TestnetValidatorAppID = 21580889
	MainnetValidatorAppID = 350338509
)

// Client ...
type Client struct {
	ac             *algod.Client
	validatorAppID uint64
	address        string
	assetsCache    map[uint64]models.Asset
}

// NewClient ...
func NewClient(ac *algod.Client, validatorAppID uint64, userAddress string) *Client {
	return &Client{
		ac:             ac,
		validatorAppID: validatorAppID,
		address:        userAddress,
		assetsCache:    map[uint64]models.Asset{},
	}
}

// Pool ...
func (c *Client) Pool(ctx context.Context, asset1ID, asset2ID uint64) (*Pool, error) {
	return FetchPool(ctx, c, asset1ID, asset2ID)
}

// PrepareAppOptInTxns ...
func (c *Client) PrepareAppOptInTxns(ctx context.Context, address string) ([]types.Transaction, error) {
	if address == "" {
		address = c.address
	}

	params, err := c.ac.SuggestedParams().Do(ctx)
	if err != nil {
		return nil, err
	}
	txns, err := PrepareAppOptInTxns(c.validatorAppID, address, params)
	if err != nil {
		return nil, err
	}

	return txns, nil
}

// IsOptedIn ...
func (c *Client) IsOptedIn(ctx context.Context, address string) (bool, error) {
	if address == "" {
		address = c.address
	}

	acc, err := c.ac.AccountInformation(address).Do(ctx)
	if err != nil {
		return false, err
	}

	for _, a := range acc.AppsLocalState {
		if a.Id == c.validatorAppID {
			return true, nil
		}
	}

	return false, nil
}

// Asset ...
// func (c *Client) Asset(ctx context.Context, assetID uint64) (models.Asset, error) {
// 	asset, ok := c.assetsCache[assetID]
// 	if !ok {
// 		asset, err := c.ac.GetAssetByID(a.ID).Do(ctx)
// 		if err != nil {
// 			return nil, err
// 		}
// 		c.assetsCache[assetID] = asset
// 	}
//
// 	return asset, nil
// }

// Submit ...
// func (c *Client) Submit(transactionGroup) {
// }

// ExcessAmounts ...
// func (c *Client) ExcessAmounts(ctx context.Context, address string) (map[uint64]uint64, error) {
// 	if address == "" {
// 		address = p.Client.Address
// 	}
// 	acc, err := c.ac.AccountInformation(address).Do(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var validatorApp models.ApplicationLocalState
// 	for _, a := range acc.AppsLocalState {
// 		if a.Id == c.validatorAppID {
// 			validatorApp = a
// 		}
// 	}
//
// 	// validatorAppState
// 	excess := map[uint64]uint64{}
// 	return excess, nil
// }
//
