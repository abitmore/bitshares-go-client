// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: limitordercanceloperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/denkhaus/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *LimitOrderCancelOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *LimitOrderCancelOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "fee_paying_account":`)

	{

		obj, err = j.FeePayingAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"order":`)

	{

		obj, err = j.Order.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"extensions":`)
	if j.Extensions != nil {
		buf.WriteString(`[`)
		for i, v := range j.Extensions {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtLimitOrderCancelOperationbase = iota
	ffjtLimitOrderCancelOperationnosuchkey

	ffjtLimitOrderCancelOperationFeePayingAccount

	ffjtLimitOrderCancelOperationOrder

	ffjtLimitOrderCancelOperationExtensions

	ffjtLimitOrderCancelOperationFee
)

var ffjKeyLimitOrderCancelOperationFeePayingAccount = []byte("fee_paying_account")

var ffjKeyLimitOrderCancelOperationOrder = []byte("order")

var ffjKeyLimitOrderCancelOperationExtensions = []byte("extensions")

var ffjKeyLimitOrderCancelOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *LimitOrderCancelOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *LimitOrderCancelOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtLimitOrderCancelOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtLimitOrderCancelOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffjKeyLimitOrderCancelOperationExtensions, kn) {
						currentKey = ffjtLimitOrderCancelOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyLimitOrderCancelOperationFeePayingAccount, kn) {
						currentKey = ffjtLimitOrderCancelOperationFeePayingAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyLimitOrderCancelOperationFee, kn) {
						currentKey = ffjtLimitOrderCancelOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyLimitOrderCancelOperationOrder, kn) {
						currentKey = ffjtLimitOrderCancelOperationOrder
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyLimitOrderCancelOperationFee, kn) {
					currentKey = ffjtLimitOrderCancelOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyLimitOrderCancelOperationExtensions, kn) {
					currentKey = ffjtLimitOrderCancelOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyLimitOrderCancelOperationOrder, kn) {
					currentKey = ffjtLimitOrderCancelOperationOrder
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyLimitOrderCancelOperationFeePayingAccount, kn) {
					currentKey = ffjtLimitOrderCancelOperationFeePayingAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtLimitOrderCancelOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtLimitOrderCancelOperationFeePayingAccount:
					goto handle_FeePayingAccount

				case ffjtLimitOrderCancelOperationOrder:
					goto handle_Order

				case ffjtLimitOrderCancelOperationExtensions:
					goto handle_Extensions

				case ffjtLimitOrderCancelOperationFee:
					goto handle_Fee

				case ffjtLimitOrderCancelOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_FeePayingAccount:

	/* handler: j.FeePayingAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.FeePayingAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Order:

	/* handler: j.Order type=types.LimitOrderID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Order.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Extensions", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Extensions = nil
		} else {

			j.Extensions = []types.Extension{}

			wantVal := true

			for {

				var tmpJExtensions types.Extension

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJExtensions type=types.Extension kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJExtensions.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Extensions = append(j.Extensions, tmpJExtensions)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
