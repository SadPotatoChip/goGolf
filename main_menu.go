package main

import (
	"image"
)

/*
	TODO
	1. promeni brojeve u %
	2. dodaj pocetak i kraj u svakom nivou
*/

var page_num = 0

var photo_num = 1

var level_num = 0	// menu

var all_levels = false
var is_main_menu = false

var backgroundImage uninteractableImage

var triangleGraphic image.Image
var boxGraphic image.Image
var special_box_graphic image.Image
var holeGraphic image.Image


var levelIsInstantiating = true

var hole_str = "images/level_5/goal.png"	// promeni
var triangle_str = "images/main_menu/test.png"			// nigde se ne koristi
var box_str = "images/main_menu/play.png"
var backgroung_str = "images/main_menu/menu_background.png"

func create_play_button() {
	if is_main_menu {
		box_str = "images/main_menu/play.png"
		prefetchGraphics()
		// TODO Zasto mora lvl.addBox()? Ne radi bez toga!
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

// dugme.Graphic.Fill(color.RGBA{0, 0, 0, 0})

// opcije u meniju
func create_second_menu_button() {
	if is_main_menu && is_menu_oppened{
		if is_menu_button_2_sellected == false {
			box_str = "images/buttons/buttons_unselected/levels.png"
		} else {
			box_str = "images/buttons/buttons_selected/levels.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 120), newV2(1190, 180)))
	}
}

func create_first_menu_button() {
	if is_main_menu {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		if is_1_menu_button_sellected {
			box_str = "images/buttons/buttons_selected/controls.png"
		} else {
			box_str = "images/buttons/buttons_unselected/controls.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 60), newV2(1190, 120)))
	}
}

func create_third_menu_button() {
	if is_main_menu {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		if is_menu_button_3_sellected == false && is_menu_oppened{
			box_str = "images/buttons/buttons_unselected/options.png"
		} else {
			box_str = "images/buttons/buttons_selected/options.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 180), newV2(1190, 240)))
	}
}

func create_forth_menu_button() {
	if is_main_menu {
		if is_menu_button_4_sellected == false && is_menu_oppened{
			box_str = "images/buttons/buttons_unselected/quit.png"
		} else {
			box_str = "images/buttons/buttons_selected/quit.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 240), newV2(1190, 300)))
	}
}


func set_first_page() {
	lvl.Instantiate("lala")
	all_levels = true
	is_main_menu = false
	level_num = 6	// jer ima 5 nivoa, a meni je 0-ti

	page_num = 1

	backgroung_str = "images/main_menu/purple.png"
	prefetchGraphics()

	if is_main_menu == false {
		box_str = "images/main_menu/grid.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))


		first_photo()
		previous_photo()
		next_photo()

	}
}

func set_second_page() {
	lvl.Instantiate("lala")
	all_levels = true
	is_main_menu = false
	level_num = 6	// jer ima 5 nivoa, a meni je 0-ti

	backgroung_str = "images/main_menu/controls.png"
	prefetchGraphics()

	if is_main_menu == false {
		box_str = "images/main_menu/grid.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))
	}
}

func set_third_page() {
	lvl.Instantiate("lala")
	all_levels = true
	is_main_menu = false
	level_num = 6	// jer ima 5 nivoa, a meni je 0-ti

	backgroung_str = "images/main_menu/purple.png"
	prefetchGraphics()

	if is_main_menu == false {
		box_str = "images/main_menu/grid.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))
	}
}

func first_photo() {
	photo_num = 1
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/main_menu/1.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(500, 100), newV2(600, 200)))
	}
}

func second_photo() {
	photo_num = 2
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/main_menu/2.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(500, 100), newV2(600, 200)))
	}
}


func third_photo() {
	photo_num = 3
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/main_menu/3.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(500, 100), newV2(600, 200)))
	}
}

func previous_photo() {
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/main_menu/back.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(530, 400), newV2(630, 500)))
	}
}

func next_photo() {
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/main_menu/next.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(630, 400), newV2(730, 500)))
	}
}


func set_main_menu() {
	lvl.Instantiate("lala")
	all_levels = false
	is_main_menu = true		// nepotrebno, moze preko level_num
	level_num = 0

	backgroung_str = "images/main_menu/menu_background.png"
	prefetchGraphics()

	create_menu_button()
	create_play_button()
}

