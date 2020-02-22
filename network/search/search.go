package search

import (
	"errors"
	"fmt"
	"io"

	"gitlab.com/MicahParks/wigole"
	"gitlab.com/MicahParks/wigole/date"
	"gitlab.com/MicahParks/wigole/user"
)

const ApiUrl = "network/search?"
const Method = "GET"

var errVariance = errors.New("variance must be between 0.001 and 0.2")

func (p *Parameters) Body() (io.Reader, error) {
	return nil, nil
}

func (p *Parameters) Url() (url string, err error) {
	url = fmt.Sprintf("onlymine=%v", p.Onlymine)
	if p.Notmine {
		url += fmt.Sprintf("&notmine=%v", p.Notmine)
	}
	if p.Latrange1 != 0 || p.Latrange2 != 0 {
		url += fmt.Sprintf("&latrange1=%f&latrange2=%f", p.Latrange1, p.Latrange2)
	}
	if p.Longrange1 != 0 || p.Longrange2 != 0 {
		url += fmt.Sprintf("&longrange1=%f&longrange2=%f", p.Longrange1, p.Longrange2)
	}
	if !p.Lastupdt.IsZero() {
		url += fmt.Sprintf("&lastupdt=%s", date.String(p.Lastupdt))
	}
	if !p.StartTransID.IsZero() {
		url += fmt.Sprintf("&startTransID=%d", p.StartTransID.Year()) // Will break in about 8000 years.
	}
	if !p.EndTransID.IsZero() {
		url += fmt.Sprintf("&endTransID=%d", p.EndTransID.Year()) // Will break in about 8000 years.
	}
	if len(p.Encryption) != 0 {
		// It's possible for a user of the API to make their own Encryption type, but we'll allow it.
		url += fmt.Sprintf("&encryption=%s", p.Encryption)
	}
	url += fmt.Sprintf("&freenet=%v", p.Freenet)
	url += fmt.Sprintf("&paynet=%v", p.Paynet)
	if len(p.Netid) != 0 {
		url += fmt.Sprintf("&netid=%s", p.Netid)
	}
	if len(p.Ssid) != 0 {
		url += fmt.Sprintf("&ssid=%s", p.Ssid)
	}
	if len(p.Ssidlike) != 0 {
		url += fmt.Sprintf("&ssidlike=%s", p.Ssidlike)
	}
	if p.MinQoS != 8 {
		url += fmt.Sprintf("&minQoS=%d", p.MinQoS)
	}
	if p.Variance < 0.001 && p.Variance > 0.2 {
		url += fmt.Sprintf("&variance=%f", p.Variance)
	} else if p.Variance != 0 {
		return "", errVariance
	}
	if len(p.HouseNumber) != 0 {
		url += fmt.Sprintf("&houseNumber=%s", p.HouseNumber)
	}
	if len(p.Road) != 0 {
		url += fmt.Sprintf("&road=%s", p.Road)
	}
	if len(p.City) != 0 {
		url += fmt.Sprintf("&city=%s", p.City)
	}
	if len(p.Region) != 0 {
		url += fmt.Sprintf("&region=%s", p.Region)
	}
	if len(p.PostalCode) != 0 {
		url += fmt.Sprintf("&postalCode=%s", p.PostalCode)
	}
	if len(p.Country) != 0 {
		url += fmt.Sprintf("&country=%s", p.Country)
	}
	if p.ResultsPerPage != 0 {
		url += fmt.Sprintf("&resultsPerPage=%d", p.ResultsPerPage)
	}
	if len(p.SearchAfter) != 0 {
		url += fmt.Sprintf("&searchAfter=%s", p.SearchAfter)
	}
	return url, nil
}

func (p *Parameters) Do(u *user.User) (*wigole.Response, error) {
	resp := &wigole.Response{}
	if err := wigole.Do(p, Method, resp, ApiUrl, u); err != nil {
		return nil, err
	}
	return resp, nil
}

func New() *Parameters {
	return &Parameters{
		MinQoS: 8, // Max of 7. This let's you know it's uninitialized.
	}
}