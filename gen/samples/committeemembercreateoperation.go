//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeCommitteeMemberCreate

package samples

import(
    "github.com/denkhaus/bitshares/gen/data"
    "github.com/denkhaus/bitshares/types"
)

func init(){
	data.OpSampleMap[types.OperationTypeCommitteeMemberCreate] = 
    sampleDataCommitteeMemberCreateOperation
}

var sampleDataCommitteeMemberCreateOperation = `{
  "committee_member_account": "1.2.1191",
  "fee": {
    "amount": 200000,
    "asset_id": "1.3.0"
  },
  "url": "https://github.com/gileadmcgee/dele-puppy"
}`

//end of file