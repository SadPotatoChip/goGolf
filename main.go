package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
        "os"
        "image"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type vector2 struct {
	x, y float64
}

func newV2(x, y float64) vector2 {
	tmp := new(vector2)
	tmp.x = x
	tmp.y = y
	return *tmp
}

const screenWidth float64 = 1200
const screenHeight float64 = 600

const testBallsLen = 50

var levelIsInstantiating = true
var lvl level
var player *ball

var testBalls [testBallsLen]*ball

var triangle_graphic image.Image

var test_triangle *triangle
var test_triangle2 *triangle

func main() {
    
        reader, _ := os.Open("triangle.png")
        triangle_graphic,_, _ = image.Decode(reader)
        
	var w, h int = int(screenWidth), int(screenHeight)
	if err := ebiten.Run(update, w, h, 1, "puff puff"); err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {
	if levelIsInstantiating {
		player = makeBall(500, 500,false)
		lvl.Instantiate()

		//test
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 35; i++ {
			x, y := rand.Float64()*screenWidth, rand.Float64()*screenHeight
			lvl.addBox(newBox(newV2(x, y), newV2(x+10, y+10)))

		}
                test_triangle = newTriangle(vector2{screenWidth/2,screenHeight/2}, vector2{screenWidth/2+50,screenHeight/2+50}, "top-left")
                test_triangle2 = newTriangle(vector2{200,200}, vector2{300,300}, "bottom-left")
              //  lvl.addBox(test_triangle)
                //lvl.addBox(test_triangle2)
                
		levelIsInstantiating = false
	}
	handleInput()
        screen.DrawImage(test_triangle.graphic, test_triangle.opts)
        screen.DrawImage(test_triangle2.graphic, test_triangle2.opts)
        
	player.applyNaturalForces()
	player.move()

	drawLevel(screen)
	drawPlayer(screen)

	//DEBUG
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%f \nx:%f y:%f\nv:%f h:%f\npow:%f, angle:%f, hitHeld:%t", ebiten.CurrentFPS(), player.position.x, player.position.y, player.horisonatalSpeed, player.verticalSpeed, player.controls.power, player.controls.angle/math.Pi, hitKeyIsDown))

	return nil
}

func drawLevel(screen *ebiten.Image) {
	for i := 0; i < lvl.nOfShapes; i++ {
                tmp := *lvl.maxSortedShapes[i]
                switch (tmp.(type)) {
                    case triangle: temp := tmp.(triangle)
                                        screen.DrawImage(temp.getGraphic(), temp.getOpts())
                    case box: temp := tmp.(box)
                                        screen.DrawImage(temp.getGraphic(), temp.getOpts())
                }
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
		if hitKeyDown(player.controls.hitKey) {
			fmt.Printf("hit\n")
			player.hit(player.controls.angle, player.controls.power)

		}
	}
}
