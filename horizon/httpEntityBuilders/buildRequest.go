package httpEntityBuilders

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"bitbucket.org/aiventureteam/horizon-go/horizon/httpEntities"
)

func getBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	return body
}

func BuildRequest(r *http.Request) httpEntities.Request {
	body := getBody(r)
	var request httpEntities.Request
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Fatal(err)
	}
	return request
}
