package tinyman

import (
	"encoding/json"
)

// TinymanContracts ...
var TinymanContracts ASC

func init() {
	if err := json.Unmarshal(ascBlob, &TinymanContracts); err != nil {
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

var ascBlob = []byte(`
{
  "repo": "https://github.com/tinymanorg/tinyman-contracts-v1",
  "ref": "6e150df8e3e74458f947f3399c6f37e7d4289d21",
  "contracts": {
    "pool_logicsig": {
      "type": "logicsig",
      "logic": {
        "bytecode": "BCAIAQCBgICAgICAgPABAwSAgICAgICAgPABBQYhBSQNRDEJMgMSRDEVMgMSRDEgMgMSRDIEIg1EMwEAMQASRDMBECEHEkQzARiBgoCAgICAgIDwARJEMwEZIhIzARslEhA3ARoAgAlib290c3RyYXASEEAAXDMBGSMSRDMBG4ECEjcBGgCABHN3YXASEEACEzMBGyISRDcBGgCABG1pbnQSQAE5NwEaAIAEYnVybhJAAYM3ARoAgAZyZWRlZW0SQAIzNwEaAIAEZmVlcxJAAlAAIQYhBCQjEk0yBBJENwEaARchBRI3ARoCFyQSEEQzAgAxABJEMwIQJRJEMwIhIxJEMwIiIxwSRDMCIyEHEkQzAiQjEkQzAiWAB1RNMVBPT0wSRDMCJlEADYANVGlueW1hbiBQb29sIBJEMwIngBNodHRwczovL3RpbnltYW4ub3JnEkQzAikyAxJEMwIqMgMSRDMCKzIDEkQzAiwyAxJEMwMAMQASRDMDECEEEkQzAxEhBRJEMwMUMQASRDMDEiMSRCQjE0AAEDMBATMCAQgzAwEINQFCAYkzBAAxABJEMwQQIQQSRDMEESQSRDMEFDEAEkQzBBIjEkQzAQEzAgEIMwMBCDMEAQg1AUIBVDIEIQYSRDcBHAExABNENwEcATMEFBJEMwIAMQATRDMCFDEAEkQzAwAzAgASRDMDFDMDBzMDECISTTEAEkQzBAAxABJEMwQUMwIAEkQzAQEzBAEINQFCAPwyBCEGEkQ3ARwBMQATRDcBHAEzAhQSRDMDFDMDBzMDECISTTcBHAESRDMCADEAEkQzAhQzBAASRDMDADEAEkQzAxQzAwczAxAiEk0zBAASRDMEADEAE0QzBBQxABJEMwEBMwIBCDMDAQg1AUIAjjIEIQQSRDcBHAExABNEMwIANwEcARJEMwIAMQATRDMDADEAEkQzAhQzAgczAhAiEk0xABJEMwMUMwMHMwMQIhJNMwIAEkQzAQEzAwEINQFCADwyBCUSRDcBHAExABNEMwIUMwIHMwIQIhJNNwEcARJEMwEBMwIBCDUBQgARMgQlEkQzAQEzAgEINQFCAAAzAAAxABNEMwAHMQASRDMACDQBD0M=",
        "address": "5MKWI634X65LPTLRYB6PP4IVMTV75UKTYID2BQ5ATCVKXUW5XYGTMU7BSI",
        "size": 839,
        "variables": [
          {
            "name": "TMPL_ASSET_ID_1",
            "type": "int",
            "index": 17,
            "length": 10
          },
          {
            "name": "TMPL_ASSET_ID_2",
            "type": "int",
            "index": 5,
            "length": 10
          },
          {
            "name": "TMPL_VALIDATOR_APP_ID",
            "type": "int",
            "index": 75,
            "length": 10
          }
        ],
        "source": "https://github.com/tinymanorg/tinyman-contracts-v1/tree/6e150df8e3e74458f947f3399c6f37e7d4289d21/contracts/pool_logicsig.teal.tmpl"
      },
      "name": "pool_logicsig"
    },
    "validator_app": {
      "type": "app",
      "approval_program": {
        "bytecode": "BCAHAAHoB+UHBf///////////wHAhD0mCwFvAWUBcAJhMQJhMgJsdARzd2FwBG1pbnQBdAJwMQJwMjEZgQQSMRkhBBIRMRmBAhIRQATAMRkjEjEbIhIQQASyNhoAgAZjcmVhdGUSQASjMRkjEjYaAIAJYm9vdHN0cmFwEhBAA8IzAhIzAggINTQiK2I1ZSI0ZXAARDUBIicEYjVmNGZAABEiYCJ4CTEBCDMACAk1AkIACCI0ZnAARDUCIicFYjVnKDRlFlA1byI0b2I1PSg0ZhZQNXAiNHBiNT4oNGcWUDVxIjRxYjU/IipiNUA0ATQ9CTVHNAI0Pgk1SDEAKVA0ZRZQNXkxAClQNGYWUDV6MQApUDRnFlA1ezYaAIAGcmVkZWVtEkAAWjYaAIAEZmVlcxJAABw2GgAnBhI2GgAnBxIRNhoAgARidXJuEhFAAFwANGdJRDMCERJEMwISRDMCFDIJEkQ0PzMCEgk1PzRAMwISCTVAIio0QGYiNHE0P2YjQzQ0RCIoMwIRFlBKYjQ0CWYjMQApUDMCERZQSmI0NAlJQQADZiNDSGgjQzIHIicIYglJNfpBAD4igAJjMUpiIicJYjT6Cx4hBSMeHzX7SEhINPtmIoACYzJKYiInCmI0+gseIQUjHh81+0hISDT7ZiInCDIHZjMDEjMDCAg1NTYcATEAE0Q0Z0EAIiI0Z3AARDUGIhw0Bgk0Pwg1BDYaACcGEkABEDRnMwQREkQ2GgAnBxJAAE0zBBI0Rx00BCMdH0hITEhJNRA0NAk1yTMEEjRIHTQEIx0fSEhMSEk1ETQ1CTXKNBA0ERBENEc0EAk1UTRINBEJNVI0BDMEEgk1U0IB+jRHNDQINVE0SDQ1CDVSNAQiEkAALjQ0NAQdNEcjHR9ISExINDU0BB00SCMdH0hITEhKDU1JNAQINVMzBBIJNctCAbciJwUzBBFJNWdmKDRnFlA1cSI0Z3AAREQ0ZzRlE0Q0ZzRmE0QzBBIkCEkdNfA0NDQ1HTXxSgxAAAgSRDTwNPEORDMEEiQIIwhJHTXwNDQ0NR018UoNQAAIEkQ08DTxDUQkNT80BDMEEiQICDVTQgFHMwIRNGUSMwMRNGYSEEk1ZEAAGTMCETRmEjMDETRlEhBENEg1EjRHNRNCAAg0RzUSNEg1EzYaAYACZmkSQABaNhoBgAJmbxJENDUkCzQSHTQTNDUJJR0fSEhMSCMISTUVIg00NTQTDBBENDQ0FQk0ZEEAEzXJNEc0FQg1UTRINDUJNVJCAGc1yjRINBUINVI0RzQ1CTVRQgBUNDRJNRUlCzQTHTQSJAs0NCULHh9ISExISTUUIg00FDQTDBBENBQ0NQk0ZEEAEzXKNEc0NAg1UTRINBQJNVJCABM1yTRHNBQJNVE0SDQ0CDVSQgAANBUhBAs0BB2BoJwBNBIdH0hITEhJNSo0BAg1U0IAOyIrNhoBF0k1ZWYiJwQ2GgIXSTVmZjRlcQNEgAEtUIAEQUxHTzRmQQAGSDRmcQNEUDMCJkkVgQ1MUhJDIio0QDQqCGYiNHE0PzQqCDTLCGYiNG80PTTJCGYiNHA0PjTKCGYigAJzMTRRZiKAAnMyNFJmIicJNFIhBh00USMdH0hITEhmIicKNFEhBh00UiMdH0hITEhmIoADaWx0NFNmNMtBAAkjNHtKYjTLCGY0yUEACSM0eUpiNMkIZjTKQQAJIzR6SmI0yghmI0MjQyJD",
        "address": "PWOZOLBQJ5IMSYIFBZH5YOCP4NQ72MLWN3VD37WMWG2HHBEEIZJSJDRLFY",
        "size": 1296,
        "variables": [],
        "source": "https://github.com/tinymanorg/tinyman-contracts-v1/tree/6e150df8e3e74458f947f3399c6f37e7d4289d21/contracts/validator_approval.teal"
      },
      "clear_program": {
        "bytecode": "BIEB",
        "address": "P7GEWDXXW5IONRW6XRIRVPJCT2XXEQGOBGG65VJPBUOYZEJCBZWTPHS3VQ",
        "size": 3,
        "variables": [],
        "source": "https://github.com/tinymanorg/tinyman-contracts-v1/tree/6e150df8e3e74458f947f3399c6f37e7d4289d21/contracts/validator_clear_state.teal"
      },
      "global_state_schema": {
        "num_uints": 0,
        "num_byte_slices": 0
      },
      "local_state_schema": {
        "num_uints": 16,
        "num_byte_slices": 0
      },
      "name": "validator_app"
    }
  }
}
`)
