package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gitlab.com/MicahParks/wigole"
	pUser "gitlab.com/MicahParks/wigole/api/profile/user"
)

type creds struct {
	Password string
	Username string
}

func main() {
	// Get and print your user info.
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
	u := wigole.New(cred.Password, cred.Username)
	pU := pUser.New()
	resp, err := pU.Do(u)
	if err != nil {
		if err == wigole.ErrTooMany {
			println("Too many queries of that type for today.")
			return
		}
		if err == wigole.ErrAuth {
			println("Failed to authenticate with creds.json.")
			return
		}
		panic(err)
	}
	fmt.Printf("%+v\n", *resp)
}
