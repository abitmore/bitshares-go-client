//This file is generated by btsgen. DO NOT EDIT.
//operation sample data for OperationTypeProposalCreate

package samples

func init() {

	sampleDataProposalCreateOperation[5] = `{
  "expiration_time": "2015-10-20T15:55:00",
  "extensions": [],
  "fee": {
    "amount": 2000005,
    "asset_id": "1.3.0"
  },
  "fee_paying_account": "1.2.90742",
  "proposed_ops": [
    {
      "op": [
        31,
        {
          "fee": {
            "amount": 2000000,
            "asset_id": "1.3.0"
          },
          "new_parameters": {     
            "extensions": [],       
            "current_fees": {
              "parameters": [
                [
                  0,
                  {
                    "fee": 2000000,
                    "price_per_kbyte": 1000000
                  }
                ],
                [
                  1,
                  {
                    "fee": 500000
                  }
                ],
                [
                  2,
                  {
                    "fee": 0
                  }
                ],
                [
                  3,
                  {
                    "fee": 2000000
                  }
                ],
                [
                  4,
                  {}
                ]
              ],
              "scale": 20000
            }            
          }
        }
      ]
    }
  ],
  "review_period_seconds": 3600
}`

}

//end of file