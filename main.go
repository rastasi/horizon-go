package main

import (
	"bitbucket.org/aiventureteam/horizon-go/horizon"
	"bitbucket.org/aiventureteam/horizon-go/horizon/entities"
)

func main() {
	incomings := []entities.Incoming{}
	outgoings := []entities.Outgoing{}
	endpoints := []entities.Endpoint{}
	verifyToken := "verifyToken"
	pageAccessToken := "pageAccessToken"
	configuration := entities.Configuration{
		VerifyToken:     verifyToken,
		PageAccessToken: pageAccessToken,
		Incomings:       incomings,
		Outgoings:       outgoings,
		Endpoints:       endpoints,
	}
	horizon.Start(configuration)
}
