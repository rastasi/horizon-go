package entity

type Configuration struct {
	VerifyToken     string
	PageAccessToken string
	MessagesURL     string
	Incomings       []Incoming
	Outgoings       []Outgoing
	Endpoints       []Endpoint
}
