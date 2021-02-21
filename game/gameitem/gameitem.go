package gameitem

import (
	"candy/graphics"
	"candy/pubsub"
)

type Type int

const (
	NoneType Type = iota
	SpeedType
	PowerType
	CandyType
	FirstAidKitType
)

var Types = []Type{
	NoneType,
	SpeedType,
	PowerType,
	CandyType,
	FirstAidKitType,
}

func (t Type) GetBound() graphics.Bound {
	switch t {
	case SpeedType:
		return graphics.Bound{
			X:      761,
			Y:      204,
			Width:  60,
			Height: 60,
		}
	case PowerType:
		return graphics.Bound{
			X:      761,
			Y:      264,
			Width:  60,
			Height: 60,
		}
	case CandyType:
		return graphics.Bound{
			X:      761,
			Y:      144,
			Width:  60,
			Height: 60,
		}
	case FirstAidKitType:
		return graphics.Bound{
			X:      761,
			Y:      84,
			Width:  60,
			Height: 60,
		}
	default:
		return graphics.Bound{}
	}
}

func (t Type) CanAutoUse() bool {
	switch t {
	case PowerType, SpeedType, CandyType:
		return true
	}
	return false
}

type GameItem interface {
	GetType() Type
	Use()
}

var _ GameItem = (*Power)(nil)

type Power struct {
	pubSub *pubsub.PubSub
}

func (p Power) Use() {
	p.pubSub.Publish(pubsub.IncreasePlayerPower, 1)
}

func (p Power) GetType() Type {
	return PowerType
}

var _ GameItem = (*Speed)(nil)

type Speed struct {
}

func (p Speed) Use() {
}

func (p Speed) GetType() Type {
	return SpeedType
}

var _ GameItem = (*Candy)(nil)

type Candy struct {
}

func (p Candy) Use() {
}

func (p Candy) GetType() Type {
	return CandyType
}

var _ GameItem = (*FirstAidKit)(nil)

type FirstAidKit struct {
}

func (p FirstAidKit) Use() {
}

func (p FirstAidKit) GetType() Type {
	return FirstAidKitType
}

type None struct {
}

var _ GameItem = (*None)(nil)

func (n None) GetType() Type {
	return NoneType
}

func (n None) Use() {
	return
}

func WithPubSub(itemType Type, pubSub *pubsub.PubSub) GameItem {
	switch itemType {
	case PowerType:
		return Power{
			pubSub: pubSub,
		}
	case SpeedType:
		return Speed{}
	case CandyType:
		return Candy{}
	case FirstAidKitType:
		return FirstAidKit{}
	}
	return None{}
}
