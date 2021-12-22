package tinyman

import (
	"encoding/json"
	"io/ioutil"
)

// TinymanContracts ...
var TinymanContracts ASC

func init() {
	b, err := ioutil.ReadFile("asc.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &TinymanContracts); err != nil {
		panic(err)
	}
}

// ASC ...
type ASC struct {
	Contracts Contracts `json:"contracts"`
}

// Contracts ...
type Contracts struct {
	PoolLogicSig Contract `json:"pool_logicsig"`
	// validatorApp ValidatorContract
}

// Contract ...
type Contract struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Logic Logic  `json:"logic"`
}

// Logic ...
type Logic struct {
	Bytecode  string     `json:"bytecode"`
	Address   string     `json:"address"`
	Size      int        `json:"size"`
	Variables []Variable `json:"variables"`
}

// Variable ...
type Variable struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Index  int    `json:"index"`
	Length int    `json:"length"`
}
