package search

import (
	"io"
	"net/url"
	"strconv"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "network/search"
	Method  = "GET"
)

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (values url.Values, err error) {
	values, err = p.SsidUrl()
	if err != nil {
		return url.Values{}, err
	}
	if len(p.Encryption) != 0 {
		// It's possible for a user of the API to make their own Encryption type, but we'll allow it.
		values.Set("encryption", string(p.Encryption))
	}
	values.Set("freenet", strconv.FormatBool(p.Freenet))
	values.Set("paynet", strconv.FormatBool(p.Paynet))
	if len(p.Netid) != 0 {
		values.Set("netid", p.Netid)
	}
	return values, nil
}

func (p *Parameters) Do(u *wigole.User) (*NetSearchResponse, error) {
	resp := &NetSearchResponse{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	p := Parameters{}
	p.SearchSsid = *wigole.NewSsid()
	return &p
}