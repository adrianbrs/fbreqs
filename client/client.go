package client

// Client contém os dados do Client
type Client struct {
	Request *Request
}

// New retorna uma nova instância do Client
func New() *Client {
	return &Client{
		Request: GetRequestConfig(),
	}
}
