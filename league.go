package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func httpCall(url string) (response []byte) {
	method := "GET"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, nil)
	check(err)
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	check(err)
	return body
}

func main() {

	playerScore := string("https://127.0.0.1:2999/liveclientdata/playerscores?summonerName=TrollAddict")
	allGameData := string("https://127.0.0.1:2999/liveclientdata/allgamedata")
	allGameDataResponse := httpCall(allGameData)
	var allGameDat map[string]interface{}
	if err := json.Unmarshal(allGameDataResponse, &allGameDat); err != nil {
		panic(err)
	}
	response := []byte(httpCall(playerScore))

	fmt.Println(response)
	var dat map[string]interface{}
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}
	file, _ := json.MarshalIndent(allGameDat, "", " ")
	_ = ioutil.WriteFile("test2.json", file, 0644)
}
