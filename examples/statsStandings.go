package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/MicahParks/wigole"
	"github.com/MicahParks/wigole/api/stats/standings"
)

type creds struct {
	ApiName  string
	ApiToken string
}

func main() {
	// Query for user standings. Print the first user's username and how many WiFi networks they've discovered.
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
	s := standings.New()
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
	println(resp.Results[0].UserName, resp.Results[0].DiscoveredWiFi)
}
