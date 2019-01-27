package httpEntities

type Response struct {
	Recipient Participant `json:"recipient"`
	Message   Message     `json:"message"`
}
