/*
 * u ovom fajlu su funkcije
 * koje postavljaju nivoe, izlaze iz aplikacije...
 */

package main

import (
	"fmt"
	"os"
	"github.com/hajimehoshi/ebiten"
)

var x_pos, y_pos int
var other_ball = false
var cont = false

/*
 *
 * indikator da li smo trenutno u options
 *
 */
var opt = false

/*
 *
 * indikator da li smo trenutno u levels
 *
 */
var lev = false

/*
 * kordinate centra i precnika okruglog dugmeta
 *
 */
var x_center = 1150
var y_center = 60
var radius = 50

/*
 *
 * indikator da li je meni vec ucitan
 *
 */
var main_menu_already_set = false

/*
 *
 * indikator koje je dugme iz menija (tj. opcija) selektovano
 *
 */
var is_1_menu_button_sellected = false
var is_menu_button_2_sellected = false
var is_menu_button_3_sellected = false
var is_menu_button_4_sellected = false
var is_menu_oppened = false

/*
 * proverava koji su dugmici kliknuti
 * i izvrsava odgovarajuce akcije
 *
 */
func check_button_clicks() {

	/* PLAY button is clicked --> pokreni igricu */
	play_button_is_clicked()

	/* MENU button */
	if is_menu_button_selected() {
		is_menu_oppened = true
		create_first_menu_button()
		create_second_menu_button()
		create_third_menu_button()
		main_menu_already_set = false
	} else {
		if main_menu_already_set == false && !(x_pos > 1000 && x_pos < 1190 && y_pos > 60 && y_pos < 300) && is_main_menu == true {
			set_main_menu()
			is_menu_oppened = false
			main_menu_already_set = true
		}
	}

	/*
	* prva opcija menija
 	*
 	*/
	if is_first_option_selected() {
		is_1_menu_button_sellected = true
		create_first_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			cont = true
			is_main_menu = false
			set_first_page()
		}
	} else {
		if is_not_first_option_selected() {
			is_1_menu_button_sellected = false
			create_first_menu_button()
			is_main_menu = true
		}
	}

	/*
 	* druga opcija menija
 	*
 	*/
	if is_second_option_selected() {
		is_menu_button_2_sellected = true
		create_second_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			lev = true
			is_main_menu = false
			set_second_page()
		}
	} else {
		if is_not_second_option_selected() {
			is_menu_button_2_sellected = false
			create_second_menu_button()
		}
	}

	/*
	 * 3. opcija menija
	 *
	 */

	if is_third_option_selected() {
		is_menu_button_3_sellected = true
		create_third_menu_button()
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			os.Exit(0)
		}
	} else {
		if is_not_third_option_selected() {
			is_menu_button_3_sellected = false
			create_third_menu_button()
		}
	}

	/*
	 * kad se klikne na grid, vraca se u meni
	 *
	 */
	if is_grid_clicked() {
		if mouseButtonDown(ebiten.MouseButtonLeft) {
			set_main_menu()
		}
	}

	/* izaberi lopticu */
	/* ova funkcija se ne koristi nigde zbog problema sa klikom!

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
	*/

	/*
	 * predji na prvi nivo
	 *
	 */
	if x_pos > 50 && x_pos < 270 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_first_level()
	}

	/*
	 * predji na drugi nivo
	 *
	 */
	if x_pos > 270 && x_pos < 490 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_second_level()
	}

	/*
	 * predji na treci nivo
	 *
	 */
	if x_pos > 490 && x_pos < 710 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_third_level()
	}

	/*
	 * predji na cetvrti nivo
	 *
	 */
	if x_pos > 710 && x_pos < 930 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_forth_level()
	}

	/*
	 * predji na peti nivo
	 *
	 */
	if x_pos > 930 && x_pos < 1150 && y_pos > 200 && y_pos < 360 && (is_main_menu == false) && lev == true && mouseButtonDown(ebiten.MouseButtonLeft) && level_num == 0 {
		set_fifth_level()
	}

	/*if ((x_pos - x_center) * (x_pos - x_center) + (y_pos - y_center) * (y_pos - y_center)) <= (radius * radius) && mouseButtonDown(ebiten.MouseButtonLeft) {
		level_num=0
		set_main_menu()
		}
	}*/

}

/*
 * izlaz iz aplikacije i
 * povratak u glavni meni
 *
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


	/*
	 * Pokusaj da napravim da se listaju slike nivoa
	 * klikom na strelice: <-- i -->,
	 * ali sam obrisala zbog toga sto ne registruje klik
	 * kao jedan klik nego kao 2, 3, 5... klika
	 *
	 */

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

	// previous
	// (505, 420), newV2(630, 500)

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
