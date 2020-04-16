package sender

import (
	"caixa-falso/client"
	"net/http"
)

// GetHTTPClient retorna um *http.Client
func GetHTTPClient(req *client.Request) (httpClient *http.Client, err error) {
	httpClient = http.DefaultClient

	// Pega o proxy disponivel
	proxy, err := req.Proxy.GetProxy()

	// Verifica se há um proxy disponível
	if err == nil {
		// Cria a conexão com proxy
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
	}

	return httpClient, nil
}
