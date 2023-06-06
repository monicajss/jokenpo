package entities

const (
	ROCK    string = "rock"
	PAPER          = "paper"
	SISSORS        = "sissors"
)

// Jokenpo represents the jokenpo moves
type Jokenpo struct {
	Move string `json:"move" move:"move"`
}
