// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: proposalupdateoperation.go

package operations

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/denkhaus/bitshares/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"reflect"
)

// MarshalJSON marshal bytes to json - template
func (j *ProposalUpdateOperation) MarshalJSON() ([]byte, error) {
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
func (j *ProposalUpdateOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "active_approvals_to_add":`)
	if j.ActiveApprovalsToAdd != nil {
		buf.WriteString(`[`)
		for i, v := range j.ActiveApprovalsToAdd {
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
	buf.WriteString(`,"active_approvals_to_remove":`)
	if j.ActiveApprovalsToRemove != nil {
		buf.WriteString(`[`)
		for i, v := range j.ActiveApprovalsToRemove {
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
	buf.WriteString(`,"owner_approvals_to_add":`)
	if j.OwnerApprovalsToAdd != nil {
		buf.WriteString(`[`)
		for i, v := range j.OwnerApprovalsToAdd {
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
	buf.WriteString(`,"owner_approvals_to_remove":`)
	if j.OwnerApprovalsToRemove != nil {
		buf.WriteString(`[`)
		for i, v := range j.OwnerApprovalsToRemove {
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
	buf.WriteString(`,"extensions":`)
	if j.Extensions != nil {
		buf.WriteString(`"`)
		{
			enc := base64.NewEncoder(base64.StdEncoding, buf)
			enc.Write(reflect.Indirect(reflect.ValueOf(j.Extensions)).Bytes())
			enc.Close()
		}
		buf.WriteString(`"`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"fee_paying_account":`)

	{

		obj, err = j.FeePayingAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"key_approvals_to_add":`)
	if j.KeyApprovalsToAdd != nil {
		buf.WriteString(`[`)
		for i, v := range j.KeyApprovalsToAdd {
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
	buf.WriteString(`,"key_approvals_to_remove":`)
	if j.KeyApprovalsToRemove != nil {
		buf.WriteString(`[`)
		for i, v := range j.KeyApprovalsToRemove {
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
	buf.WriteString(`,"proposal":`)

	{

		obj, err = j.Proposal.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

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
	ffjtProposalUpdateOperationbase = iota
	ffjtProposalUpdateOperationnosuchkey

	ffjtProposalUpdateOperationActiveApprovalsToAdd

	ffjtProposalUpdateOperationActiveApprovalsToRemove

	ffjtProposalUpdateOperationOwnerApprovalsToAdd

	ffjtProposalUpdateOperationOwnerApprovalsToRemove

	ffjtProposalUpdateOperationExtensions

	ffjtProposalUpdateOperationFeePayingAccount

	ffjtProposalUpdateOperationKeyApprovalsToAdd

	ffjtProposalUpdateOperationKeyApprovalsToRemove

	ffjtProposalUpdateOperationProposal

	ffjtProposalUpdateOperationFee
)

var ffjKeyProposalUpdateOperationActiveApprovalsToAdd = []byte("active_approvals_to_add")

var ffjKeyProposalUpdateOperationActiveApprovalsToRemove = []byte("active_approvals_to_remove")

var ffjKeyProposalUpdateOperationOwnerApprovalsToAdd = []byte("owner_approvals_to_add")

var ffjKeyProposalUpdateOperationOwnerApprovalsToRemove = []byte("owner_approvals_to_remove")

var ffjKeyProposalUpdateOperationExtensions = []byte("extensions")

var ffjKeyProposalUpdateOperationFeePayingAccount = []byte("fee_paying_account")

var ffjKeyProposalUpdateOperationKeyApprovalsToAdd = []byte("key_approvals_to_add")

var ffjKeyProposalUpdateOperationKeyApprovalsToRemove = []byte("key_approvals_to_remove")

var ffjKeyProposalUpdateOperationProposal = []byte("proposal")

var ffjKeyProposalUpdateOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *ProposalUpdateOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *ProposalUpdateOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtProposalUpdateOperationbase
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
				currentKey = ffjtProposalUpdateOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyProposalUpdateOperationActiveApprovalsToAdd, kn) {
						currentKey = ffjtProposalUpdateOperationActiveApprovalsToAdd
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProposalUpdateOperationActiveApprovalsToRemove, kn) {
						currentKey = ffjtProposalUpdateOperationActiveApprovalsToRemove
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyProposalUpdateOperationExtensions, kn) {
						currentKey = ffjtProposalUpdateOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyProposalUpdateOperationFeePayingAccount, kn) {
						currentKey = ffjtProposalUpdateOperationFeePayingAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProposalUpdateOperationFee, kn) {
						currentKey = ffjtProposalUpdateOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'k':

					if bytes.Equal(ffjKeyProposalUpdateOperationKeyApprovalsToAdd, kn) {
						currentKey = ffjtProposalUpdateOperationKeyApprovalsToAdd
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProposalUpdateOperationKeyApprovalsToRemove, kn) {
						currentKey = ffjtProposalUpdateOperationKeyApprovalsToRemove
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyProposalUpdateOperationOwnerApprovalsToAdd, kn) {
						currentKey = ffjtProposalUpdateOperationOwnerApprovalsToAdd
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProposalUpdateOperationOwnerApprovalsToRemove, kn) {
						currentKey = ffjtProposalUpdateOperationOwnerApprovalsToRemove
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyProposalUpdateOperationProposal, kn) {
						currentKey = ffjtProposalUpdateOperationProposal
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyProposalUpdateOperationFee, kn) {
					currentKey = ffjtProposalUpdateOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationProposal, kn) {
					currentKey = ffjtProposalUpdateOperationProposal
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationKeyApprovalsToRemove, kn) {
					currentKey = ffjtProposalUpdateOperationKeyApprovalsToRemove
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationKeyApprovalsToAdd, kn) {
					currentKey = ffjtProposalUpdateOperationKeyApprovalsToAdd
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyProposalUpdateOperationFeePayingAccount, kn) {
					currentKey = ffjtProposalUpdateOperationFeePayingAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationExtensions, kn) {
					currentKey = ffjtProposalUpdateOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationOwnerApprovalsToRemove, kn) {
					currentKey = ffjtProposalUpdateOperationOwnerApprovalsToRemove
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationOwnerApprovalsToAdd, kn) {
					currentKey = ffjtProposalUpdateOperationOwnerApprovalsToAdd
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationActiveApprovalsToRemove, kn) {
					currentKey = ffjtProposalUpdateOperationActiveApprovalsToRemove
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalUpdateOperationActiveApprovalsToAdd, kn) {
					currentKey = ffjtProposalUpdateOperationActiveApprovalsToAdd
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtProposalUpdateOperationnosuchkey
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

				case ffjtProposalUpdateOperationActiveApprovalsToAdd:
					goto handle_ActiveApprovalsToAdd

				case ffjtProposalUpdateOperationActiveApprovalsToRemove:
					goto handle_ActiveApprovalsToRemove

				case ffjtProposalUpdateOperationOwnerApprovalsToAdd:
					goto handle_OwnerApprovalsToAdd

				case ffjtProposalUpdateOperationOwnerApprovalsToRemove:
					goto handle_OwnerApprovalsToRemove

				case ffjtProposalUpdateOperationExtensions:
					goto handle_Extensions

				case ffjtProposalUpdateOperationFeePayingAccount:
					goto handle_FeePayingAccount

				case ffjtProposalUpdateOperationKeyApprovalsToAdd:
					goto handle_KeyApprovalsToAdd

				case ffjtProposalUpdateOperationKeyApprovalsToRemove:
					goto handle_KeyApprovalsToRemove

				case ffjtProposalUpdateOperationProposal:
					goto handle_Proposal

				case ffjtProposalUpdateOperationFee:
					goto handle_Fee

				case ffjtProposalUpdateOperationnosuchkey:
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

handle_ActiveApprovalsToAdd:

	/* handler: j.ActiveApprovalsToAdd type=types.AccountIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AccountIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ActiveApprovalsToAdd = nil
		} else {

			j.ActiveApprovalsToAdd = []types.AccountID{}

			wantVal := true

			for {

				var tmpJActiveApprovalsToAdd types.AccountID

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

				/* handler: tmpJActiveApprovalsToAdd type=types.AccountID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJActiveApprovalsToAdd.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.ActiveApprovalsToAdd = append(j.ActiveApprovalsToAdd, tmpJActiveApprovalsToAdd)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ActiveApprovalsToRemove:

	/* handler: j.ActiveApprovalsToRemove type=types.AccountIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AccountIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ActiveApprovalsToRemove = nil
		} else {

			j.ActiveApprovalsToRemove = []types.AccountID{}

			wantVal := true

			for {

				var tmpJActiveApprovalsToRemove types.AccountID

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

				/* handler: tmpJActiveApprovalsToRemove type=types.AccountID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJActiveApprovalsToRemove.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.ActiveApprovalsToRemove = append(j.ActiveApprovalsToRemove, tmpJActiveApprovalsToRemove)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OwnerApprovalsToAdd:

	/* handler: j.OwnerApprovalsToAdd type=types.AccountIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AccountIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.OwnerApprovalsToAdd = nil
		} else {

			j.OwnerApprovalsToAdd = []types.AccountID{}

			wantVal := true

			for {

				var tmpJOwnerApprovalsToAdd types.AccountID

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

				/* handler: tmpJOwnerApprovalsToAdd type=types.AccountID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJOwnerApprovalsToAdd.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.OwnerApprovalsToAdd = append(j.OwnerApprovalsToAdd, tmpJOwnerApprovalsToAdd)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OwnerApprovalsToRemove:

	/* handler: j.OwnerApprovalsToRemove type=types.AccountIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AccountIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.OwnerApprovalsToRemove = nil
		} else {

			j.OwnerApprovalsToRemove = []types.AccountID{}

			wantVal := true

			for {

				var tmpJOwnerApprovalsToRemove types.AccountID

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

				/* handler: tmpJOwnerApprovalsToRemove type=types.AccountID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJOwnerApprovalsToRemove.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.OwnerApprovalsToRemove = append(j.OwnerApprovalsToRemove, tmpJOwnerApprovalsToRemove)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Extensions", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Extensions = nil
		} else {
			b := make([]byte, base64.StdEncoding.DecodedLen(fs.Output.Len()))
			n, err := base64.StdEncoding.Decode(b, fs.Output.Bytes())
			if err != nil {
				return fs.WrapErr(err)
			}

			v := reflect.ValueOf(&j.Extensions).Elem()
			v.SetBytes(b[0:n])

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

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

handle_KeyApprovalsToAdd:

	/* handler: j.KeyApprovalsToAdd type=types.PublicKeys kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for PublicKeys", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.KeyApprovalsToAdd = nil
		} else {

			j.KeyApprovalsToAdd = []types.PublicKey{}

			wantVal := true

			for {

				var tmpJKeyApprovalsToAdd types.PublicKey

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

				/* handler: tmpJKeyApprovalsToAdd type=types.PublicKey kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJKeyApprovalsToAdd.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.KeyApprovalsToAdd = append(j.KeyApprovalsToAdd, tmpJKeyApprovalsToAdd)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_KeyApprovalsToRemove:

	/* handler: j.KeyApprovalsToRemove type=types.PublicKeys kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for PublicKeys", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.KeyApprovalsToRemove = nil
		} else {

			j.KeyApprovalsToRemove = []types.PublicKey{}

			wantVal := true

			for {

				var tmpJKeyApprovalsToRemove types.PublicKey

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

				/* handler: tmpJKeyApprovalsToRemove type=types.PublicKey kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJKeyApprovalsToRemove.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.KeyApprovalsToRemove = append(j.KeyApprovalsToRemove, tmpJKeyApprovalsToRemove)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Proposal:

	/* handler: j.Proposal type=types.ProposalID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Proposal.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
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
