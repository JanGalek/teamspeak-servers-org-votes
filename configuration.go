package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Connection struct {
	Host     string
	Port     string
	Database string
	UserName string
	Password string
}

type Api struct {
	Key string
}

type Configuration struct {
	Connection Connection
	Api Api
}

func loadConfiguration() (*Configuration)  {
	configuration := new(Configuration)

	file, err := os.Open("./config/db.yaml")

	checkError(err)

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&configuration)

	checkError(err)
	return configuration
}

func getApi() (Api)  {
	configuration := loadConfiguration()

	return configuration.Api
}

func getConnection() (Connection)  {
	configuration := loadConfiguration()

	return configuration.Connection
}