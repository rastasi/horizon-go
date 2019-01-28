package horizon

import (
	"log"

	"fmt"
	"net/http"

	"bitbucket.org/aiventureteam/horizon-go/entity"
	"bitbucket.org/aiventureteam/horizon-go/webhookEndpoint"
	"github.com/gorilla/mux"
)

func Start(configuration entity.Configuration) {
	configuration.MessagesURL = "https://graph.facebook.com/v2.6/me/messages"
	router := mux.NewRouter()
	router.HandleFunc("/webhook", webhookEndpoint.GetWebhook(configuration)).Methods("GET")
	router.HandleFunc("/webhook", webhookEndpoint.PostWebhook(configuration)).Methods("POST")
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", router))
}
