package webhookEndpoints

import (
	"bytes"
	"net/http"

	"bitbucket.org/aiventureteam/horizon-go/horizon/entities"
	"bitbucket.org/aiventureteam/horizon-go/horizon/httpEntityBuilders"
)

func PostWebhook(configuration entities.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := httpEntityBuilders.BuildRequest(r)

		if request.Object == "page" {
			event := request.Entry[0].Messaging[0]
			if event.Message.Text != "" {
				senderID := event.Sender.Id
				text := getResponseText(event.Message.Text)
				response := httpEntityBuilders.BuildResponse(senderID, text)
				sendResponse(response)
			}
		}
		w.Write([]byte("EVENT_RECEIVED"))
	}
}

func sendResponse(response []byte) {
	request, err := http.NewRequest("POST", messagesURL, bytes.NewBuffer(response))
	request.Header.Set("Content-Type", "application/json")

	q := request.URL.Query()
	q.Add("access_token", pageAccessToken)
	request.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func getResponseText(input string) string {
	if input == "hello" {
		return "hello!"
	}
	if input == "help" {
		return "Can I help you?"
	}
	if input == "contact" {
		return "www.aiventure.me"
	}
	return "Sorry, I don't undersand."
}
