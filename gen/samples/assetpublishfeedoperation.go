//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAssetPublishFeed

package samples

import(
    "github.com/denkhaus/bitshares/gen/data"
    "github.com/denkhaus/bitshares/types"
)

func init(){
	data.OpSampleMap[types.OperationTypeAssetPublishFeed] = 
    sampleDataAssetPublishFeedOperation
}

var sampleDataAssetPublishFeedOperation = `{
  "asset_id": "1.3.113",
  "extensions": [],
  "fee": {
    "amount": 200000,
    "asset_id": "1.3.0"
  },
  "feed": {
    "core_exchange_rate": {
      "base": {
        "amount": 2809162,
        "asset_id": "1.3.113"
      },
      "quote": {
        "amount": 1000000000,
        "asset_id": "1.3.0"
      }
    },
    "maintenance_collateral_ratio": 1750,
    "maximum_short_squeeze_ratio": 1100,
    "settlement_price": {
      "base": {
        "amount": 2809162,
        "asset_id": "1.3.113"
      },
      "quote": {
        "amount": 1000000000,
        "asset_id": "1.3.0"
      }
    }
  },
  "publisher": "1.2.35248"
}`

//end of file