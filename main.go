package main

import (
	"bitbucket.org/aiventureteam/horizon-go/entity"
	"bitbucket.org/aiventureteam/horizon-go/horizon"
)

const (
	verifyToken     = "rubyforever"
	pageAccessToken = "EAACChyvB7ZCMBAPXK1i9VDFJyPdEEArsMuAuGNFsuCvfe9Sp6LCX21Rkmii580x25GKUVrEYQkx7BDPdOgnibTt9I9jIx4uIgduUbaafmkmiBGn01KxhIU8yKaXjpBQfRZBDhmy2oiP1wsqh0CKgoDZCj6LEGI5mVD3BcZBtBAZDZD"
)

func main() {
	incomings := []entity.Incoming{}
	outgoings := []entity.Outgoing{}
	endpoints := []entity.Endpoint{}
	configuration := entity.Configuration{
		VerifyToken:     verifyToken,
		PageAccessToken: pageAccessToken,
		Incomings:       incomings,
		Outgoings:       outgoings,
		Endpoints:       endpoints,
	}
	horizon.Start(configuration)
}
