package gamestate

import (
	"math/rand"
)

const squareWidth = 60

const defaultRows = 12
const defaultCols = 15

type squareType int

const (
	EmptySquare squareType = iota
	GreenTile
	YellowTile
)

var squareTypes = []squareType{
	EmptySquare,
	GreenTile,
	YellowTile,
}

type GameItemType int

const (
	ItemTypeNone GameItemType = iota
	ItemTypeSpeed
	ItemTypePower
	ItemTypeCandy
	ItemTypeFirstAidKit
)

var gameItemTypes = []GameItemType{
	ItemTypeNone,
	ItemTypeSpeed,
	ItemTypePower,
	ItemTypeCandy,
	ItemTypeFirstAidKit,
}

type Square struct {
	SquareType   squareType
	GameItemType GameItemType
}

type Map struct {
	Squares [][]Square
}

func newMap() Map {
	squares := make([][]Square, 0)

	for row := 0; row < defaultRows; row++ {
		gridRow := make([]Square, 0)

		for col := 0; col < defaultCols; col++ {
			gridRow = append(gridRow, randomSquare())
		}
		squares = append(squares, gridRow)
	}

	return Map{
		Squares: squares,
	}
}

func randomSquare() Square {
	index := rand.Intn(len(squareTypes) * 2)
	if index >= len(squareTypes) || squareTypes[index] == EmptySquare {
		return Square{SquareType: EmptySquare}
	}
	sqType := squareTypes[index]

	index = rand.Intn(len(gameItemTypes))
	return Square{SquareType: sqType, GameItemType: gameItemTypes[index]}
}
