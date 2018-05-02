package main

import (
	"fmt"
	//"image/color"
	"github.com/hajimehoshi/ebiten"
        "math"
	//"image"
	// "image/color"
	"encoding/json"
	"io/ioutil"
)

const maxLevelObjects int = 200

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
	MaxSortedShapes [maxLevelObjects]shape `json:"MaxSortedShapes"`
	MinSortedShapes [maxLevelObjects]shape `json:"MinSortedShapes"`
	NumOfShapes     int
	ShapeTypes		string
}

type box struct {
	Collider boxCollider
	Graphic  *ebiten.Image
	Opts     *ebiten.DrawImageOptions
}

type triangle struct {
	Collider triangleCollider
	Graphic  *ebiten.Image
	Opts     *ebiten.DrawImageOptions
}

type uninteractableImage struct {
	Graphic *ebiten.Image
	Opts    *ebiten.DrawImageOptions
}

func (img *uninteractableImage)draw(screen *ebiten.Image){
	screen.DrawImage(img.Graphic, img.Opts)
}

func (t triangle)getMin() vector2{
    return t.Collider.Min;
}

func (t triangle)getMax() vector2{
    return t.Collider.Max;
}

func (b box)getMin() vector2{
    return b.Collider.Min;
}

func (b box)getMax() vector2{
    return b.Collider.Max;
}

func (t triangle)getGraphic() *ebiten.Image{
    return t.Graphic
}

func (b box)getGraphic() *ebiten.Image{
    
    return b.Graphic
}

func (t triangle)getOpts() *ebiten.DrawImageOptions{
    return t.Opts
}

func (b box)getOpts() *ebiten.DrawImageOptions{
    return b.Opts
}

func (l *level) Instantiate(filePath string) {
	l.ShapeTypes=""
	rawData, err := ioutil.ReadFile("levels/"+filePath)
	if err != nil {
		fmt.Println("failed to load data from json")
		panic("rip")
	}
	jsonData :=level{[maxLevelObjects]shape{},
						[maxLevelObjects]shape{},
						0,""}
	if err = json.Unmarshal(rawData, jsonData); err==nil{
		fmt.Println("failed to unmarshal")
		panic("rip")
	}


	//shapes:=jsonData.MaxSortedShapes
	fmt.Println(jsonData)

	/*types:= []rune(jsonData.ShapeTypes)
	for i:=0;i<jsonData.NumOfShapes;i++{
		min:=jsonData.MaxSortedShapes[i].(box).Collider.Min
		max:=jsonData.MaxSortedShapes[i].getMax()
		graphic:=jsonData.MaxSortedShapes[i].getGraphic()
		opts:=jsonData.MaxSortedShapes[i].getOpts()
		switch types[i] {
		case 'B':
			tmp:=newBox(min,max)
			tmp.Opts=opts
			tmp.Graphic=graphic
			tmp.Graphic.Fill(color.White)
			l.addBox(tmp)
		}

	}*/


	//empty level

	l.NumOfShapes = 0
	l.MaxSortedShapes = [maxLevelObjects]shape{}
	l.MinSortedShapes =	[maxLevelObjects]shape{}
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
	if l.MaxSortedShapes[0] != nil {
		//maxSorted
		i:=0
		var originalArray []shape
		for _,s := range l.MaxSortedShapes{
			originalArray=append(originalArray,s)
		}
		for i<l.NumOfShapes && b.getMax().X>=l.MaxSortedShapes[i].getMax().X {
			i++
		}
		l.MaxSortedShapes[i]=b
		for i<l.NumOfShapes {
			l.MaxSortedShapes[i+1]=originalArray[i]
			i++
		}

		i=0
		var originalArray2 []shape
		for _,s := range l.MinSortedShapes{
			originalArray2=append(originalArray2,s)
		}
		for i<l.NumOfShapes && b.getMin().X>=l.MinSortedShapes[i].getMin().X {
			i++
		}
		l.MinSortedShapes[i]=b
		for i<l.NumOfShapes {
			l.MinSortedShapes[i+1]=originalArray2[i]
			i++
		}



	} else {
		var interfacedBox shape= b
		l.MaxSortedShapes[0] = interfacedBox
		l.MinSortedShapes[0] = interfacedBox
	}
	l.NumOfShapes++

	l.ShapeTypes=""
	for j:=0;j<l.NumOfShapes;j++{
		switch (l.MaxSortedShapes[j]).(type) {
			case *box: l.ShapeTypes+="B"
			case *triangle: l.ShapeTypes+="T"
		}
	}
}

func (l *level) addTriangle(t *triangle){
	if l.MaxSortedShapes[0] != nil {
		//maxSorted
		i:=0
		var originalArray []shape
		for _,s := range l.MaxSortedShapes{
			originalArray=append(originalArray,s)
		}
		for i<l.NumOfShapes && t.getMax().X>=l.MaxSortedShapes[i].getMax().X {
			i++
		}
		l.MaxSortedShapes[i]=t
		for i<l.NumOfShapes {
			l.MaxSortedShapes[i+1]=originalArray[i]
			i++
		}

		i=0
		var originalArray2 []shape
		for _,s := range l.MinSortedShapes{
			originalArray2=append(originalArray2,s)
		}
		for i<l.NumOfShapes && t.getMin().X>=l.MinSortedShapes[i].getMin().X {
			i++
		}
		l.MinSortedShapes[i]=t
		for i<l.NumOfShapes {
			l.MinSortedShapes[i+1]=originalArray2[i]
			i++
		}



	} else {
		var interfacedBox shape= t
		l.MaxSortedShapes[0] = interfacedBox
		l.MinSortedShapes[0] = interfacedBox
	}
	l.NumOfShapes++

	l.ShapeTypes=""
	for j:=0;j<l.NumOfShapes;j++{
		switch (l.MaxSortedShapes[j]).(type) {
		case *box: l.ShapeTypes+="B"
		case *triangle: l.ShapeTypes+="T"
		}
	}
}

func newBox(min, max vector2) *box {
	if min.X > max.X || min.Y > max.Y {
		fmt.Printf("Invalid box: (%f,%f)(%f,%f)", min.X, min.Y, max.X, max.Y)
		return nil
	}
	tmp := new(box)
	tmp.Collider.Min = min
	tmp.Collider.Max = max
	tmp.Collider.Mid = newV2((min.X+max.Y)/2, (min.Y+max.Y)/2)
	tmp.Graphic, _ = ebiten.NewImageFromImage(boxGraphic, ebiten.FilterNearest)
	//tmp.Graphic.Fill(color.White)
	tmp.Opts = &ebiten.DrawImageOptions{}
	tmp.Opts.GeoM.Translate(min.X, min.Y)
	return tmp
}

func newTriangle(min, max vector2, side string) *triangle {
	if min.X > max.X || min.Y > max.Y {
		fmt.Printf("Invalid triangle: (%f,%f)(%f,%f)", min.X, min.Y, max.X, max.Y)
		return nil
	}
	tmp := new(triangle)
	tmp.Collider.Min = min
	tmp.Collider.Max = max
	tmp.Collider.Mid = newV2((min.X+max.Y)/2, (min.Y+max.Y)/2)
 	tmp.Graphic, _ = ebiten.NewImageFromImage(triangleGraphic, ebiten.FilterNearest)
	tmp.Opts = &ebiten.DrawImageOptions{}
	
        //tmp.Opts.GeoM.Scale(0.001*(max.X-min.X), 0.001*(max.Y-min.Y))
        
        switch side{
            case "top-left": tmp.Opts.GeoM.Rotate(-math.Pi/2)
            case "top-right": 
            case "bottom-left": tmp.Opts.GeoM.Rotate(math.Pi)
            case "bottom-right": tmp.Opts.GeoM.Rotate(math.Pi/2)
        }

        tmp.Opts.GeoM.Translate(min.X, min.Y)

	return tmp
}