/*
 * u ovom fajlu su funkcije koje kreiraju
 * stranice koje se prikazuju u zavisnosti
 * koja je opcija iz menija odabrana
 */

package main

import (
	"image"
)

func set_first_page() {
	lvl.Instantiate(100,100)
	is_main_menu = false

	backgroung_str = "images/main_menu/controls_blue.png"
	prefetchGraphics()

		box_str = "images/main_menu/grid.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))

}

func set_second_page() {
	lvl.Instantiate(100,100)

	is_main_menu = false

	backgroung_str = "images/main_menu/purple.png"
	prefetchGraphics()

		box_str = "images/main_menu/grid.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))

		box_str = "images/balls/nivoi.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(50, 200), newV2(1150, 400)))

		//first_photo()
		//previous_photo()
		//next_photo()
}

func set_third_page() {
	lvl.Instantiate(100,100)
	is_main_menu = false

	backgroung_str = "images/main_menu/purple.png"
	prefetchGraphics()

		box_str = "images/main_menu/grid.png"
		box_str = "images/balls/2.png"
		prefetchGraphics()
		lvl.add_uninteractable_image(newSpecialBox(newV2(1150, 20), newV2(1180, 50)))

	cont = false
	lev = false

	//ball_pick()
	//full_screen()
	speaker_on()
	//speaker_off()
}

/* izmenjeno zbog problema sa klikom */

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

/* --------------- */
