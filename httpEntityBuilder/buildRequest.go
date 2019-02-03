package httpEntityBuilder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rastasi/horizon-go/httpEntity"
)

func getBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	return body
}

func BuildRequest(r *http.Request) httpEntity.Request {
	body := getBody(r)
	var request httpEntity.Request
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Fatal(err)
	}
	return request
}
