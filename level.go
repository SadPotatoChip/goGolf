package main

import (
	"fmt"
	//"image/color"
	"github.com/hajimehoshi/ebiten"
        "math"
	//"image"
	// "image/color"
	//"encoding/json"
	//"io/ioutil"
	"io/ioutil"
	"encoding/json"
	"image/color"
)

const maxLevelObjects int = 200

///Keeping arrays for clarity, each represents the boxes in the level sorted by their
///min-max x coordinate, a different array is used for collision detection depending on
///the balls horizontal direction for element elimination
type level struct {
	hole	*hole
	MaxSortedShapes [maxLevelObjects]shape `json:"MaxSortedShapes"`
	MinSortedShapes [maxLevelObjects]shape `json:"MinSortedShapes"`
	NumOfShapes     int
	ShapeTypes		string
}

func (l *level) Instantiate(filePath string) {
	l.ShapeTypes=""


	//fmt.Println(jsonData)


	//clear level
	l.NumOfShapes = 0
	l.MaxSortedShapes = [maxLevelObjects]shape{}
	l.MinSortedShapes =	[maxLevelObjects]shape{}
	ground := newBox(newV2(0, screenHeight), newV2(screenWidth, screenHeight+90))
	ceiling:= newBox(newV2(0,-80), newV2(screenWidth-1, 0))
	wallLeft:= newBox(newV2(0,-80), newV2(0, screenHeight))
	wallRight:= newBox(newV2(screenWidth,0), newV2(screenWidth+90, screenHeight))
	l.addBox(ground)
	l.addBox(ceiling)
	l.addBox(wallLeft)
	l.addBox(wallRight)
	//l.hole=newHole(vector2{340,260},vector2{600,600})


}

func parseJsonFromPath(filePath string) level{
	rawData, err := ioutil.ReadFile("levels/"+filePath)
	if err != nil {
		fmt.Println("failed to load data from json")
		panic("rip")
	}
	jsonData :=level{nil,[maxLevelObjects]shape{},
	[maxLevelObjects]shape{},
	0,""}
	if err = json.Unmarshal(rawData, jsonData); err==nil{
		fmt.Println("failed to unmarshal")
		panic("rip")
	}
	return jsonData
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
	if boxGraphic!=nil {
		graphicTmp, err := ebiten.NewImageFromImage(boxGraphic, ebiten.FilterNearest)
		if err!=nil {
			fmt.Println("failed to load graphic for box, using white instead")
			tmp.Graphic.Fill(color.White)
		}else {
			tmp.Graphic = graphicTmp
		}
	}else{
		fmt.Println("boxGraphic not set, using white instead")
		tmp.Graphic, _=ebiten.NewImage(int(math.Abs(max.X-min.X))+1,
			int(math.Abs(max.Y-min.Y))+1,
			ebiten.FilterDefault)
		tmp.Graphic.Fill(color.White)
	}
	tmp.Opts = &ebiten.DrawImageOptions{}
	tmp.Opts.GeoM.Translate(min.X, min.Y)
	return tmp
}

//u ovoj funkciji se pravi samo kvadrat bez postavljanja collider...
func newSpecialBox(min, max vector2) *box {
	if min.X > max.X || min.Y > max.Y {
		fmt.Printf("Invalid box: (%f,%f)(%f,%f)", min.X, min.Y, max.X, max.Y)
		return nil
	}
	tmp := new(box)
	if special_box_graphic!=nil {
		graphicTmp, err := ebiten.NewImageFromImage(special_box_graphic, ebiten.FilterNearest)
		if err!=nil {
			fmt.Println("failed to load graphic for box, using white instead")
			tmp.Graphic.Fill(color.White)
		}else {
			tmp.Graphic = graphicTmp
		}
	}else{
		fmt.Println("boxGraphic not set, using white instead")
		tmp.Graphic, _=ebiten.NewImage(int(math.Abs(max.X-min.X))+1,
			int(math.Abs(max.Y-min.Y))+1,
			ebiten.FilterDefault)
		tmp.Graphic.Fill(color.White)
	}
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
	if triangleGraphic!=nil {
		graphicTmp, err := ebiten.NewImageFromImage(triangleGraphic, ebiten.FilterNearest)
		if err!=nil {
			fmt.Println("failed to load graphic for box, using white instead")
			tmp.Graphic.Fill(color.White)
		}else {
			tmp.Graphic = graphicTmp
		}
	}else{
		fmt.Println("boxGraphic not set, using white instead")
		tmp.Graphic, _=ebiten.NewImage(int(math.Abs(max.X-min.X))+1,
			int(math.Abs(max.Y-min.Y))+1,
			ebiten.FilterDefault)
		tmp.Graphic.Fill(color.White)
	}
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

func newHole(min, max vector2) *hole{
	if min.X > max.X || min.Y > max.Y {
		fmt.Printf("Invalid box: (%f,%f)(%f,%f)", min.X, min.Y, max.X, max.Y)
		return nil
	}
	tmp := new(hole)
	tmp.Collider.Min = min
	tmp.Collider.Max = max
	tmp.Collider.Mid = newV2((min.X+max.Y)/2, (min.Y+max.Y)/2)
	if holeGraphic!=nil {
		graphicTmp, err := ebiten.NewImageFromImage(holeGraphic, ebiten.FilterNearest)
		if err!=nil {
			fmt.Println("failed to load graphic for hole, using white instead")
			tmp.Graphic.Fill(color.White)
		}else {
			tmp.Graphic = graphicTmp
		}
	}else{
		fmt.Println("HoleGraphic not set, using white instead")
		tmp.Graphic, _=ebiten.NewImage(int(math.Abs(max.X-min.X))+1,
			int(math.Abs(max.Y-min.Y))+1,
			ebiten.FilterDefault)
		tmp.Graphic.Fill(color.White)
	}
	tmp.Opts = &ebiten.DrawImageOptions{}
	tmp.Opts.GeoM.Translate(min.X, min.Y)
	return tmp
}