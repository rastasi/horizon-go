package entities

type Configuration struct {
	VerifyToken     string
	PageAccessToken string
	Incomings       []Incoming
	Outgoings       []Outgoing
	Endpoints       []Endpoint
}
