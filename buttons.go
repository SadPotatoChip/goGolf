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

var is_1_menu_button_sellected = false
var is_menu_button_2_sellected = false
var is_menu_button_3_sellected = false
var is_menu_button_4_sellected = false
var is_menu_oppened = false

func checkButtonClicks(){
	if x_pos > 500 && x_pos < 720 && y_pos > 200 && y_pos < 260 && is_main_menu {
		//play_button_is_sellected()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_first_level()
		}
	} else {	// else mora ovako } else { inace ne radi
		//clear_play_button()
	}
/*
	if x_pos > 500 && x_pos < 720 && y_pos > 300 && y_pos < 360 && is_main_menu {
		levels_button_is_sellected()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_all_levels()
		}
	} else {
		clear_levels_button()
	}
*/
	if x_pos > 1140 && x_pos < 1170 && y_pos > 10 && y_pos < 40 && is_main_menu {
		//if mouseButtonDown(ebiten.MouseButtonLeft) {
			is_menu_oppened = true
			create_first_menu_button()
			create_second_menu_button()
			create_third_menu_button()
			create_forth_menu_button()
		//}
	} else {
			is_menu_oppened = false
			clear_first_menu_button()
			clear_second_menu_button()
			clear_third_menu_button()
			clear_forth_menu_button()
	}

// (1000, 60), (1190, 120)))
	if x_pos > 1000 && x_pos < 1190 && y_pos > 60 && y_pos < 120 && is_main_menu && is_menu_oppened{
		is_1_menu_button_sellected = true
		create_first_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_first_page()
		}
	} else {
		if is_menu_oppened {
			is_1_menu_button_sellected = true
			create_first_menu_button()
		}
	}

// (1000, 120), (1190, 180)
	if x_pos > 1000 && x_pos < 1190 && y_pos > 120 && y_pos < 180 && is_main_menu {
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_second_page()
		}
	}

// (1000, 180), (1190, 240)
	if x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu {
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_third_page()
		}
	}

// (1000, 240), (1190, 300)
	if x_pos > 1000 && x_pos < 1190 && y_pos > 240 && y_pos < 300 && is_main_menu {
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			os.Exit(0)
		}
	}

// kad se klikne na grid treba da se vrati u meni
// (1150, 20), (1180, 50)

	if x_pos > 1150 && x_pos < 1180 && y_pos > 20 && y_pos < 50 && (is_main_menu == false) {
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_main_menu()
		}
	}

// next
// (530, 400), (630, 500)
	if x_pos > 530 && x_pos < 630 && y_pos > 400 && y_pos < 500 && (is_main_menu == false) && mouseButtonDown(ebiten.MouseButtonLeft) {
		fmt.Println("dadada")
		switch (photo_num){
			case 1:	
			case 2: first_photo()
			case 3: second_photo()
		}
	}

// previous
// (630, 400), (730, 500)
	if x_pos > 630 && x_pos < 730 && y_pos > 400 && y_pos < 500 && (is_main_menu == false) && mouseButtonDown(ebiten.MouseButtonLeft) {
		fmt.Println("dadada")
		switch (photo_num){
			case 1:	second_photo()
				fmt.Println("1")
			case 2: third_photo()
			case 3: 
		}
	}

	if ((x_pos - x_center) * (x_pos - x_center) + (y_pos - y_center) * (y_pos - y_center)) <= (radius * radius) && mouseButtonDown(ebiten.MouseButtonLeft) {
		switch (level_num){
			case 0:	// kad je u meniju ne reaguje na next dugme (i ne postoji u meniju)
			case 1: set_second_level()
			/*case 2: set_third_level()
			case 3: set_forth_level()
			case 4: set_fifth_level()
			case 5: set_main_menu()*/
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
