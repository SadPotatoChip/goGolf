package main

/*
 * u ovom fajlu se nalaze funkcije za kreiranje
 * opcija u meniju, dugmeta play, dugmeta za meni
 */

import (
	"image"
)

/* indikatori - zbog problema sa klikom, izbaceno je da se nivoi biraju klikom na <-- -->
var no_more_levels_next = false
var no_more_levels_back = false
*/

var mute = false

/* brojaci */
var page_num = 0
var photo_num = 1
var level_num = 0

/* indikatori gde se nalazimo*/
var all_levels = false
var is_main_menu = false

var backgroundImage uninteractableImage

var triangleGraphic image.Image
var boxGraphic image.Image
var special_box_graphic image.Image
var holeGraphic image.Image


var levelIsInstantiating = true
var menuOpen=false

var hole_str = "images/level_5/goal.png"
var triangle_str = "images/main_menu/test.png"
var box_str = "images/main_menu/play.png"
var backgroung_str = "images/main_menu/menu_background.png"

func create_play_button() {
	if is_main_menu {
		box_str = "images/main_menu/play.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(565, 190), newV2(650, 250)))
	}
}

func create_menu_button() {
	if is_main_menu {
		box_str = "images/main_menu/menu.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1140, 10), newV2(1200, 100)))
	}
}

func create_first_menu_button() {
	if is_main_menu && is_menu_oppened{// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		if is_1_menu_button_sellected {
			box_str = "images/buttons/buttons_selected/controls.png"
		} else {
			box_str = "images/buttons/buttons_unselected/controls.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 60), newV2(1180, 120)))
	}
}

func create_second_menu_button() {
	if is_main_menu && is_menu_oppened{
		if is_menu_button_2_sellected == false {
			box_str = "images/buttons/buttons_unselected/levels.png"
		} else {
			box_str = "images/buttons/buttons_selected/levels.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 120), newV2(1180, 180)))
	}
}

func create_third_menu_button() {
	if is_main_menu && is_menu_oppened{
		if is_menu_button_3_sellected == false && is_menu_oppened{
			box_str = "images/buttons/buttons_unselected/quit.png"
		} else {
			box_str = "images/buttons/buttons_selected/quit.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 180), newV2(1180, 240)))
	}
}


func speaker_on() {
	if is_main_menu == false {
		mute = false
		box_str = "images/balls/sound.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(650, 100), newV2(1050, 500)))
	}
}

func speaker_off() {
	if is_main_menu == false {
		mute = true
		box_str = "images/main_menu/speaker_off.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(570, 200), newV2(600, 400)))
	}
}

func full_screen() {
	if is_main_menu == false {
		mute = true
		box_str = "images/main_menu/full-screen.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(770, 200), newV2(800, 400)))
	}
}


func set_main_menu() {
	lvl.Instantiate(100,100)
	all_levels = false
	is_main_menu = true
	level_num = 0
	songPaths[0]="music_1.mp3"
	if selectedSong==0{
		audioPlayMainTrack()
	}

	backgroung_str = "images/main_menu/menu_background.png"
	prefetchGraphics()

	create_menu_button()
	create_play_button()
}

