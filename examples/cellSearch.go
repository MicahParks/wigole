package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/MicahParks/wigole"
	"github.com/MicahParks/wigole/api/cell/search"
)

type creds struct {
	ApiName  string
	ApiToken string
}

func main() {
	cred := creds{}
	c, err := ioutil.ReadFile("creds.json")
	if err != nil {
		println("Create a properly formatted 'creds.json' file in the working directory.")
		return
	}
	if err = json.Unmarshal(c, &cred); err != nil {
		println("JSON failure for 'creds.json' file in the working directory.")
		return
	}
	u := wigole.NewUser(cred.ApiName, cred.ApiToken)
	s := search.New()
	s.Latrange1 = 37.0078
	s.Latrange2 = 38.0348
	s.Longrange1 = -122.6535
	s.Longrange2 = -121.3351
	resp, err := s.Do(u)
	if err != nil {
		if errors.Is(err, wigole.ErrFail) {
			println(err.Error())
			return
		}
		if err == wigole.ErrAuth {
			println("Failed to authenticate with creds.json.")
			return
		}
		panic(err)
	}
	println(resp.Results[0].Postalcode)
}
