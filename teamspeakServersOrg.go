package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getUrl(api Api, month string) (string)  {
	return fmt.Sprintf("https://teamspeak-servers.org/api/?object=servers&element=voters&key=%s&month=%s&limit=500&format=json", api.Key, month)
}

func loadData(month string) (Server)  {
	api := getApi()

	url := getUrl(api, month)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	rr := Server{}

	jsonError := json.Unmarshal(data, &rr)

	if jsonError != nil {
		log.Fatal(jsonError)
	}

	return rr
}