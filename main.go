package main

import (
	"fmt"
	"math"
        "os"
        "image"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	//"math/rand"
	//"time"
)

type vector2 struct {
	X, Y float64
}

func newV2(x, y float64) vector2 {
	tmp := new(vector2)
	tmp.X = x
	tmp.Y = y
	return *tmp
}


const screenWidth float64 = 1200
const screenHeight float64 = 600

var current_level = 0

var levelIsInstantiating = true
var lvl level
var player *ball

var backgroundImage uninteractableImage


//debug
var triangleGraphic image.Image
var boxGraphic image.Image

var triangle_str = "triangle.png"
var box_str = "box.png"
var backgroung_str = "1.png"

// textures
//var texture_array

func main() {
	//preprocess testing textures import
    
    if current_level == 1 {
        triangle_str = "triangle1000x1000.png"
        box_str = "triangle1000x1000.png"
        backgroung_str = "1.png"
    }
    
	reader, _ := os.Open(triangle_str)
	triangleGraphic,_, _ = image.Decode(reader)
	
        reader3, _ := os.Open(box_str)
	boxGraphic,_, _ = image.Decode(reader3)
        
        reader2, _ := os.Open(backgroung_str)
	tmp,_, _ := image.Decode(reader2)
	testBackground,_:=ebiten.NewImageFromImage(tmp,ebiten.FilterDefault)

	backgroundImage = uninteractableImage{testBackground,&ebiten.DrawImageOptions{}}
        
	var w, h int = int(screenWidth), int(screenHeight)
	if err := ebiten.Run(update, w, h, 1, "puff puff"); err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {
	backgroundImage.draw(screen)

	if levelIsInstantiating {
		player = makeBall(500, 500,false)
		lvl.Instantiate("testlvl1.json")

                set_first_level()
                current_level = 1
                
                /* don't touch this!
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 0; i++ {
			x, y := rand.Float64()*screenWidth, rand.Float64()*screenHeight
			lvl.addBox(newBox(newV2(x, y), newV2(x+20, y+30)))

		}*/

        

		//save level
		//lvl.makeJson("testlvl1")
		levelIsInstantiating = false
	}
        
  // TODO Izdvojicu sve fje u vezi pravljenja nivoa u poseban fajl, samo hocu prvo da mi sve proradi!
        
        if ebiten.IsKeyPressed(ebiten.KeySpace) && current_level == 1 {
            fmt.Println("da")
            set_second_level()
            current_level++
        }
   
   // TODO ako stavim dva if-a ukljucuje se samo nivo u poslednjem if-u
  /*
    if ebiten.IsKeyPressed(ebiten.KeySpace) && current_level == 1 {
            fmt.Println("daa")
            set_third_level()
            current_level++
        }
        
    */    
  // TODO ne mogu da promenim background i slicice za trougao i kvadrat

	handleInput()
        
	player.applyNaturalForces()
	player.move()

	drawLevel(screen)
	drawPlayer(screen)

	//DEBUG
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%f \nx:%f y:%f\nh:%f v:%f\npow:%f, angle:%f, hitHeld:%t", ebiten.CurrentFPS(), player.position.X, player.position.Y, player.horisonatalSpeed, player.verticalSpeed, player.controls.power, player.controls.angle/math.Pi, hitKeyIsDown))

	return nil
}

func set_first_level() {
    
//                triangle_str = "triangle.png"
 //               box_str = "box.str"
   //             backgroung_str = "1.png"
                
    		// levo
                lvl.addTriangle(newTriangle(vector2{400,500}, vector2{700,700}, "top-right"))
                lvl.addTriangle(newTriangle(vector2{200,400}, vector2{700,700}, "top-right"))
                
                lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
                lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
                lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))                
                lvl.addBox(newBox(newV2(300, 500), newV2(400, 600)))
                lvl.addBox(newBox(newV2(0, 400), newV2(100, 500)))
                lvl.addBox(newBox(newV2(100, 400), newV2(200, 500)))
                
                // desno
                lvl.addTriangle(newTriangle(vector2{900,500}, vector2{1200,700}, "top-left"))
                lvl.addTriangle(newTriangle(vector2{700,600}, vector2{1200,900}, "top-left"))
                
                lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
                lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))
                lvl.addBox(newBox(newV2(1000, 500), newV2(1100, 600)))
                lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))
                lvl.addBox(newBox(newV2(1000, 400), newV2(1100, 500)))
                lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
}

func set_second_level() {
    
        lvl.Instantiate("lala")
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
        
}

func set_third_level() {
    
        lvl.Instantiate("lalaa")

                lvl.addBox(newBox(newV2(1000, 500), newV2(1100, 600)))
                lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))

        
}

func drawLevel(screen *ebiten.Image) {
	for i := 0; i < lvl.NumOfShapes; i++ {
			screen.DrawImage(lvl.MaxSortedShapes[i].getGraphic(),
				lvl.MaxSortedShapes[i].getOpts())
	}
}

func drawPlayer(screen *ebiten.Image) {
	screen.DrawImage(player.graphic, player.opts)

	//debug
	screen.DrawImage(player.collisonGhost.graphic, player.collisonGhost.opts)
}

func handleInput() {
	if player.isGrounded {
		if ebiten.IsKeyPressed(player.controls.angleUpKey) {
			player.controls.changeAngle(+1)
		}
		if ebiten.IsKeyPressed(player.controls.angleDownKey) {
			player.controls.changeAngle(-1)
		}
		if ebiten.IsKeyPressed(player.controls.powerUpKey) {
			player.controls.changePower(+1)
		}
		if ebiten.IsKeyPressed(player.controls.powerDownKey) {
			player.controls.changePower(-1)
		}
		/*if hitKeyDown(player.controls.hitKey) {
			fmt.Printf("hit\n")
			player.hit(player.controls.angle, player.controls.power)

		}*/
	}
}
