package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten"
)

var x_pos, y_pos int
var other_ball = false
var cont = false

var opt = false
var lev = false
var x_center = 1150
var y_center = 60
var radius = 50

var main_menu_already_set = false

var is_1_menu_button_sellected = false
var is_menu_button_2_sellected = false
var is_menu_button_3_sellected = false
var is_menu_button_4_sellected = false
var is_menu_oppened = false

func checkButtonClicks() {

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
		//create_forth_menu_button()
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
	if x_pos > 1000 && x_pos < 1190 && y_pos > 60 && y_pos < 120 && is_main_menu && is_menu_oppened && !(page_num == 1) {
		is_1_menu_button_sellected = true
		create_first_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			cont = true
			is_main_menu = false
			set_first_page()
		}
	} else {
		if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 120 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu) {
			is_1_menu_button_sellected = false
			create_first_menu_button()
			is_main_menu = true
		}
	}
	//}
	// 2 MENU OPTION
	// (1000, 120), (1190, 180)

	if x_pos > 1000 && x_pos < 1190 && y_pos > 120 && y_pos < 180 && is_main_menu && is_menu_oppened { //TODO
		is_menu_button_2_sellected = true
		create_second_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			lev = true
			is_main_menu = false
			set_second_page()
		}
	} else {
		if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 120 && y_pos < 180 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu) {
			is_menu_button_2_sellected = false
			create_second_menu_button()
		}
	}

	// 3 MENU OPTION
	// (1000, 180), (1190, 240)
	/*
		if x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu && is_menu_oppened{
			is_menu_button_3_sellected = true
			create_third_menu_button()
			if mouseButtonDown(ebiten.MouseButtonLeft) {
				opt = true
				is_main_menu = false
				set_third_page()
			}
		} else {
			if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu){
				is_menu_button_3_sellected = false
				create_third_menu_button()
			}
		}
	*/
	// 4 MENU OPTION
	// (1000, 240), (1190, 300)

	if x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu && is_menu_oppened {
		is_menu_button_3_sellected = true
		create_third_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			os.Exit(0)
		}
	} else {
		if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu) {
			is_menu_button_3_sellected = false
			create_third_menu_button()
		}
	}

	// kad se klikne na grid treba da se vrati u meni
	// (1150, 20), (1180, 50)

	if x_pos > 1150 && x_pos < 1180 && y_pos > 20 && y_pos < 50 && (is_main_menu == false) && (lev || opt || cont) {
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_main_menu()
		}
	}

	// (150, 100), (550, 500) izaberi lopticu
	if x_pos > 360 && x_pos < 470 && y_pos > 250 && y_pos < 350 && is_main_menu == false && opt == true {
		fmt.Printf("crvena: %d   %d", x_pos, y_pos)
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			str_2 = "images/main_menu/tennis-ball.png"
		}
	} else {
		if x_pos > 230 && x_pos < 330 && y_pos > 250 && y_pos < 350 && is_main_menu == false && opt == true {
			{
				fmt.Printf("bela: %d   %d", x_pos, y_pos)
				if mouseButtonDown(ebiten.MouseButtonLeft) {
					str_2 = "images/main_menu/golf-ball.png"
				}
			}
		} else {
			if is_main_menu == false && opt == true {
				set_third_page()
			}
		}
	}
	/*
			if mouseButtonDown(ebiten.MouseButtonLeft) {
				os.Exit(0)
			}
		} else {
			if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 240 && y_pos < 300 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu){
				is_menu_button_4_sellected = false
				create_forth_menu_button()
			}
		}
	*/

	// (300, 100), (600, 200) predji na drugi nivo
	if x_pos > 50 && x_pos < 270 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_first_level()
	}

	// (300, 100), (600, 200) predji na drugi nivo
	if x_pos > 270 && x_pos < 490 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_second_level()
	}
	// (300, 100), (600, 200) predji na drugi nivo
	if x_pos > 490 && x_pos < 710 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_third_level()
	}
	// (300, 100), (600, 200) predji na drugi nivo
	if x_pos > 710 && x_pos < 930 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_forth_level()
	}
	// (300, 100), (600, 200) predji na drugi nivo
	if x_pos > 930 && x_pos < 1150 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_fifth_level()
	}
	// next
	// (605, 420), (730, 600)
	//var yes = true

	/*
	   var counter = 0
	   	if x_pos > 605 && x_pos < 730 && y_pos > 420 && y_pos < 600 && (is_main_menu == false) && mouseButtonDown(ebiten.MouseButtonLeft) {
	   		//counter++
	   		fmt.Println("klik 1")
	   		switch (photo_num){
	   			case 1:	second_photo()
	   				//yes = !yes
	   				//fmt.Println(yes)
	   			case 2: if counter % 3 != 0 {
	   					//fmt.Println(yes)
	   					third_photo()
	   					}
	   			case 3 : if counter % 3 != 0 {
	   					forth_photo()
	   					}
	   			case 4 : if counter % 3 != 0 {
	   					fifth_photo()
	   					no_more_levels_next = true
	   					change_next_button()
	   				}
	   			case 5:
	   		}
	   	}
	*/

	// previous
	// (505, 420), newV2(630, 500)
	/*
		if x_pos > 505 && x_pos < 630 && y_pos > 420 && y_pos < 500 && (is_main_menu == false) && mouseButtonDown(ebiten.MouseButtonLeft) {
			fmt.Println("1 klik")
			switch (photo_num){
				case 1:
				case 2: first_photo()
					no_more_levels_back = true
					change_previous_button()
				case 3: second_photo()
					no_more_levels_back = false
					previous_photo()
				case 4: third_photo()
					no_more_levels_back = false
					previous_photo()
				case 5: forth_photo()
					no_more_levels_back = false
					previous_photo()
			}
		}
	*/
}

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

	if ebiten.IsKeyPressed(ebiten.Key1) {
		set_main_menu()
	}
}
