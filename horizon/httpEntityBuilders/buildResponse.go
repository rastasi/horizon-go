package httpEntityBuilders

import (
	"encoding/json"
	"log"

	"bitbucket.org/aiventureteam/horizon-go/horizon/httpEntities"
)

func getJSONResponse(response httpEntities.Response) []byte {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	return jsonResponse
}

func BuildResponse(senderID string, text string) []byte {
	response := httpEntities.Response{Recipient: httpEntities.Participant{Id: senderID}, Message: httpEntities.Message{Text: text}}
	jsonResponse := getJSONResponse(response)
	return jsonResponse
}
