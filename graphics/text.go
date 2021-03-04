package graphics

import (
	"bufio"
	"io/ioutil"

	"golang.org/x/image/font"
)

type ebitenText struct {
	buf         *bufio.ReadWriter
	textContent string
	fontFace    font.Face
	graphics    *Ebiten
	x           int
	y           int
	width       int
	height      int
	alignment   alignment
}

func (t *ebitenText) Write(p []byte) (int, error) {
	return t.buf.Write(p)
}

func (t *ebitenText) Draw() {
	_ = t.buf.Flush()
	buf, _ := ioutil.ReadAll(t.buf)
	t.textContent = string(buf)
}

