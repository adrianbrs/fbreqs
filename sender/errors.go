package sender

import "fmt"

// Error erro do sender
type Error struct {
	ID          string
	Name        string
	Description string
}

func (err *Error) Error() string {
	return fmt.Sprintf("[%s] - %s: %s", err.ID, err.Name, err.Description)
}
