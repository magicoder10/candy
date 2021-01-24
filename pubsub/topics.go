package pubsub

type topic int

const (
	OnCandyExploding topic = iota
	OnDropCandy
)

type OnDropCandyPayload struct {
	X int
	Y int
	Width int
	Height int
}
