package client

import "fmt"

// Error erro do sender
type Error struct {
	ID          string
	Description string
}

func (err *Error) Error() string {
	return fmt.Sprintf("[Error] %s: %s", err.ID, err.Description)
}

var (
	EmptyRequestURL = &Error{
		ID:          "EmptyRequestURL",
		Description: "The request URL is empty",
	}
	InvalidProxy = &Error{
		ID:          "InvalidProxy",
		Description: "Proxy inválido",
	}
	InvalidRequestAmount = &Error{
		ID:          "InvalidRequestAmount",
		Description: "Amount of requests should be greater than zero",
	}
	EmptyProxyList = &Error{
		ID:          "EmptyProxyList",
		Description: "Could not find any proxy in proxy list",
	}
)
