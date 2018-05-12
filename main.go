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

	reader4, _ := os.Open(box_str)
	special_box_graphic, _, _ = image.Decode(reader4)

	reader2, _ := os.Open(backgroung_str)
	tmp, _, _ := image.Decode(reader2)

	reader6, _ := os.Open(hole_str)
	holeGraphic, _, _ = image.Decode(reader6)

	testBackground, _ := ebiten.NewImageFromImage(tmp, ebiten.FilterDefault)
	backgroundImage = uninteractableImage{testBackground,&ebiten.DrawImageOptions{}}
}

func update(screen *ebiten.Image) error {
	backgroundImage.draw(screen)

	x_pos, y_pos = ebiten.CursorPosition()

	if levelIsInstantiating {
		lvl.Instantiate("testlvl1.json")

		set_main_menu()
                
                /* don't touch this!
                you're a potato :p

		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 0; i++ {
			x, y := rand.Float64()*screenWidth, rand.Float64()*screenHeight
			lvl.addBox(newBox(newV2(x, y), newV2(x+20, y+30)))

		}*/


		levelIsInstantiating = false
	}
	
	checkButtonClicks()
	check_pressed_keys()



	handleInput()
        
	player.applyNaturalForces()
	player.move()

	if checkIfBallIsInHole(){
		switch (level_num){
			case 1: set_second_level()
			case 2: set_third_level()
			case 3: set_forth_level()
			case 4: set_fifth_level()
			case 5: set_main_menu()
		}
		fmt.Println("GG")		// kad loptica upadne u rupu
	}	// player.horizontalSpeed = 0 da stane
		// ....verticalSpeed = 0
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
	
	for i := 0; i < lvl.num_of_uninteractable_image; i++ {
			screen.DrawImage(lvl.uninteractable_image_array[i].Graphic,
				lvl.uninteractable_image_array[i].Opts)
	}
	
	if lvl.hole!=nil {
		screen.DrawImage(lvl.hole.Graphic, lvl.hole.Opts)
	}
}

func drawPlayer(screen *ebiten.Image) {
	screen.DrawImage(player.graphic, player.opts)
	if player.isGrounded && player.horisonatalSpeed==0{
		for i := 0; i < len(player.indicatorGhost); i++ {
			screen.DrawImage(player.indicatorGhost[i].graphic,player.indicatorGhost[i].opts)
		}
	}
	//screen.DrawImage(player.collisonGhost.graphic,player.collisonGhost.opts)
}

func handleInput() {
	if player.isGrounded && player.horisonatalSpeed==0{
		if ebiten.IsKeyPressed(player.controls.angleUpKey) {
			player.controls.changeAngle(+1)
			player.setIndicators()
		}
		if ebiten.IsKeyPressed(player.controls.angleDownKey) {
			player.controls.changeAngle(-1)
			player.setIndicators()
		}
		if ebiten.IsKeyPressed(player.controls.powerUpKey) {
			player.controls.changePower(+1)
			player.setIndicators()
		}
		if ebiten.IsKeyPressed(player.controls.powerDownKey) {
			player.controls.changePower(-1)
			player.setIndicators()
		}
		if hitKeyDown(player.controls.hitKey) {
			player.hit(player.controls.angle, player.controls.power)

		}
	}
}

func checkIfBallIsInHole() bool{
	if lvl.hole!=nil{
		if player.position.X+player.size > lvl.hole.Collider.Min.X && player.position.X < lvl.hole.Collider.Max.X &&
			player.position.Y+player.size > lvl.hole.Collider.Min.Y && player.position.Y < lvl.hole.Collider.Max.Y {
			return true
		}
		return false
	}
	return false
}
