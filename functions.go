/*
 * u ovom fajlu su pomocne funkcije
 * koje na osnovu polozaja kursora
 * odredjuju da li je neki objekat
 * selektovan ili je na njega kliknuto
 */

package main

import (
	"github.com/hajimehoshi/ebiten"
)

/*
 * koordinate centa i duzina poluprecnika play dugmeta
 */
var x_center_coord = 610
var y_center_coord = 230
var radius_coord = 110

/*
 * PLAY button
 *
 */
func play_button_is_clicked() {
	if ((x_pos-x_center_coord)*(x_pos-x_center_coord) + (y_pos-y_center_coord)*(y_pos-y_center_coord)) <= (radius_coord * radius_coord){
		if x_pos > 500 && x_pos < 720 && y_pos > 200 && y_pos < 260 && is_main_menu {
			if mouseButtonDown(ebiten.MouseButtonLeft) {
				set_first_level()
			}
		}
	}
}

/*
 * MENU button
 *
 */
func is_menu_button_selected() bool {
	if x_pos > 1140 && x_pos < 1170 && y_pos > 10 && y_pos < 50 && is_main_menu {
		return true
	} else {
		return false
	}
}

/*
 * ako je selektovana prva opcija menija vraca true
 *
 */
func is_first_option_selected() bool {
	if x_pos > 1000 && x_pos < 1190 && y_pos > 60 && y_pos < 120 && is_main_menu && is_menu_oppened && !(page_num == 1) {
		return true;
	} else {
		return false;
	}
}

/*
 * nije mogla da se iskoristi prethodna funkcija za ispitivanje
 * da li nije selektovana prva opcija menija, zato pisem novu
 * (isto vazi za osale opcije)
 *
 */
func is_not_first_option_selected() bool {
	if main_menu_already_set == false && !(x_pos > 1000 && x_pos < 1190 && y_pos > 60 && y_pos < 300) && is_main_menu == true {
		return true
	} else {
		return false
	}
}

/*
 * ako je selektovana druga opcija menija vraca true
 *
 */
func is_second_option_selected() bool {
	if x_pos > 1000 && x_pos < 1190 && y_pos > 120 && y_pos < 180 && is_main_menu && is_menu_oppened {
		return true
	} else {
		return false
	}
}

func is_not_second_option_selected() bool {
	if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 120 && y_pos < 180 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu) {
		return true
	} else {
		return false
	}
}

/*
 * ako je selektovana treca opcija menija vraca true
 *
 */
func is_third_option_selected() bool {
	if x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu && is_menu_oppened {
		return true
	} else {
		return false
	}
}

func is_not_third_option_selected() bool {
	if is_menu_oppened && !(x_pos > 1000 && x_pos < 1190 && y_pos > 180 && y_pos < 240 && is_main_menu && is_menu_oppened) && (x_pos > 1000 && x_pos < 1190 && y_pos > 10 && y_pos < 300 && is_main_menu) {
		return true
	} else {
		return false
	}
}

/*
 * ako je kliknuto na grid, vraca se u meni
 *
 */
func is_grid_clicked() bool {
	if x_pos > 1150 && x_pos < 1180 && y_pos > 20 && y_pos < 50 && (is_main_menu == false) && (lev || opt || cont) {
		return true
	} else {
		return false
	}
}
