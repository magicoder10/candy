package pubsub

type topic int

const (
	OnCandyExploding topic = iota
	OnPlayerWalking
	IncreasePlayerPower
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
