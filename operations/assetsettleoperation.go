package operations

//go:generate ffjson   $GOFILE

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	op := &AssetSettleOperation{}
	types.OperationMap[op.Type()] = op
}

type AssetSettleOperation struct {
	Account    types.GrapheneID  `json:"account"`
	Amount     types.AssetAmount `json:"amount"`
	Fee        types.AssetAmount `json:"fee"`
	Extensions types.Extensions  `json:"extensions"`
}

func (p *AssetSettleOperation) ApplyFee(fee types.AssetAmount) {
	p.Fee = fee
}

func (p AssetSettleOperation) Type() types.OperationType {
	return types.OperationTypeAssetSettle
}

func (p AssetSettleOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode operation id")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode Account")
	}

	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode Amount")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode Extensions")
	}

	return nil
}

//NewAssetSettleOperation creates a new AssetSettleOperation
func NewAssetSettleOperation() *AssetSettleOperation {
	tx := AssetSettleOperation{
		Extensions: types.Extensions{},
	}
	return &tx
}