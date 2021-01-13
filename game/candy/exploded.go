package candy

var _ State = (*explodedState)(nil)

type explodedState struct {
	shared
}

func (e explodedState) Exploded() bool {
	return true
}
