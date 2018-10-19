package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getUrl(api Api) (string)  {
	return "https://teamspeak-servers.org/api/?object=servers&element=voters&key=" + api.Key + "&month=current&format=json"
}

func loadData() (Server)  {
	api := getApi()

	url := getUrl(api)
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