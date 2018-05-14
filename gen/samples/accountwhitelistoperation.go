//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeAccountWhitelist

package samples

import(
    "github.com/denkhaus/bitshares/gen/data"
    "github.com/denkhaus/bitshares/types"
)

func init(){
	data.OpSampleMap[types.OperationTypeAccountWhitelist] = 
    sampleDataAccountWhitelistOperation
}

var sampleDataAccountWhitelistOperation = `{
  "account_to_list": "1.2.96537",
  "authorizing_account": "1.2.35824",
  "extensions": [],
  "fee": {
    "amount": 600000,
    "asset_id": "1.3.0"
  },
  "new_listing": 2
}`

//end of file