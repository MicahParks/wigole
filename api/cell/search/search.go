package search

import (
	"io"
	"net/url"
	"strconv"

	"gitlab.com/MicahParks/wigole"
)

const (
	ApiPath = "cell/search"
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
	if len(p.Cell_op) != 0 {
		values.Set("cell_op", string(p.Cell_op))
	}
	if len(p.Cell_net) != 0 {
		values.Set("cell_net", string(p.Cell_net))
	}
	if len(p.Cell_id) != 0 {
		values.Set("cell_id", string(p.Cell_id))
	}
	values.Set("showGsm", strconv.FormatBool(p.ShowGsm))
	values.Set("showCdma", strconv.FormatBool(p.ShowCdma))
	values.Set("showLte", strconv.FormatBool(p.ShowLte))
	values.Set("showWcdma", strconv.FormatBool(p.ShowWcdma))
	return values, nil
}

func (p *Parameters) Do(u *wigole.User) (*Response, error) {
	resp := &Response{}
	if err := wigole.Do(ApiPath, p, Method, resp, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	p := Parameters{
		ShowGsm:   true,
		ShowCdma:  true,
		ShowLte:   true,
		ShowWcdma: true,
	}
	p.SearchSsid = *wigole.NewSsid()
	return &p
}