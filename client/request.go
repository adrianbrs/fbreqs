package client

import (
	"github.com/spf13/viper"
)

// Request configuration values
type Request struct {
	URL    string
	Amount int
	Proxy  *Proxy
}

// GetRequestConfig retorna uma instância do Request
func GetRequestConfig() *Request {
	req := &Request{
		URL:    viper.GetString("request.url"),
		Amount: viper.GetInt("request.amount"),
		Proxy:  GetProxyConfig(),
	}

	// Validate request URL
	if req.URL == "" {
		panic(EmptyRequestURL)
	}

	// Validate request amount
	if req.Amount <= 0 {
		panic(InvalidRequestAmount)
	}

	return req
}
