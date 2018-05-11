package main

import (
	"image"
)

/*
	TODO
	1. promeni brojeve u %
	2. dodaj pocetak i kraj u svakom nivou
*/

var no_more_levels_next = false
var no_more_levels_back = false

var mute = false

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
		prefetchGraphics()					// TODO proveeeri granice 1190 je bilo
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 120), newV2(1180, 180)))
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

func create_third_menu_button() {
	if is_main_menu && is_menu_oppened{// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		if is_menu_button_3_sellected == false && is_menu_oppened{
			box_str = "images/buttons/buttons_unselected/options.png"
		} else {
			box_str = "images/buttons/buttons_selected/options.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 180), newV2(1180, 240)))
	}
}

func create_forth_menu_button() {
	if is_main_menu && is_menu_oppened{
		if is_menu_button_4_sellected == false && is_menu_oppened{
			box_str = "images/buttons/buttons_unselected/quit.png"
		} else {
			box_str = "images/buttons/buttons_selected/quit.png"
		}
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1000, 240), newV2(1180, 300)))
	}
}


func set_second_page() {
	lvl.Instantiate("lala")
	//all_levels = true
	is_main_menu = false

	//page_num = 1

	backgroung_str = "images/main_menu/purple.png"
	prefetchGraphics()

	//if is_main_menu == false {	// TODO
		box_str = "images/main_menu/grid.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))

		box_str = "images/balls/nivoi.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(50, 200), newV2(1150, 400)))

		//first_photo()
		//previous_photo()
		//next_photo()

	//}
}

func set_first_page() {
	lvl.Instantiate("lala")
	is_main_menu = false
	//if cont == true {

	backgroung_str = "images/main_menu/controls.png"
	prefetchGraphics()

	//if is_main_menu == false {
		box_str = "images/main_menu/grid.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))
	//}
	//opt = false
	//lev = false
	//}
}

func set_third_page() {
	lvl.Instantiate("lala")
	is_main_menu = false

	//if opt == true {
	backgroung_str = "images/main_menu/purple.png"
	prefetchGraphics()

	//if is_main_menu == false {
		box_str = "images/main_menu/grid.png"
		box_str = "images/balls/2.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))
		//lvl.add_uninteractable_image(newSpecialBox(newV2(150, 100), newV2(550, 500)))
	//}

	cont = false
	lev = false

	//ball_pick()
	//full_screen()
	speaker_on()
	//speaker_off()
	//}
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

func create_balls() {
	box_str = "images/balls/all_balls.png"
	prefetchGraphics()
	lvl.add_uninteractable_image(newSpecialBox(newV2(350, 100), newV2(500, 200)))


}

func ball_pick() {
	if is_main_menu == false {
		mute = true
		box_str = "images/balls/two_balls_big.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(100, 100), newV2(400, 400)))
	}
}

func first_photo() {
	photo_num = 1
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/lev_1.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(300, 100), newV2(600, 200)))
	}
}

func second_photo() {
	photo_num = 2
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/lev_2.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(300, 100), newV2(600, 200)))
	}
}


func third_photo() {
	photo_num = 3
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/lev_3.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(300, 100), newV2(600, 200)))
	}
}

func forth_photo() {
	photo_num = 4
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/lev_4.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(300, 100), newV2(600, 200)))
	}
}

func fifth_photo() {
	photo_num = 5
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/lev_5.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(300, 100), newV2(600, 200)))
	}
}

func previous_photo() {
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/back.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(505, 420), newV2(630, 500)))
	}
}

func next_photo() {
	if is_main_menu == false {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/next.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(605, 420), newV2(730, 600)))
	}
}

func change_previous_button() {
	if is_main_menu == false && no_more_levels_back == true {
	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/back_grey.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(505, 420), newV2(630, 500)))
	}
}

func change_next_button() {
	if is_main_menu == false && no_more_levels_next == true {
	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/pick_the_level/next_grey.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(605, 420), newV2(730, 600)))
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

