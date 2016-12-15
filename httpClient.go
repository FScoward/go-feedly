package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func get(client *http.Client, url string, accessToken string, typedef interface{}) interface{} {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", accessToken)
	resp, _ := client.Do(req)
	error := json.NewDecoder(resp.Body).Decode(&typedef)
	if error != nil {
		log.Fatal(error)
	}

	return typedef
}
