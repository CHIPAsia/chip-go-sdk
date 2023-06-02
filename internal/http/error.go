package http

type Error struct {
	Type    string `json:"type"`
	ID      int    `json:"id,omitempty"`
	Message string `json:"message"`
}

const (
	NotFound   = 404
	OK         = 200
	BadRequest = 400
)
