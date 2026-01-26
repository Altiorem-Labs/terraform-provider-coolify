package client

import "net/http"

type CoolifyClient struct {
	Endpoint string
	Token    string
	Client   *http.Client
}
