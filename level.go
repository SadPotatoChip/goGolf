package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const maxLevelObjects int = 2000


///Keeping arrays for clarity, each represents the boxes in the level sorted by their
///min-max x coordinate, a different array is used for collision detection depending on
///the balls horizontal direction for element elimination
type level struct {
	maxSortedBoxes *[maxLevelObjects]*box
	minSortedBoxes *[maxLevelObjects]*box
	nOfBoxes       int
}

type box struct {
	collider boxCollider
	graphic  *ebiten.Image
	opts     *ebiten.DrawImageOptions
}

func (l *level) Instantiate( /* probably a json file*/ ) {
	l.nOfBoxes = 0
	l.maxSortedBoxes = new([maxLevelObjects]*box)
	l.minSortedBoxes = new([maxLevelObjects]*box)
	ground := newBox(newV2(0, screenHeight-10), newV2(screenWidth, screenHeight+90))
	ceiling:= newBox(newV2(0,-80), newV2(screenWidth, 10))
	wallLeft:= newBox(newV2(0,-80), newV2(10, screenHeight))
	wallRight:= newBox(newV2(screenWidth-10,0), newV2(screenWidth+90, screenHeight))
	l.addBox(ground)
	l.addBox(ceiling)
	l.addBox(wallLeft)
	l.addBox(wallRight)
}

func (l *level) addBox(b *box) {
	//manually realocate arrays so that it is sorted
	if l.maxSortedBoxes[0] != nil {
		//maxSorted
		var i int
		tmp := *l.maxSortedBoxes
		for i < l.nOfBoxes && b.collider.max.x > l.maxSortedBoxes[i].collider.max.x {
			i++
		}
		l.maxSortedBoxes[i] = b
		for i < l.nOfBoxes {
			l.maxSortedBoxes[i+1] = tmp[i]
			i++
		}

		//minSorted
		i=0
		tmp = *l.minSortedBoxes
		for i < l.nOfBoxes && b.collider.min.x > l.minSortedBoxes[i].collider.min.x {
			i++
		}
		l.minSortedBoxes[i] = b
		for i < l.nOfBoxes {
			l.minSortedBoxes[i+1] = tmp[i]
			i++
		}
	} else {
		l.maxSortedBoxes[0] = b
		l.minSortedBoxes[0] = b
	}
	l.nOfBoxes++
	//debug
	/*for i:=0;i< l.nOfBoxes;i++{
		fmt.Printf("%f ",l.min[i].collider.min.x )
	}
	fmt.Printf("\n")*/

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
