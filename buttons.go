package main

import (
	"fmt"
	"os"
	"github.com/hajimehoshi/ebiten"
)

var x_pos, y_pos int

var x_center = 1150
var y_center = 60
var radius = 50

func checkButtonClicks(){
	if x_pos > 500 && x_pos < 720 && y_pos > 200 && y_pos < 260 && is_main_menu {
		play_button_is_sellected()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			fmt.Println("pointer is on the first button")
			set_first_level()
		}
	} else {	// else mora ovako } else { inace ne radi
		clear_play_button()
	}

	if x_pos > 500 && x_pos < 720 && y_pos > 300 && y_pos < 360 && is_main_menu {
		levels_button_is_sellected()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			fmt.Println("pointer is on the second button")
			set_all_levels()
		}
	} else {
		clear_levels_button()
	}

	if ((x_pos - x_center) * (x_pos - x_center) + (y_pos - y_center) * (y_pos - y_center)) <= (radius * radius) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		fmt.Println("pointer is on the NEXT button")
		switch (level_num){
			case 0:	// kad je u meniju ne reaguje na next dugme (i ne postoji u meniju)
			case 1: set_second_level()
			case 2: set_third_level()
			case 3: set_forth_level()
			case 4: set_fifth_level()
			case 5: set_main_menu()
		}
	}
}

func check_pressed_keys() {
  	if ebiten.IsKeyPressed(ebiten.Key1) && all_levels {
		fmt.Println("go back")
		set_main_menu()
	}

  	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		fmt.Println("exiting...")
		os.Exit(0)
	}
}
