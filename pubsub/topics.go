package pubsub

type topic int

const (
	OnCandyExploding topic = iota
	OnCandyStartExploding
	OnPlayerWalking
	IncreasePlayerPower
	IncreasePlayerSpeed
	IncreaseCandyLimit
	OnDropCandy
)

type OnDropCandyPayload struct {
	X          int
	Y          int
	Width      int
	Height     int
	PowerLevel int
}

type OnPlayerWalkingPayload struct {
	X      int
	Y      int
	Width  int
	Height int
}
