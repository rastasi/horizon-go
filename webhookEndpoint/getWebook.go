package webhookEndpoint

import (
	"net/http"

	"github.com/rastasi/horizon-go/entity"
)

func GetWebhook(configuration entity.Configuration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mode := r.URL.Query()["hub.mode"][0]
		token := r.URL.Query()["hub.verify_token"][0]
		challenge := r.URL.Query()["hub.challenge"][0]

		if mode != "" && token != "" {
			if mode == "subscribe" && token == configuration.VerifyToken {
				w.Write([]byte(challenge))
			} else {
				w.Write([]byte("Forbidden"))
			}
		}
	}
}
