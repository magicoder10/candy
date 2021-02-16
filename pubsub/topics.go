package pubsub

type topic int

const (
	OnCandyExploding topic = iota
	OnPlayerWalking
	IncrementPlayerPower
	IncreasePlayerSpeed
	OnDropCandy
)

type OnDropCandyPayload struct {
	X      int
	Y      int
	Width  int
	Height int
}

type OnPlayerWalkingPayload struct {
	X      int
	Y      int
	Width  int
	Height int
}
