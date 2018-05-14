//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeTransfer

package samples

import(
    "github.com/denkhaus/bitshares/gen/data"
    "github.com/denkhaus/bitshares/types"
)

func init(){
	data.OpSampleMap[types.OperationTypeTransfer] = 
    sampleDataTransferOperation
}

var sampleDataTransferOperation = `{
  "amount": {
    "amount": 13795774,
    "asset_id": "1.3.350"
  },
  "extensions": [],
  "fee": {
    "amount": 4242186,
    "asset_id": "1.3.0"
  },
  "from": "1.2.95726",
  "memo": {
    "from": "BTS6GuuJFaPkr9E8Zh8PZ6VyCpe3A89uti2G38Dc6jN1gRaoR4uss",
    "message": "40b22e49bab023a2c460b3d8f2b2c4d0a748c52af2d712518b967ba4d51ba8e5a460271b83e880f1e30b34acc0ce5967",
    "nonce": "370035771605783",
    "to": "BTS8MmcVDiutGynSpi5vSr8tWbrTDWYWpAkTXUD24sJu45DBFLSRK"
  },
  "to": "1.2.32567"
}`

//end of file