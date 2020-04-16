package sender

import (
	"fmt"
	"net/url"

	"github.com/Pallinder/go-randomdata"
)

// GetData returns data to be sent in request
func GetData() map[string][]string {
	acesso := ""
	if randomdata.Number(0, 100) > 1 {
		acesso = "FISICA"
	} else {
		acesso = "pj"
	}

	stName := randomdata.FirstName(randomdata.Number(0, 1))
	ndName := randomdata.LastName()
	username := stName + ndName[:randomdata.Number(1, len(ndName))]

	formData := url.Values{
		"user":   {username},
		"pass":   {fmt.Sprintf("%d", randomdata.Number(10000000, 99999999))},
		"acesso": {acesso},
		"sender": {"account"},
		"assina": {fmt.Sprintf("%d", randomdata.Number(100000, 999999))},
	}
	return formData
}
