package httpEntities

type Messaging struct {
	Sender    Participant `json:"sender"`
	Recipient Participant `json:"recipient"`
	Timestamp int         `json:"timestamp"`
	Message   Message     `json:"message"`
}

type Entry struct {
	Id        string      `json:"id"`
	Time      int         `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

type Request struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}
