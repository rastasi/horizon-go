package httpEntityBuilder

import (
	"encoding/json"
	"log"

	"github.com/rastasi/horizon-go/httpEntity"
)

func getJSONResponse(response httpEntity.Response) []byte {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	return jsonResponse
}

func BuildResponse(senderID string, text string) []byte {
	response := httpEntity.Response{Recipient: httpEntity.Participant{Id: senderID}, Message: httpEntity.Message{Text: text}}
	jsonResponse := getJSONResponse(response)
	return jsonResponse
}
