package horizon

import (
	"log"

	"fmt"
	"net/http"

	"bitbucket.org/aiventureteam/horizon-go/horizon/entities"
	"bitbucket.org/aiventureteam/horizon-go/horizon/webhookEndpoints"
	"github.com/gorilla/mux"
)

const (
	verifyToken     = "rubyforever"
	pageAccessToken = "EAACChyvB7ZCMBAPXK1i9VDFJyPdEEArsMuAuGNFsuCvfe9Sp6LCX21Rkmii580x25GKUVrEYQkx7BDPdOgnibTt9I9jIx4uIgduUbaafmkmiBGn01KxhIU8yKaXjpBQfRZBDhmy2oiP1wsqh0CKgoDZCj6LEGI5mVD3BcZBtBAZDZD"
	messagesURL     = "https://graph.facebook.com/v2.6/me/messages"
)

func Start(configuration entities.Configuration) {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", webhookEndpoints.GetWebhook(configuration)).Methods("GET")
	router.HandleFunc("/webhook", webhookEndpoints.PostWebhook(configuration)).Methods("POST")
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", router))
}
