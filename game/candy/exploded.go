package candy

var _ state = (*explodedState)(nil)

type explodedState struct {
	sharedState
}

func (e explodedState) exploded() bool {
	return true
}
