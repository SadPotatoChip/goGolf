package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

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

const testBallsN = 50

var levelisInstantiating = true
var lvl level
var player *ball

var testBalls [testBallsN]*ball

func main() {
	var w, h int = int(screenWidth), int(screenHeight)
	if err := ebiten.Run(update, w, h, 1, "puff puff"); err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {
	if levelisInstantiating {
		player = makeBall(500, 500)
		lvl.Instantiate()
		//test
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 10; i++ {
			x, y := rand.Float64()*screenWidth, rand.Float64()*screenHeight
			lvl.addBox(newBox(newV2(x, y), newV2(x+200, y+200)))

		}

		levelisInstantiating = false
	}
	handleInput()

	player.applyNaturalForces()
	player.move()

	drawLevel(screen)
	drawPlayer(screen)

	//DEBUG
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%f \nx:%f y:%f\nv:%f h:%f\npow:%f, angle:%f, hitHeld:%t", ebiten.CurrentFPS(), player.position.x, player.position.y, player.horisonatalSpeed, player.verticalSpeed, player.controls.power, player.controls.angle/math.Pi, hitKeyIsDown))

	return nil
}

func drawLevel(screen *ebiten.Image) {
	for i := 0; i < lvl.nOfBoxes; i++ {
		screen.DrawImage(lvl.maxSortedBoxes[i].graphic, lvl.maxSortedBoxes[i].opts)
	}
}

func drawPlayer(screen *ebiten.Image) {
	screen.DrawImage(player.graphic, player.opts)

	//screen.DrawImage(player.controls.indicator.graphic, player.controls.indicator.opts)
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
