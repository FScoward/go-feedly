package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	// feedly から更新情報を取得する
	fmt.Println("Hello go feedly")
	var config Config
	_, err := toml.DecodeFile("../config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}
	client := new(http.Client)
	//getCode(client)
	accessToken := config.Feedly.AccessToken
	//profile(client, accessToken)
	fetchCategories(client, accessToken)
	unreadCounts := getUnreadFeeds(client, accessToken)
	for i, obj:= range unreadCounts.Unreadcounts {
		fmt.Println(i, obj)
	}


}

func getCode(client *http.Client) {
	feedly_auth_url := "https://cloud.feedly.com/v3/auth/auth?client_id=feedly&redirect_uri=http://localhost&scope=https://cloud.feedly.com/subscriptions&response_type=code&provider=google&migrate=false"
	req, _ := http.NewRequest("GET", feedly_auth_url, nil)
	//req.Header.Set("Authorization OAuth", access_token)
	resp, _ := client.Do(req)
	dumpResp, _ := httputil.DumpResponse(resp, false)
	fmt.Printf("%s", dumpResp)
}

func authToken(client *http.Client, clientSecret string, code string) {
	fmt.Println("### auth token ###")
	url := "https://cloud.feedly.com/v3/auth/token" +
		"?client_id=feedly" +
		"&client_secret=" + clientSecret +
		"&grant_type=authorization_code" +
		"&redirect_uri=http%3A%2F%2Fwww.feedly.com%2Ffeedly.html" +
		"&code=" + code
	fmt.Println(url)
	resp, _ := client.PostForm(url, nil)
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
}

func profile(client *http.Client, accessToken string) {
	url := "https://cloud.feedly.com/v3/profile"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", accessToken)
	resp, _ := client.Do(req)
	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
}

func fetchCategories(client *http.Client, accessToken string) {
	url := "https://cloud.feedly.com/v3/categories"
	var cat []Category
	get(client, url, accessToken, &cat)
	for i, m := range cat {
		fmt.Println(i, m)
	}
}

func getUnreadFeeds(client *http.Client, accessToken string) *UnreadCounts {
	url := "https://cloud.feedly.com//v3/markers/counts"
	var unreadCounts = new (UnreadCounts)
	get(client, url, accessToken, &unreadCounts)
	return unreadCounts
}
