package restclient

import (
	"net/http"
	"time"
)

type RestClient struct {
	client      *http.Client
	BearerToken string
}

func NewRestClient() (RestClient, error) {
	httpClient := http.Client{Timeout: time.Duration(20) * time.Second}
	client := RestClient{
		client:      &httpClient,
		BearerToken: bearer{}.token,
	}

	return client, nil
}
