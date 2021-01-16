package player

import (
	"fmt"

	"candy/graphics"
	"candy/screen/game/gameitem"
	"candy/screen/game/square"

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
	gameItem        gameitem.GameItem
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
	items   map[gameitem.GameItem]int
}

func (b *BackPack) AddItem(gameItem gameitem.GameItem) {
	if _, ok := b.items[gameItem]; !ok {
		// Find the first empty box
		index := 0
		for index < len(b.boxes) && b.boxes[index].gameItem != gameitem.None {
			index++
		}
		// No empty box
		if index == len(b.boxes) {
			return
		}
		b.boxes[index].gameItem = gameItem
	}
	b.items[gameItem]++
}

func (b *BackPack) TakeItem(boxIndex int) gameitem.GameItem {
	box := b.boxes[boxIndex]
	if box.gameItem != gameitem.None {
		b.items[box.gameItem]--

		if boxIndex > reservedBoxesEnd && b.items[box.gameItem] == 0 {
			b.boxes[boxIndex].gameItem = gameitem.None
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
	items := make(map[gameitem.GameItem]int)

	x := screenX + gameItemPadding
	y := screenY + gameItemPadding

	boxes[0] = newBox(g, x, y, true)
	boxes[0].gameItem = gameitem.Power
	boxes[1] = newBox(g, x+boxWidth, y, true)
	boxes[1].gameItem = gameitem.Speed
	boxes[2] = newBox(g, x+boxWidth*2, y, true)
	boxes[2].gameItem = gameitem.Candy
	for i := 0; i < boxCount-3; i++ {
		boxes[i+3] = newBox(g, screenX+usableItemBoxX+i*boxWidth, y, false)
	}

	items[gameitem.Power] = 0
	items[gameitem.Speed] = 0
	items[gameitem.Candy] = 0
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
		gameItem:        gameitem.None,
		text:            g.NewText(basicfont.Face7x13, x+itemCountXOffset, y+itemCountYOffset, 28, 18, 1.2, graphics.AlignCenter),
		alwaysShowCount: alwaysShowCount,
	}
}
