package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
)

// Create hash value from the string
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}


// Get FMG API Data
func CollectData(urlString string) string {
	apiKey := "70c55be6005581c638125a74ec3c5d54"
	baseURL := "https://financialmodelingprep.com/api/v3/"
	url := baseURL + urlString + apiKey
	//fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	return responseString
}