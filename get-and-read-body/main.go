package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Element struct {
	Lyrics string `json:"lyrics"`
}

var client = http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	url := "https://api.lyrics.ovh/v1/kyuss/gardenia"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("invalid status code %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var e Element
	if err := json.Unmarshal(body, &e); err != nil {
		panic(err)
	}
	fmt.Println(e.Lyrics)
}
