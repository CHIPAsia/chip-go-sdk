package chip

type RelatedObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type StatusHistory struct {
	Status        Status        `json:"status"`
	Timestamp     int           `json:"timestamp"`
	RelatedObject RelatedObject `json:"related_object"`
}
