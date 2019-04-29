package abn

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var guid string

// SetGUID obtain an API key from https://abr.business.gov.au/Tools/WebServicesAgreement
func SetGUID(key string) {
	guid = key
}

// LookUpABN Look up an Australian business number
func LookUpABN(abn string) map[string]interface{} {
	response, err := http.Get("https://abr.business.gov.au/json/AbnDetails.aspx?abn=" + abn + "&guid=" + guid)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return decode(body)
}

func decode(body []byte) map[string]interface{} {
	var data map[string]interface{}
	len := len(body) - 1
	err := json.Unmarshal(body[9:len], &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
