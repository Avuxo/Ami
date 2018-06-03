package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)
/*

Config
This code deals with the loading and parsing of the configuration files
that handle user config. This is where things like API keys are loaded
from `config.json'.

*/

// structure for the config file.
type Config struct {
	ApiKey string `json:"apiKey"`
}

// read the config file into a string
func parseConfigFile() (Config) {
	// Load the `config.json' configuration file and decode it
	// into an internal go structure.
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()
	
	configuration := Config{}
	
	// read the file into memory and convert it into an internal struct.
	fileBytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(fileBytes, &configuration)
	
	return configuration
}
