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
	//"io/ioutil"
	//"encoding/json"
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
	uninteractable_image_array [maxLevelObjects] uninteractableImage
	num_of_uninteractable_image int
	NumOfShapes     int
	ShapeTypes	string
}

func (l *level) Instantiate(filePath string) {
	l.ShapeTypes=""
	if levelIsInstantiating==false {
		player = makeBall(300, 300, false, "images/main_menu/golf-ball.png")
	}else{
		player = makeBall(-1000, -1000, false, "images/main_menu/golf-ball.png")
	}

	//fmt.Println(jsonData)


	//clear level
	l.NumOfShapes = 0
	l.MaxSortedShapes = [maxLevelObjects]shape{}
	l.MinSortedShapes =	[maxLevelObjects]shape{}
	l.uninteractable_image_array = [maxLevelObjects]uninteractableImage{}
	l.num_of_uninteractable_image = 0
	ground := newBox(newV2(0, screenHeight), newV2(screenWidth, screenHeight+90))
	ceiling:= newBox(newV2(0,-80), newV2(screenWidth-1, 0))
	wallLeft:= newBox(newV2(0,-80), newV2(0, screenHeight))
	wallRight:= newBox(newV2(screenWidth,0), newV2(screenWidth+90, screenHeight))
	l.addBox(ground)
	l.addBox(ceiling)
	l.addBox(wallLeft)
	l.addBox(wallRight)

	//debug


	//l.hole=newHole(vector2{340,260},vector2{600,600})


}
/*
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
*/
func (l *level) add_uninteractable_image (u uninteractableImage) {
    l.uninteractable_image_array[l.num_of_uninteractable_image] = u
    l.num_of_uninteractable_image++
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
func newSpecialBox(min, max vector2) uninteractableImage {
	if min.X > max.X || min.Y > max.Y {
		fmt.Printf("Invalid box: (%f,%f)(%f,%f)", min.X, min.Y, max.X, max.Y)
		return uninteractableImage{}  // konstruktor, napravi strukturu i vrati strukturu a ne pointer
	}
	tmp := uninteractableImage{}
        
        graphicTmp, err := ebiten.NewImageFromImage(special_box_graphic, ebiten.FilterNearest)
		if err!=nil {
			fmt.Println("failed to load graphic for box, using white instead")
			tmp.Graphic.Fill(color.White)
		}else {
			tmp.Graphic = graphicTmp
		}
		
        /*
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
	}*/
        
	tmp.Opts = &ebiten.DrawImageOptions{}
	tmp.Opts.GeoM.Translate(min.X, min.Y)
	return tmp
}

func newTriangle(min, max vector2, side string) *triangle {
	graphicSize:=100.0
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
			fmt.Println("failed to load graphic for triangle, using white instead")
			tmp.Graphic.Fill(color.White)
		}else {
			tmp.Graphic = graphicTmp
		}
	}else{
		fmt.Println("triangleGraphic not set, using white instead")
		tmp.Graphic, _=ebiten.NewImage(int(math.Abs(max.X-min.X)),
			int(math.Abs(max.Y-min.Y)),
			ebiten.FilterDefault)
		tmp.Graphic.Fill(color.White)
	}
	tmp.Opts = &ebiten.DrawImageOptions{}
	tmp.Opts.GeoM.Scale((tmp.Collider.Max.X-tmp.Collider.Min.X)/graphicSize,
		(tmp.Collider.Max.Y-tmp.Collider.Min.Y)/graphicSize)
	tmp.Collider.Missing_part=side


	switch side{
	case "top-left":
		tmp.Opts.GeoM.Rotate(-math.Pi/2)
		tmp.Opts.GeoM.Translate(min.X, min.Y+(max.Y-min.Y))
	case "top-right":
		tmp.Opts.GeoM.Translate(min.X, min.Y)
	case "bottom-left":
		tmp.Opts.GeoM.Rotate(math.Pi)
		tmp.Opts.GeoM.Translate(min.X+(max.X-min.X), min.Y+(max.Y-min.Y))
	case "bottom-right":
		tmp.Opts.GeoM.Rotate(math.Pi/2)
		tmp.Opts.GeoM.Translate(min.X+(max.X-min.X), min.Y)
	}



	//debug
	drawDebugSquares(tmp)

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

func drawDebugSquares(tmp *triangle){
	debugBox1:=newBox(tmp.Collider.Max,vector2{tmp.Collider.Max.X+10,tmp.Collider.Max.Y+10})
	debugBox2:=newBox(tmp.Collider.Min,vector2{tmp.Collider.Min.X+10,tmp.Collider.Min.Y+10})
	debugBox1.Graphic,_=ebiten.NewImage(10,10,ebiten.FilterDefault)
	debugBox2.Graphic,_=ebiten.NewImage(10,10,ebiten.FilterDefault)
	debugBox1.Graphic.Fill(color.RGBA{255, 0, 0, 255})
	debugBox2.Graphic.Fill(color.White)

	lvl.addBox(debugBox1)
	lvl.addBox(debugBox2)
}
