//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAssetCreate

package samples

import(
    "github.com/denkhaus/bitshares/gen/data"
    "github.com/denkhaus/bitshares/types"
)

func init(){
	data.OpSampleMap[types.OperationTypeAssetCreate] = 
    sampleDataAssetCreateOperation
}

var sampleDataAssetCreateOperation = `{
  "common_options": {
    "blacklist_authorities": [],
    "blacklist_markets": [],
    "core_exchange_rate": {
      "base": {
        "amount": 1,
        "asset_id": "1.3.0"
      },
      "quote": {
        "amount": 1,
        "asset_id": "1.3.1"
      }
    },
    "description": "Open Ledger Asset Namespace",
    "extensions": [],
    "flags": 0,
    "issuer_permissions": 79,
    "market_fee_percent": 0,
    "max_market_fee": 0,
    "max_supply": "1000000000000000",
    "whitelist_authorities": [],
    "whitelist_markets": []
  },
  "extensions": [],
  "fee": {
    "amount": 1000000000,
    "asset_id": "1.3.0"
  },
  "is_prediction_market": false,
  "issuer": "1.2.1090",
  "precision": 5,
  "symbol": "OPEN"
}`

//end of file