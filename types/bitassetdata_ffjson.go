// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: bitassetdata.go

package types

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"reflect"
)

// MarshalJSON marshal bytes to json - template
func (j *BitAssetData) MarshalJSON() ([]byte, error) {
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
func (j *BitAssetData) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
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
	buf.WriteString(`,"current_feed_publication_time":`)

	{

		obj, err = j.MembershipExpirationDate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	if j.IsPredictionMarket {
		buf.WriteString(`,"is_prediction_market":true`)
	} else {
		buf.WriteString(`,"is_prediction_market":false`)
	}
	/* Struct fall back. type=types.Price kind=struct */
	buf.WriteString(`,"settlement_price":`)
	err = buf.Encode(&j.SettlementPrice)
	if err != nil {
		return err
	}
	buf.WriteString(`,"feeds":`)
	if j.Feeds != nil {
		buf.WriteString(`[`)
		for i, v := range j.Feeds {
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
	buf.WriteString(`,"options":`)

	{

		err = j.Options.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	/* Struct fall back. type=types.PriceFeed kind=struct */
	buf.WriteString(`,"current_feed":`)
	err = buf.Encode(&j.CurrentFeed)
	if err != nil {
		return err
	}
	buf.WriteString(`,"force_settled_volume":`)
	fflib.FormatBits2(buf, uint64(j.ForcedSettledVolume), 10, false)
	buf.WriteString(`,"settlement_fund":`)
	fflib.FormatBits2(buf, uint64(j.SettlementFund), 10, false)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtBitAssetDatabase = iota
	ffjtBitAssetDatanosuchkey

	ffjtBitAssetDataID

	ffjtBitAssetDataMembershipExpirationDate

	ffjtBitAssetDataIsPredictionMarket

	ffjtBitAssetDataSettlementPrice

	ffjtBitAssetDataFeeds

	ffjtBitAssetDataOptions

	ffjtBitAssetDataCurrentFeed

	ffjtBitAssetDataForcedSettledVolume

	ffjtBitAssetDataSettlementFund
)

var ffjKeyBitAssetDataID = []byte("id")

var ffjKeyBitAssetDataMembershipExpirationDate = []byte("current_feed_publication_time")

var ffjKeyBitAssetDataIsPredictionMarket = []byte("is_prediction_market")

var ffjKeyBitAssetDataSettlementPrice = []byte("settlement_price")

var ffjKeyBitAssetDataFeeds = []byte("feeds")

var ffjKeyBitAssetDataOptions = []byte("options")

var ffjKeyBitAssetDataCurrentFeed = []byte("current_feed")

var ffjKeyBitAssetDataForcedSettledVolume = []byte("force_settled_volume")

var ffjKeyBitAssetDataSettlementFund = []byte("settlement_fund")

// UnmarshalJSON umarshall json - template of ffjson
func (j *BitAssetData) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *BitAssetData) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtBitAssetDatabase
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
				currentKey = ffjtBitAssetDatanosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyBitAssetDataMembershipExpirationDate, kn) {
						currentKey = ffjtBitAssetDataMembershipExpirationDate
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataCurrentFeed, kn) {
						currentKey = ffjtBitAssetDataCurrentFeed
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyBitAssetDataFeeds, kn) {
						currentKey = ffjtBitAssetDataFeeds
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataForcedSettledVolume, kn) {
						currentKey = ffjtBitAssetDataForcedSettledVolume
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyBitAssetDataID, kn) {
						currentKey = ffjtBitAssetDataID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataIsPredictionMarket, kn) {
						currentKey = ffjtBitAssetDataIsPredictionMarket
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyBitAssetDataOptions, kn) {
						currentKey = ffjtBitAssetDataOptions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyBitAssetDataSettlementPrice, kn) {
						currentKey = ffjtBitAssetDataSettlementPrice
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitAssetDataSettlementFund, kn) {
						currentKey = ffjtBitAssetDataSettlementFund
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataSettlementFund, kn) {
					currentKey = ffjtBitAssetDataSettlementFund
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataForcedSettledVolume, kn) {
					currentKey = ffjtBitAssetDataForcedSettledVolume
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBitAssetDataCurrentFeed, kn) {
					currentKey = ffjtBitAssetDataCurrentFeed
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataOptions, kn) {
					currentKey = ffjtBitAssetDataOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataFeeds, kn) {
					currentKey = ffjtBitAssetDataFeeds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataSettlementPrice, kn) {
					currentKey = ffjtBitAssetDataSettlementPrice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitAssetDataIsPredictionMarket, kn) {
					currentKey = ffjtBitAssetDataIsPredictionMarket
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyBitAssetDataMembershipExpirationDate, kn) {
					currentKey = ffjtBitAssetDataMembershipExpirationDate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBitAssetDataID, kn) {
					currentKey = ffjtBitAssetDataID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtBitAssetDatanosuchkey
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

				case ffjtBitAssetDataID:
					goto handle_ID

				case ffjtBitAssetDataMembershipExpirationDate:
					goto handle_MembershipExpirationDate

				case ffjtBitAssetDataIsPredictionMarket:
					goto handle_IsPredictionMarket

				case ffjtBitAssetDataSettlementPrice:
					goto handle_SettlementPrice

				case ffjtBitAssetDataFeeds:
					goto handle_Feeds

				case ffjtBitAssetDataOptions:
					goto handle_Options

				case ffjtBitAssetDataCurrentFeed:
					goto handle_CurrentFeed

				case ffjtBitAssetDataForcedSettledVolume:
					goto handle_ForcedSettledVolume

				case ffjtBitAssetDataSettlementFund:
					goto handle_SettlementFund

				case ffjtBitAssetDatanosuchkey:
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

	/* handler: j.ID type=types.AssetBitAssetDataID kind=struct quoted=false*/

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

handle_MembershipExpirationDate:

	/* handler: j.MembershipExpirationDate type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MembershipExpirationDate.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IsPredictionMarket:

	/* handler: j.IsPredictionMarket type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.IsPredictionMarket = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.IsPredictionMarket = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SettlementPrice:

	/* handler: j.SettlementPrice type=types.Price kind=struct quoted=false*/

	{
		/* Falling back. type=types.Price kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.SettlementPrice)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Feeds:

	/* handler: j.Feeds type=types.AssetFeeds kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AssetFeeds", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Feeds = nil
		} else {

			j.Feeds = []AssetFeed{}

			wantVal := true

			for {

				var tmpJFeeds AssetFeed

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

				/* handler: tmpJFeeds type=types.AssetFeed kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJFeeds.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Feeds = append(j.Feeds, tmpJFeeds)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Options:

	/* handler: j.Options type=types.BitassetOptions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			err = j.Options.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CurrentFeed:

	/* handler: j.CurrentFeed type=types.PriceFeed kind=struct quoted=false*/

	{
		/* Falling back. type=types.PriceFeed kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.CurrentFeed)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ForcedSettledVolume:

	/* handler: j.ForcedSettledVolume type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ForcedSettledVolume.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SettlementFund:

	/* handler: j.SettlementFund type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.SettlementFund.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
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

// MarshalJSON marshal bytes to json - template
func (j *BitassetOptions) MarshalJSON() ([]byte, error) {
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
func (j *BitassetOptions) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"feed_lifetime_sec":`)
	fflib.FormatBits2(buf, uint64(j.FeedLifetimeSec), 10, false)
	buf.WriteString(`,"minimum_feeds":`)
	fflib.FormatBits2(buf, uint64(j.MinimumFeeds), 10, false)
	buf.WriteString(`,"force_settlement_delay_sec":`)
	fflib.FormatBits2(buf, uint64(j.ForceSettlementDelaySec), 10, false)
	buf.WriteString(`,"force_settlement_offset_percent":`)
	fflib.FormatBits2(buf, uint64(j.ForceSettlementOffsetPercent), 10, false)
	buf.WriteString(`,"maximum_force_settlement_volume":`)
	fflib.FormatBits2(buf, uint64(j.MaximumForceSettlementVolume), 10, false)
	buf.WriteString(`,"short_backing_asset":`)

	{

		obj, err = j.ShortBackingAsset.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

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
	buf.WriteByte('}')
	return nil
}

const (
	ffjtBitassetOptionsbase = iota
	ffjtBitassetOptionsnosuchkey

	ffjtBitassetOptionsFeedLifetimeSec

	ffjtBitassetOptionsMinimumFeeds

	ffjtBitassetOptionsForceSettlementDelaySec

	ffjtBitassetOptionsForceSettlementOffsetPercent

	ffjtBitassetOptionsMaximumForceSettlementVolume

	ffjtBitassetOptionsShortBackingAsset

	ffjtBitassetOptionsExtensions
)

var ffjKeyBitassetOptionsFeedLifetimeSec = []byte("feed_lifetime_sec")

var ffjKeyBitassetOptionsMinimumFeeds = []byte("minimum_feeds")

var ffjKeyBitassetOptionsForceSettlementDelaySec = []byte("force_settlement_delay_sec")

var ffjKeyBitassetOptionsForceSettlementOffsetPercent = []byte("force_settlement_offset_percent")

var ffjKeyBitassetOptionsMaximumForceSettlementVolume = []byte("maximum_force_settlement_volume")

var ffjKeyBitassetOptionsShortBackingAsset = []byte("short_backing_asset")

var ffjKeyBitassetOptionsExtensions = []byte("extensions")

// UnmarshalJSON umarshall json - template of ffjson
func (j *BitassetOptions) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *BitassetOptions) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtBitassetOptionsbase
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
				currentKey = ffjtBitassetOptionsnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffjKeyBitassetOptionsExtensions, kn) {
						currentKey = ffjtBitassetOptionsExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyBitassetOptionsFeedLifetimeSec, kn) {
						currentKey = ffjtBitassetOptionsFeedLifetimeSec
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitassetOptionsForceSettlementDelaySec, kn) {
						currentKey = ffjtBitassetOptionsForceSettlementDelaySec
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitassetOptionsForceSettlementOffsetPercent, kn) {
						currentKey = ffjtBitassetOptionsForceSettlementOffsetPercent
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyBitassetOptionsMinimumFeeds, kn) {
						currentKey = ffjtBitassetOptionsMinimumFeeds
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBitassetOptionsMaximumForceSettlementVolume, kn) {
						currentKey = ffjtBitassetOptionsMaximumForceSettlementVolume
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyBitassetOptionsShortBackingAsset, kn) {
						currentKey = ffjtBitassetOptionsShortBackingAsset
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyBitassetOptionsExtensions, kn) {
					currentKey = ffjtBitassetOptionsExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitassetOptionsShortBackingAsset, kn) {
					currentKey = ffjtBitassetOptionsShortBackingAsset
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitassetOptionsMaximumForceSettlementVolume, kn) {
					currentKey = ffjtBitassetOptionsMaximumForceSettlementVolume
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitassetOptionsForceSettlementOffsetPercent, kn) {
					currentKey = ffjtBitassetOptionsForceSettlementOffsetPercent
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitassetOptionsForceSettlementDelaySec, kn) {
					currentKey = ffjtBitassetOptionsForceSettlementDelaySec
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitassetOptionsMinimumFeeds, kn) {
					currentKey = ffjtBitassetOptionsMinimumFeeds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBitassetOptionsFeedLifetimeSec, kn) {
					currentKey = ffjtBitassetOptionsFeedLifetimeSec
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtBitassetOptionsnosuchkey
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

				case ffjtBitassetOptionsFeedLifetimeSec:
					goto handle_FeedLifetimeSec

				case ffjtBitassetOptionsMinimumFeeds:
					goto handle_MinimumFeeds

				case ffjtBitassetOptionsForceSettlementDelaySec:
					goto handle_ForceSettlementDelaySec

				case ffjtBitassetOptionsForceSettlementOffsetPercent:
					goto handle_ForceSettlementOffsetPercent

				case ffjtBitassetOptionsMaximumForceSettlementVolume:
					goto handle_MaximumForceSettlementVolume

				case ffjtBitassetOptionsShortBackingAsset:
					goto handle_ShortBackingAsset

				case ffjtBitassetOptionsExtensions:
					goto handle_Extensions

				case ffjtBitassetOptionsnosuchkey:
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

handle_FeedLifetimeSec:

	/* handler: j.FeedLifetimeSec type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.FeedLifetimeSec.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MinimumFeeds:

	/* handler: j.MinimumFeeds type=types.UInt8 kind=uint8 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MinimumFeeds.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ForceSettlementDelaySec:

	/* handler: j.ForceSettlementDelaySec type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ForceSettlementDelaySec.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ForceSettlementOffsetPercent:

	/* handler: j.ForceSettlementOffsetPercent type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ForceSettlementOffsetPercent.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MaximumForceSettlementVolume:

	/* handler: j.MaximumForceSettlementVolume type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.MaximumForceSettlementVolume.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ShortBackingAsset:

	/* handler: j.ShortBackingAsset type=types.AssetID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ShortBackingAsset.UnmarshalJSON(tbuf)
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
