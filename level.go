package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const maxLevelObjects int = 50

type level struct {
	boxes    *[maxLevelObjects]*box
	nOfBoxes int
}

type box struct {
	collider boxCollider
	graphic  *ebiten.Image
	opts     *ebiten.DrawImageOptions
}

func (l *level) Instantiate( /* probably a json file*/ ) {
	l.nOfBoxes = 0
	l.boxes = new([maxLevelObjects]*box)
	ground := newBox(newV2(0, screenHeight-10), newV2(screenWidth, screenHeight+90))
	l.addBox(ground)
}

func (l *level) addBox(b *box) {
	l.boxes[l.nOfBoxes] = b
	l.nOfBoxes++
}

func newBox(min, max vector2) *box {
	if min.x > max.x || min.y > max.y {
		fmt.Printf("Invalid box: (%f,%f)(%f,%f)", min.x, min.y, max.x, max.y)
		return nil
	}
	tmp := new(box)
	tmp.collider.min = min
	tmp.collider.max = max
	tmp.collider.mid = newV2((min.x+max.y)/2, (min.y+max.y)/2)
	tmp.graphic, _ = ebiten.NewImage(int(max.x-min.x), int(max.y-min.y), ebiten.FilterNearest)
	tmp.graphic.Fill(color.White)
	tmp.opts = &ebiten.DrawImageOptions{}
	tmp.opts.GeoM.Translate(min.x, min.y)
	return tmp
}
