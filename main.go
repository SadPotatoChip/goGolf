package main

import (
	"fmt"
	"math"
	"os"
	"image"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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

var lvl level
var player *ball

func main() {
	//preprocess testing textures import
	prefetchGraphics()

	var w, h int = int(screenWidth), int(screenHeight)
	if err := ebiten.Run(update, w, h, 1, "puff puff"); err != nil {
		panic(err)
	}
}

func prefetchGraphics() {
	reader, _ := os.Open(triangle_str)
	triangleGraphic, _, _ = image.Decode(reader)

	reader3, _ := os.Open(box_str)
	boxGraphic, _, _ = image.Decode(reader3)

	reader2, _ := os.Open(backgroung_str)
	tmp, _, _ := image.Decode(reader2)
	testBackground, _ := ebiten.NewImageFromImage(tmp, ebiten.FilterDefault)
	backgroundImage = uninteractableImage{testBackground,&ebiten.DrawImageOptions{}}
}

func update(screen *ebiten.Image) error {
	backgroundImage.draw(screen)

	if levelIsInstantiating {
		player = makeBall(500, 500,false)
		lvl.Instantiate("testlvl1.json")

		set_first_level()
		current_level = 1
                
                /* don't touch this!
                you're a potato :p

		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 0; i++ {
			x, y := rand.Float64()*screenWidth, rand.Float64()*screenHeight
			lvl.addBox(newBox(newV2(x, y), newV2(x+20, y+30)))

		}*/
		//save level
		//lvl.makeJson("testlvl1")
		levelIsInstantiating = false
	}
        
  	if ebiten.IsKeyPressed(ebiten.Key1) && current_level == 1 {
  		fmt.Println("da")
  		set_second_level()
  		current_level++
	}

	handleInput()
        
	player.applyNaturalForces()
	player.move()

	drawLevel(screen)
	drawPlayer(screen)

	//DEBUG
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%f \nx:%f y:%f\nh:%f v:%f\npow:%f, angle:%f, hitHeld:%t",
				ebiten.CurrentFPS(), player.position.X, player.position.Y, player.horisonatalSpeed, 					player.verticalSpeed, player.controls.power, player.controls.angle/math.Pi, hitKeyIsDown))

	return nil
}

func drawLevel(screen *ebiten.Image) {
	for i := 0; i < lvl.NumOfShapes; i++ {
			screen.DrawImage(lvl.MaxSortedShapes[i].getGraphic(),
				lvl.MaxSortedShapes[i].getOpts())
	}
}

func drawPlayer(screen *ebiten.Image) {
	screen.DrawImage(player.graphic, player.opts)
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
			player.hit(player.controls.angle, player.controls.power)

		}
	}
}
