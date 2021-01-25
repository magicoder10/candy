package game

import (
	"fmt"

	"candy/game/gameitem"
	"candy/game/square"
	"candy/graphics"

	"golang.org/x/image/font/basicfont"
)

const boxCount = 10
const reservedBoxesEnd = 2
const boxWidth = 70
const gameItemPadding = boxWidth - square.Width
const usableItemBoxX = 244
const itemCountXOffset = 33
const itemCountYOffset = 10

var backpackImageBound = graphics.Bound{
	X:      0,
	Y:      384,
	Width:  900,
	Height: 94,
}

var countImageBound = graphics.Bound{
	X:      701,
	Y:      0,
	Width:  28,
	Height: 18,
}

type box struct {
	x               int
	y               int
	gameItemType        gameitem.Type
	text            graphics.Text
	alwaysShowCount bool
}

func (b *box) draw(batch graphics.Batch, count int) {
	batch.DrawSprite(b.x+5, b.y+5, 0, b.gameItem.GetBound(), 1)

	if count > 0 || b.alwaysShowCount {
		batch.DrawSprite(b.x+34, b.y+8, 0, countImageBound, 1)

		fmt.Fprintf(b.text, "%d", count)
		b.text.Draw()
	}
}

type BackPack struct {
	screenX int
	screenY int
	boxes   []*box
	items   map[gameitem.Type]int
}

func (b *BackPack) AddItem(gameItem gameitem.GameItem) {
	if gameItem.GetType() == gameitem.NoneType {
		return
	}
	if _, ok := b.items[gameItem.GetType()]; !ok {
		// Find the first empty box
		index := 0
		for index < len(b.boxes) && b.boxes[index].gameItem != gameitem.NoneType {
			index++
		}
		// No empty box
		if index == len(b.boxes) {
			return
		}
		b.boxes[index].gameItem = gameItem.GetType()
	}
	b.items[gameItem.GetType()]++
	if gameItem.GetType().CanAutoUse() {
		gameItem.Use()
	}
}

func (b *BackPack) TakeItem(boxIndex int) gameitem.Type {
	box := b.boxes[boxIndex]
	if box.gameItemType != gameitem.NoneType {
		b.items[box.gameItem]--

		if boxIndex > reservedBoxesEnd && b.items[box.gameItem] == 0 {
			b.boxes[boxIndex].gameItemType = gameitem.NoneType
			delete(b.items, box.gameItem)
		}
	}
	return box.gameItem
}

func (b BackPack) Draw(batch graphics.Batch) {
	batch.DrawSprite(b.screenX, b.screenY, 0, backpackImageBound, 1)
	for i := range b.boxes {
		b.drawBox(batch, i)
	}
}

func (b BackPack) drawBox(batch graphics.Batch, index int) {
	b.boxes[index].draw(batch, b.items[b.boxes[index].gameItem])
}

func NewBackPack(g graphics.Graphics, screenX int, screenY int) BackPack {
	boxes := make([]*box, boxCount)
	items := make(map[gameitem.Type]int)

	x := screenX + gameItemPadding
	y := screenY + gameItemPadding

	boxes[0] = newBox(g, x, y, true)
	boxes[0].gameItem = gameitem.PowerType
	boxes[1] = newBox(g, x+boxWidth, y, true)
	boxes[1].gameItem = gameitem.SpeedType
	boxes[2] = newBox(g, x+boxWidth*2, y, true)
	boxes[2].gameItem = gameitem.CandyType
	for i := 0; i < boxCount-3; i++ {
		boxes[i+3] = newBox(g, screenX+usableItemBoxX+i*boxWidth, y, false)
	}

	items[gameitem.PowerType] = 0
	items[gameitem.SpeedType] = 0
	items[gameitem.CandyType] = 0
	return BackPack{
		screenX: screenX,
		screenY: screenY,
		boxes:   boxes,
		items:   items,
	}
}

func newBox(g graphics.Graphics, x int, y int, alwaysShowCount bool) *box {
	return &box{
		x:               x,
		y:               y,
		gameItem:        gameitem.NoneType,
		text:            g.NewText(basicfont.Face7x13, x+itemCountXOffset, y+itemCountYOffset, 28, 18, 1.2, graphics.AlignCenter),
		alwaysShowCount: alwaysShowCount,
	}
}
