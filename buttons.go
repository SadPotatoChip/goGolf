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

var main_menu_already_set = false

var is_1_menu_button_sellected = false
var is_menu_button_2_sellected = false
var is_menu_button_3_sellected = false
var is_menu_button_4_sellected = false
var is_menu_oppened = false

func checkButtonClicks(){

	// PLAY button
	if x_pos > 500 && x_pos < 720 && y_pos > 200 && y_pos < 260 && is_main_menu {
		//play_button_is_sellected()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_first_level()
		}
	}

	// MENU button
	if x_pos > 1140 && x_pos < 1170 && y_pos > 10 && y_pos < 50 && is_main_menu {
		//if mouseButtonDown(ebiten.MouseButtonLeft) {
			is_menu_oppened = true
			create_first_menu_button()
			create_second_menu_button()
			create_third_menu_button()
			create_forth_menu_button()
			main_menu_already_set = false
		//}
	} else {
		if main_menu_already_set == false && !(x_pos > 1000 && x_pos < 1190 && y_pos > 60 && y_pos < 300) && is_main_menu == true {
			set_main_menu()
			is_menu_oppened = false
			main_menu_already_set = true
		}
	}

// 1 MENU OPTION
// (1000, 60), (1190, 120)))
	if x_pos > 1000 && x_pos < 1190 && y_pos > 60 && y_pos < 120 && is_main_menu && is_menu_oppened && !(page_num == 1){
		is_1_menu_button_sellected = true
		create_first_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_first_page()
			is_main_menu = false
		}
	} else {
		if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 120 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu){
			is_1_menu_button_sellected = false
			create_first_menu_button()
			is_main_menu = true
		}
	}
	//}
// 2 MENU OPTION
// (1000, 120), (1190, 180)

	if x_pos > 1000 && x_pos < 1190 && y_pos > 120 && y_pos < 180 && is_main_menu && is_menu_oppened{	//TODO
		is_menu_button_2_sellected = true
		create_second_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_second_page()
			is_main_menu = false
		}
	} else {
		if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 120 && y_pos < 180 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu){
			is_menu_button_2_sellected = false
			create_second_menu_button()
		}
	}

// 3 MENU OPTION
// (1000, 180), (1190, 240)

	if x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu {
		is_menu_button_3_sellected = true
		create_third_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_third_page()
			is_main_menu = false
		}
	} else {
		if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu){
			is_menu_button_3_sellected = false
			create_third_menu_button()
		}
	}

// 4 MENU OPTION
// (1000, 240), (1190, 300)

	if x_pos > 1000 && x_pos < 1190 && y_pos > 240 && y_pos < 300 && is_main_menu {
		is_menu_button_4_sellected = true
		create_forth_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			os.Exit(0)
		}
	} else {
		if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 240 && y_pos < 300 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu){
			is_menu_button_4_sellected = false
			create_forth_menu_button()
		}
	}

// kad se klikne na grid treba da se vrati u meni
// (1150, 20), (1180, 50)
/*
	if x_pos > 1150 && x_pos < 1180 && y_pos > 20 && y_pos < 50 && (is_main_menu == false) {
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_main_menu()
		}
	}
*/
// next
// (605, 420), (730, 600)

	if x_pos > 605 && x_pos < 730 && y_pos > 420 && y_pos < 600 && (is_main_menu == false) && mouseButtonDown(ebiten.MouseButtonLeft) {
		fmt.Println("dadada")
		switch (photo_num){
			case 1:	second_photo()
			case 2: third_photo()
			case 3: forth_photo()
			case 4: fifth_photo()
			case 5:
		}
	}
}
// previous
// (630, 400), (730, 500)
/*
	if x_pos > 630 && x_pos < 730 && y_pos > 400 && y_pos < 500 && (is_main_menu == false) && mouseButtonDown(ebiten.MouseButtonLeft) {
		fmt.Println("dadada")
		switch (photo_num){
			case 1:	second_photo()
				fmt.Println("1")
			case 2: third_photo()
			case 3: 
		}
	}
*/
/*
	if ((x_pos - x_center) * (x_pos - x_center) + (y_pos - y_center) * (y_pos - y_center)) <= (radius * radius) && mouseButtonDown(ebiten.MouseButtonLeft) {
		switch (level_num){
			case 0:
			case 1: set_second_level()
			case 2: set_third_level()
			case 3: set_forth_level()
			case 4: set_fifth_level()
			case 5: set_main_menu()
		}
	}
*/

func check_pressed_keys() {

  	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		fmt.Println("exiting...")
		os.Exit(0)
	}
}

