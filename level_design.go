package main

import (
	"image"
	//"fmt"
	//"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
)

/*

TODO

1. promeni brojeve u %
2. dodaj jos nivoa	-->	dodato
3. dodaj meni na pocetku	-->	dodato, ali tek treba dodati jos opcija...
4. dodaj pocetak i kraj u svakom nivou
5. promeni raspored kvadrata i trouglova u svakom nivou

(ovo sam sebi napisala, nemoj ti da radis - hvala)

*/

var level_num = 0	// menu

var all_levels = false
var is_main_menu = false

var backgroundImage uninteractableImage

var triangleGraphic image.Image
var boxGraphic image.Image
var special_box_graphic image.Image


var levelIsInstantiating = true

var triangle_str = "images/main_menu/test.png"			// nigde se ne koristi
var box_str = "images/main_menu/play.png"
var backgroung_str = "images/main_menu/menu_background.png"

// textures
// var texture_array


/*
	ispreturano je sta je prvi nivo, sta meni...
	na primer set_first_level() je ustvari meni
*/

func levels_button_is_sellected() {
	if is_main_menu {	// ako nema ovog uslova --> play i levels dugme se pojavljuju svuda (na svakom nivou)
		box_str = "images/main_menu/levels_light.png"
		prefetchGraphics()
		lvl.addBox(newSpecialBox(newV2(500, 300), newV2(750, 370)))
	}
}

func play_button_is_sellected() {
	if is_main_menu {
		box_str = "images/main_menu/play_light.png"
		prefetchGraphics()
		lvl.addBox(newSpecialBox(newV2(500, 200), newV2(750, 270)))
	}
}

func clear_levels_button() {
	if is_main_menu {
		box_str = "images/main_menu/levels.png"
		prefetchGraphics()
		lvl.addBox(newSpecialBox(newV2(500, 300), newV2(750, 370)))
	}
}

func clear_play_button() {
	if is_main_menu {
		box_str = "images/main_menu/play.png"
		prefetchGraphics()
		lvl.addBox(newSpecialBox(newV2(500, 200), newV2(750, 270)))
	}
}

func set_main_menu() {
	lvl.Instantiate("lala")
	all_levels = false
	is_main_menu = true		// nepotrebno, moze preko level_num
	level_num = 0

	box_str = "images/main_menu/play.png"
	backgroung_str = "images/main_menu/menu_background.png"
	prefetchGraphics()

	// TODO Zasto mora lvl.addBox()? Ne radi bez toga!
	lvl.addBox(newSpecialBox(newV2(500, 200), newV2(750, 270)))
	clear_levels_button()
}

func set_all_levels() {

	lvl.Instantiate("lala")
	all_levels = true
	is_main_menu = false
	level_num = 6	// jer ima 5 nivoa, a meni je 0-ti

	backgroung_str = "images/main_menu/all_levels.png"
	prefetchGraphics()
}


func set_first_level() {
        lvl.Instantiate("lalaa")
	all_levels = false
	is_main_menu = false
	level_num = 1

	//triangle_str = "images/level_1/snow_triangle.png"
	box_str = "images/level_1/ice.png"
	backgroung_str = "images/level_1/first_level_background.png"
	prefetchGraphics()
                
    	// levo
        //lvl.addTriangle(newTriangle(vector2{400,500}, vector2{700,700}, "top-right"))
        //lvl.addTriangle(newTriangle(vector2{200,400}, vector2{700,700}, "top-right"))
                
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))                
        lvl.addBox(newBox(newV2(300, 500), newV2(400, 600)))
        lvl.addBox(newBox(newV2(0, 400), newV2(100, 500)))
        lvl.addBox(newBox(newV2(100, 400), newV2(200, 500)))
                
        // desno
	//triangle_str = "images/level_1/snow_triangle_rotate.png"
	//prefetchGraphics()
        //lvl.addTriangle(newTriangle(vector2{900,500}, vector2{1200,700}, "top-left"))
        //lvl.addTriangle(newTriangle(vector2{700,600}, vector2{1200,900}, "top-left"))
                
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
        lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))
        lvl.addBox(newBox(newV2(1000, 500), newV2(1100, 600)))
        lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))
        lvl.addBox(newBox(newV2(1000, 400), newV2(1100, 500)))
        lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))

	// next level
	box_str = "images/level_3/next.png"
	prefetchGraphics()
	lvl.addBox(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

}

func set_second_level() {

        lvl.Instantiate("lalaa")
	all_levels = false
	is_main_menu = false
	level_num = 2

	backgroung_str = "images/level_2/second_level_background.png"

	box_str = "images/level_2/rock.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 500), newV2(600, 600)))
        lvl.addBox(newBox(newV2(600, 500), newV2(700, 600)))
        lvl.addBox(newBox(newV2(600, 400), newV2(700, 500)))

        lvl.addBox(newBox(newV2(200, 400), newV2(300, 500)))



	box_str = "images/level_2/wood1.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))

	lvl.addBox(newBox(newV2(700, 500), newV2(800, 600)))
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
	lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))

	lvl.addBox(newBox(newV2(1100, 300), newV2(1200, 400)))
	lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
	lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))

	lvl.addBox(newBox(newV2(800, 300), newV2(900, 400)))
	lvl.addBox(newBox(newV2(750, 400), newV2(850, 500)))
	lvl.addBox(newBox(newV2(850, 400), newV2(950, 500)))

	// next level
	box_str = "images/level_3/next.png"
	prefetchGraphics()
	lvl.addBox(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

}

func set_third_level() {
    
        lvl.Instantiate("lala")
	all_levels = false
	is_main_menu = false
	level_num = 3

	backgroung_str = "images/level_3/third_level_background.png"

	box_str = "images/level_3/city_bucket_1.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(400, 500), newV2(600, 600)))
        lvl.addBox(newBox(newV2(400, 400), newV2(600, 500)))
        lvl.addBox(newBox(newV2(400, 300), newV2(600, 400)))

	box_str = "images/level_3/city_bucket_4.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(700, 500), newV2(800, 600)))
        lvl.addBox(newBox(newV2(700, 400), newV2(800, 500)))
        //lvl.addBox(newBox(newV2(700, 300), newV2(800, 400)))

        lvl.addBox(newBox(newV2(0, 400), newV2(100, 500)))



	box_str = "images/level_3/city_box_3.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))

	box_str = "images/level_3/city_box_2.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(400, 50), newV2(500, 150)))
        lvl.addBox(newBox(newV2(500, 50), newV2(600, 150)))

	// goal
	box_str = "images/level_3/flag.png"
	prefetchGraphics()
	lvl.addBox(newSpecialBox(newV2(1100, 460), newV2(1200, 560)))
	// ----------------------------------------------------------
	box_str = "images/level_3/diamond.png"
	prefetchGraphics()
	// TODO
	//tmp.Graphic.Fill(color.White)
	lvl.addBox(newSpecialBox(newV2(1000, 580), newV2(1010, 600)))

	// next level
	box_str = "images/level_3/next.png"
	prefetchGraphics()
	lvl.addBox(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))
}


func set_forth_level() {
    
        lvl.Instantiate("lala")
	all_levels = false
	is_main_menu = false
	level_num = 4

	backgroung_str = "images/level_4/forth_level_background.png"

	box_str = "images/level_4/moon_box_1.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 500), newV2(600, 600)))
        lvl.addBox(newBox(newV2(600, 500), newV2(700, 600)))
        lvl.addBox(newBox(newV2(600, 400), newV2(700, 500)))

        lvl.addBox(newBox(newV2(200, 400), newV2(300, 500)))



	box_str = "images/level_4/moon_box_1.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))

	lvl.addBox(newBox(newV2(700, 500), newV2(800, 600)))
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
	lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))

	lvl.addBox(newBox(newV2(1100, 300), newV2(1200, 400)))
	lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
	lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))

	lvl.addBox(newBox(newV2(800, 300), newV2(900, 400)))
	lvl.addBox(newBox(newV2(750, 400), newV2(850, 500)))
	lvl.addBox(newBox(newV2(850, 400), newV2(950, 500)))

	// next level
	box_str = "images/level_3/next.png"
	prefetchGraphics()
	lvl.addBox(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

}


func set_fifth_level() {

        lvl.Instantiate("lalaa")
	all_levels = false
	is_main_menu = false
	level_num = 5

	backgroung_str = "images/level_5/fifth_level_background.png"

	box_str = "images/level_5/desert_box_1.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 500), newV2(600, 600)))
        lvl.addBox(newBox(newV2(600, 500), newV2(700, 600)))
        lvl.addBox(newBox(newV2(600, 400), newV2(700, 500)))

        lvl.addBox(newBox(newV2(200, 400), newV2(300, 500)))



	box_str = "images/level_5/desert_box_2.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))

	lvl.addBox(newBox(newV2(700, 500), newV2(800, 600)))
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
	lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))

	lvl.addBox(newBox(newV2(1100, 300), newV2(1200, 400)))
	lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
	lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))

	lvl.addBox(newBox(newV2(800, 300), newV2(900, 400)))
	lvl.addBox(newBox(newV2(750, 400), newV2(850, 500)))
	lvl.addBox(newBox(newV2(850, 400), newV2(950, 500)))

	// next level
	box_str = "images/level_3/next.png"
	prefetchGraphics()
	lvl.addBox(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

}
