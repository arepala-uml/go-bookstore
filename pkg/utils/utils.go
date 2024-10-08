package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"
)

// It helps the parse the body while creating book
func ParseBody(r *http.Request, x interface{}) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error reading request body: %v", err)
		return
	}

	err = json.Unmarshal([]byte(body), x)
	if err != nil {
		log.Errorf("Error unmarshalling JSON: %v", err)
		return
	}
}
