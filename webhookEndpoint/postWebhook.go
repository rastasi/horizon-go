package webhookEndpoint

import (
	"bytes"
	"net/http"

	"bitbucket.org/aiventureteam/horizon-go/entity"
	"bitbucket.org/aiventureteam/horizon-go/httpEntityBuilder"
)

func PostWebhook(configuration entity.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := httpEntityBuilder.BuildRequest(r)

		if request.Object == "page" {
			event := request.Entry[0].Messaging[0]
			if event.Message.Text != "" {
				senderID := event.Sender.Id
				text := getResponseText(event.Message.Text)
				response := httpEntityBuilder.BuildResponse(senderID, text)
				sendResponse(response, configuration)
			}
		}
		w.Write([]byte("EVENT_RECEIVED"))
	}
}

func sendResponse(response []byte, configuration entity.Configuration) {
	request, err := http.NewRequest("POST", configuration.MessagesURL, bytes.NewBuffer(response))
	request.Header.Set("Content-Type", "application/json")

	q := request.URL.Query()
	q.Add("access_token", configuration.PageAccessToken)
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
