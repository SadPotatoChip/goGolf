package main

import (
	"fmt"
	//"image/color"
	"github.com/hajimehoshi/ebiten"
        "math"
	//"image"
	"image/color"
)

const maxLevelObjects int = 2000

type shape interface{
    
    getMin() vector2
    getMax() vector2
  	getGraphic() *ebiten.Image
    getOpts() *ebiten.DrawImageOptions
}

///Keeping arrays for clarity, each represents the boxes in the level sorted by their
///min-max x coordinate, a different array is used for collision detection depending on
///the balls horizontal direction for element elimination
type level struct {
	maxSortedShapes *[maxLevelObjects]*shape
	minSortedShapes *[maxLevelObjects]*shape
	nOfShapes       int
}

type box struct {
	collider boxCollider
	graphic  *ebiten.Image
	opts     *ebiten.DrawImageOptions
}

type triangle struct {
	collider triangleCollider
	graphic  *ebiten.Image
	opts     *ebiten.DrawImageOptions
}

type uninteractableImage struct {
	graphic *ebiten.Image
	opts    *ebiten.DrawImageOptions
}

func (img *uninteractableImage)draw(screen *ebiten.Image){
	screen.DrawImage(img.graphic, img.opts)
}


func (t triangle)getMin() vector2{
    return t.collider.min;
}

func (t triangle)getMax() vector2{
    return t.collider.max;
}

func (b box)getMin() vector2{
    return b.collider.min;
}

func (b box)getMax() vector2{
    return b.collider.max;
}

func (t triangle)getGraphic() *ebiten.Image{
    return t.graphic
}

func (b box)getGraphic() *ebiten.Image{
    
    return b.graphic
}

func (t triangle)getOpts() *ebiten.DrawImageOptions{
    return t.opts
}

func (b box)getOpts() *ebiten.DrawImageOptions{
    return b.opts
}

func (l *level) Instantiate( /* probably a json file*/ ) {
	l.nOfShapes = 0
	l.maxSortedShapes = new([maxLevelObjects]*shape)
	l.minSortedShapes = new([maxLevelObjects]*shape)
	ground := newBox(newV2(0, screenHeight-10), newV2(screenWidth, screenHeight+90))
	ceiling:= newBox(newV2(0,-80), newV2(screenWidth-1, 10))
	wallLeft:= newBox(newV2(0,-80), newV2(10, screenHeight))
	wallRight:= newBox(newV2(screenWidth-10,0), newV2(screenWidth+90, screenHeight))
	l.addBox(ground)
	l.addBox(ceiling)
	l.addBox(wallLeft)
	l.addBox(wallRight)
}


func (l *level) addBox(b *box) {
	//manually realocate arrays so that it is sorted
	if l.maxSortedShapes[0] != nil {
		//maxSorted
		i:=1
		originalArray:=*l.maxSortedShapes
		tmp := *l.maxSortedShapes[0]
		for ;i < l.nOfShapes && b.collider.max.x > tmp.getMax().x; i++ {
			tmp = *l.maxSortedShapes[i]
		}

		if i>0 && i<l.nOfShapes{
			i--
		}

		var interfacedBox shape= b
		l.maxSortedShapes[i] = &interfacedBox
		for i < l.nOfShapes {
			l.maxSortedShapes[i+1] = originalArray[i]
			i++
		}

		//minSorted
		i=1
		originalArray=*l.minSortedShapes
		tmp = *l.minSortedShapes[0]
		for ;i < l.nOfShapes && b.collider.min.x > tmp.getMin().x; i++ {
			tmp = *l.minSortedShapes[i]
		}

		if i>0 && i<l.nOfShapes{
			i--
		}

		interfacedBox= b
		l.minSortedShapes[i] = &interfacedBox
		for i < l.nOfShapes {
			l.minSortedShapes[i+1] = originalArray[i]
			i++
		}
	} else {
		var interfacedBox shape= b
		l.maxSortedShapes[0] = &interfacedBox
		l.minSortedShapes[0] = &interfacedBox
	}
	l.nOfShapes++
	//debug
	for j:=0;j<l.nOfShapes;j++{
		switch (*l.maxSortedShapes[j]).(type) {
			case *box: fmt.Printf("B")
			case *triangle: fmt.Printf("T")
		}
	}
	fmt.Printf("\n")

}

func (l *level) addTriangle(t *triangle){
	if l.maxSortedShapes[0] != nil {
		//maxSorted
		i:=1
		originalArray:=*l.maxSortedShapes
		tmp := *l.maxSortedShapes[0]
			for ;i < l.nOfShapes && t.collider.max.x > tmp.getMax().x; i++ {
			tmp = *l.maxSortedShapes[i]
		}
	
		if i>0 && i<l.nOfShapes{
			i--
		}
	
		var interfacedBox shape= t
		l.maxSortedShapes[i] = &interfacedBox
		for i < l.nOfShapes {
			l.maxSortedShapes[i+1] = originalArray[i]
			i++
		}
	
		//minSorted
		i=1
		originalArray=*l.minSortedShapes
		tmp = *l.minSortedShapes[0]
		for ;i < l.nOfShapes && t.collider.min.x > tmp.getMin().x; i++ {
			tmp = *l.minSortedShapes[i]
		}
	
		if i>0 && i<l.nOfShapes{
			i--
		}

		interfacedBox= t
		l.minSortedShapes[i] = &interfacedBox
		for i < l.nOfShapes {
			l.minSortedShapes[i+1] = originalArray[i]
			i++
		}
	} else {
		var interfacedBox shape = t
		l.maxSortedShapes[0] = &interfacedBox
		l.minSortedShapes[0] = &interfacedBox
	}
	
	l.nOfShapes++
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

func newTriangle(min, max vector2, side string) *triangle {
	if min.x > max.x || min.y > max.y {
		fmt.Printf("Invalid triangle: (%f,%f)(%f,%f)", min.x, min.y, max.x, max.y)
		return nil
	}
	tmp := new(triangle)
	tmp.collider.min = min
	tmp.collider.max = max
	tmp.collider.mid = newV2((min.x+max.y)/2, (min.y+max.y)/2)
 	tmp.graphic, _ = ebiten.NewImageFromImage(triangle_graphic, ebiten.FilterNearest)
	tmp.opts = &ebiten.DrawImageOptions{}
	
        tmp.opts.GeoM.Scale(0.001*(max.x-min.x), 0.001*(max.y-min.y))
        
        switch side{
            case "top-left": tmp.opts.GeoM.Rotate(-math.Pi/2)
            case "top-right": 
            case "bottom-left": tmp.opts.GeoM.Rotate(math.Pi)
            case "bottom-right": tmp.opts.GeoM.Rotate(math.Pi/2)
        }

        tmp.opts.GeoM.Translate(min.x, min.y)

	return tmp
}