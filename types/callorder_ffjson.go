// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: callorder.go

package types

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CallOrder) MarshalJSON() ([]byte, error) {
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
func (j *CallOrder) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"borrower":`)

	{

		obj, err = j.Borrower.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"collateral":`)
	fflib.FormatBits2(buf, uint64(j.Collateral), 10, false)
	buf.WriteString(`,"debt":`)
	fflib.FormatBits2(buf, uint64(j.Debt), 10, false)
	buf.WriteString(`,"call_price":`)

	{

		err = j.CallPrice.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCallOrderbase = iota
	ffjtCallOrdernosuchkey

	ffjtCallOrderID

	ffjtCallOrderBorrower

	ffjtCallOrderCollateral

	ffjtCallOrderDebt

	ffjtCallOrderCallPrice
)

var ffjKeyCallOrderID = []byte("id")

var ffjKeyCallOrderBorrower = []byte("borrower")

var ffjKeyCallOrderCollateral = []byte("collateral")

var ffjKeyCallOrderDebt = []byte("debt")

var ffjKeyCallOrderCallPrice = []byte("call_price")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CallOrder) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CallOrder) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCallOrderbase
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
				currentKey = ffjtCallOrdernosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'b':

					if bytes.Equal(ffjKeyCallOrderBorrower, kn) {
						currentKey = ffjtCallOrderBorrower
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyCallOrderCollateral, kn) {
						currentKey = ffjtCallOrderCollateral
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyCallOrderCallPrice, kn) {
						currentKey = ffjtCallOrderCallPrice
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyCallOrderDebt, kn) {
						currentKey = ffjtCallOrderDebt
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyCallOrderID, kn) {
						currentKey = ffjtCallOrderID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffjKeyCallOrderCallPrice, kn) {
					currentKey = ffjtCallOrderCallPrice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCallOrderDebt, kn) {
					currentKey = ffjtCallOrderDebt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCallOrderCollateral, kn) {
					currentKey = ffjtCallOrderCollateral
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCallOrderBorrower, kn) {
					currentKey = ffjtCallOrderBorrower
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCallOrderID, kn) {
					currentKey = ffjtCallOrderID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtCallOrdernosuchkey
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

				case ffjtCallOrderID:
					goto handle_ID

				case ffjtCallOrderBorrower:
					goto handle_Borrower

				case ffjtCallOrderCollateral:
					goto handle_Collateral

				case ffjtCallOrderDebt:
					goto handle_Debt

				case ffjtCallOrderCallPrice:
					goto handle_CallPrice

				case ffjtCallOrdernosuchkey:
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

handle_ID:

	/* handler: j.ID type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Borrower:

	/* handler: j.Borrower type=types.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Borrower.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Collateral:

	/* handler: j.Collateral type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Collateral.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Debt:

	/* handler: j.Debt type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Debt.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CallPrice:

	/* handler: j.CallPrice type=types.Price kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.CallPrice.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
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