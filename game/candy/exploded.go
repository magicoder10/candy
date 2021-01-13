package candy

var _ state = (*explodedState)(nil)

type explodedState struct {
	shared
}

func (e explodedState) Exploded() bool {
	return true
}
