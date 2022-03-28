package tinyman

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/types"
)

func getStateInt(appState map[string]models.TealValue, key []byte) uint64 {
	str := base64.StdEncoding.EncodeToString(key)
	v, ok := appState[str]
	if !ok {
		return 0
	}

	return v.Uint
}

func readProgram(logic Logic, params map[string]uint64) ([]byte, error) {
	encodedTemplate := logic.Bytecode
	t, err := base64.StdEncoding.DecodeString(encodedTemplate)
	if err != nil {
		return nil, err
	}

	offset := 0
	result := make([]byte, len(t))
	copy(result, t)

	sort.Slice(logic.Variables, func(i, j int) bool {
		return logic.Variables[i].Index < logic.Variables[j].Index
	})

	for _, v := range logic.Variables {
		val := params[v.Name]
		start := v.Index - offset
		end := start + v.Length
		encodedVal, err := encodeValue(val, v.Type)
		if err != nil {
			return nil, err
		}
		encodedValLen := len(encodedVal)
		diff := v.Length - encodedValLen
		offset += diff
		temp := result[end:]
		result = append(result[:start], encodedVal...)
		result = append(result, temp...)
	}

	return result, nil
}

func encodeValue(val uint64, t string) ([]byte, error) {
	if t != "int" {
		return nil, fmt.Errorf("Unsupported value type %s!", t)
	}

	var buf []byte
	b := bytes.NewBuffer(buf)
	for {
		towrite := val & uint64(127)
		val >>= 7
		c := make([]byte, 8)
		if val != 0 {
			binary.LittleEndian.PutUint64(c, (towrite | uint64(128)))
			b.Write([]byte{c[0]})
		} else {
			binary.LittleEndian.PutUint64(c, towrite)
			b.Write([]byte{c[0]})
			break
		}
	}

	return b.Bytes(), nil
}

func getPoolLogicSigAccount(validatorAppID, asset1ID, asset2ID uint64) (crypto.LogicSigAccount, error) {
	params := map[string]uint64{
		"TMPL_ASSET_ID_1":       asset1ID,
		"TMPL_ASSET_ID_2":       asset2ID,
		"TMPL_VALIDATOR_APP_ID": validatorAppID,
	}

	poolLogicSigDef := TinymanContracts.Contracts.PoolLogicSig.Logic
	program, err := readProgram(poolLogicSigDef, params)
	if err != nil {
		return crypto.LogicSigAccount{}, err
	}

	return crypto.MakeLogicSigAccountEscrow(program, nil), nil
}

// TransactionGroup ...
type TransactionGroup struct {
	txns       []types.Transaction
	signedTxns [][]byte
}

// NewTransactionGroup ...
func NewTransactionGroup(txns []types.Transaction) (*TransactionGroup, error) {
	gid, err := crypto.ComputeGroupID(txns)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(txns); i++ {
		txns[i].Group = gid
	}

	return &TransactionGroup{
		txns:       txns,
		signedTxns: make([][]byte, len(txns)),
	}, nil
}

// SignWithLogicSig ...
func (tg *TransactionGroup) SignWithLogicSig(lsa crypto.LogicSigAccount) error {
	address, err := lsa.Address()
	if err != nil {
		return err
	}
	for i := 0; i < len(tg.txns); i++ {
		if address != tg.txns[i].Sender {
			continue
		}

		_, stx, err := crypto.SignLogicSigAccountTransaction(lsa, tg.txns[i])
		if err != nil {
			return err
		}
		tg.signedTxns[i] = stx
	}

	return nil
}

// SignWithPrivateKey ...
func (tg *TransactionGroup) SignWithPrivateKey(pk ed25519.PrivateKey) error {
	acc, err := crypto.AccountFromPrivateKey(pk)
	if err != nil {
		return err
	}

	address := acc.Address
	for i := 0; i < len(tg.txns); i++ {
		if address != tg.txns[i].Sender {
			continue
		}

		_, stx, err := crypto.SignTransaction(pk, tg.txns[i])
		if err != nil {
			return err
		}
		tg.signedTxns[i] = stx
	}

	return nil
}

// RawSignedTransactions ...
func (tg *TransactionGroup) RawSignedTransactions() []byte {
	var sg []byte
	for i := 0; i < len(tg.txns); i++ {
		sg = append(sg, tg.signedTxns[i]...)
	}

	return sg
}

// Transactions ...
func (tg *TransactionGroup) Transactions() []types.Transaction {
	return tg.txns
}
