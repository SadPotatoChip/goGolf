package main

import (
	"fmt"
	"os"
	"image"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"log"
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
var shotsTaken int

var audioContext *audio.Context
var audioPlayerMusic *audio.Player
var audioPlayerClick *audio.Player
var audioPlayerHit *audio.Player
var audioPlayerWin *audio.Player

var songPaths []string
var selectedSong=0

func main() {


	ac, err1:=audio.NewContext(44100)
	if err1!=nil{
		log.Fatal(err1)
	}
	audioContext=ac
	songPaths=[]string{"music_1.mp3","mainSong2.mp3","mainSong3.mp3"}
	audioPlayerMusic=newAudioPlayer("sounds/"+songPaths[0])
	audioPlayerClick=newAudioPlayer("sounds/click.mp3")
	audioPlayerHit=newAudioPlayer("sounds/strike.mp3")
	audioPlayerWin=newAudioPlayer("sounds/win.mp3")
	audioPlayMainTrack()

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
		lvl.Instantiate(100, 100)
		set_main_menu()
		levelIsInstantiating = false
	}
	checkButtonClicks()
	check_pressed_keys()

	handleInput()
	if !audioPlayerMusic.IsPlaying(){
		audioPlayerMusic.Rewind()
		audioPlayerMusic.Play()
	}

	player.applyNaturalForces()
	player.move()

	if checkIfBallIsInHole() {
		audioPlayerWin.Play()
		audioPlayerWin.Rewind()
		switch level_num {
		case 1:
			set_second_level()
		case 2:
			set_third_level()
		case 3:
			set_forth_level()
		case 4:
			set_fifth_level()
		case 5:
			level_num = 0
			set_main_menu()
		}
		fmt.Println("GG") // kad loptica upadne u rupu
	}
	drawLevel(screen)
	drawPlayer(screen)

	//DEBUG
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS:%f \nx:%f y:%f\nh:%f v:%f\npow:%f, angle:%f, hitHeld:%t",
	//			ebiten.CurrentFPS(), player.position.X, player.position.Y, player.horisonatalSpeed, 					player.verticalSpeed, player.controls.power, player.controls.angle/math.Pi, hitKeyIsDown))

	if shotsTaken != -1 {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Shots Taken %d", shotsTaken))
	}
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
		if hitKeyDown(player.controls.hitKey, player) {
			audioPlayerHit.Play()
			audioPlayerHit.Rewind()
			shotsTaken++
			player.hit(player.controls.angle, player.controls.power)

		}
	}
	//alternate songs
	if hitKeyDown(ebiten.KeyBackslash, player){
		audioNextSong()
	}
	if hitKeyDown(ebiten.KeyF, player){
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
			ebiten.SetScreenSize(1200,600)
		}else{
			ebiten.SetFullscreen(true)
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

func audioNextSong(){
	selectedSong++
	if selectedSong >= len(songPaths){
		selectedSong=0
	}
	audioPlayMainTrack()
}

func audioPlayMainTrack(){
	audioPlayerMusic.Close()
	audioPlayerMusic=newAudioPlayer("sounds/"+songPaths[selectedSong])
	audioPlayerMusic.Rewind()
	go func(){
		audioPlayerMusic.Play()
	}()
}

func newAudioPlayer(path string) (*audio.Player){

	f, err2 := ebitenutil.OpenFile(path)
	if err2!=nil{
		log.Fatal(err2)
	}
	d, err3 := mp3.Decode(audioContext, f)
	if err3!=nil{
		log.Fatal(err3)
	}
	ap, err4 := audio.NewPlayer(audioContext, d)
	if err3!=nil{
		log.Fatal(err4)
	}
	return ap

}