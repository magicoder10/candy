package pubsub

type topic int

const (
	OnCandyExploding topic = iota
	OnPlayerWalking
	IncrementPlayerPower
)

type OnPlayerWalkingPayload struct {
	X      int
	Y      int
	Width  int
	Height int
}
