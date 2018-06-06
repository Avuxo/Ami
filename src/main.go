package main

import (
	"fmt"
)


func main() {
	fmt.Println(":: AMI ::")

	// load the config file
	config := parseConfigFile();

	fmt.Println(config) // example print file to show it
	
	//fmt.Println(fetchMangaInfo(15125))
	fetchAnimeList("benwhom")
}
