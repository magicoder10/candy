package pubsub

type topic int

const (
	OnCandyExploding topic = iota
	OnPlayerWalking
	IncreasePlayerPower
	IncreasePlayerSpeed
	OnDropCandy
)

type OnDropCandyPayload struct {
	X              int
	Y              int
	Width          int
	Height         int
	PowerLevel     int
	MaxCandyAmount int
}

type OnPlayerWalkingPayload struct {
	X      int
	Y      int
	Width  int
	Height int
}
