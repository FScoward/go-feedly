package main

type Config struct {
	Feedly FeedlyConfig
}

type FeedlyConfig struct {
	ClientSecret string
	AccessToken string
}
