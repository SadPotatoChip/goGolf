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
2. dodaj jos nivoa	-->	dodato, treba jos 1
3. dodaj meni na pocetku	-->	dodato, ali tek treba dodati jos opcija...
4. dodaj pocetak i kraj u svakom nivou
5. promeni raspored kvadrata i trouglova u svakom nivou

(ovo sam sebi napisala, nemoj ti da radis - hvala)

*/

var all_levels = false
var is_main_menu = false

var backgroundImage uninteractableImage

var triangleGraphic image.Image
var boxGraphic image.Image
var special_box_graphic image.Image

var current_level = 0
var levelIsInstantiating = true

var triangle_str = "images/main_menu/test.png"
var box_str = "images/main_menu/play_button.png"
var backgroung_str = "images/main_menu/menu_background.png"

// textures
// var texture_array


/*
	ispreturano je sta je prvi nivo, sta meni...
	na primer set_first_level() je ustvari meni
*/

func set_first_level() {
	//lvl.Instantiate("lala")
	is_main_menu = true

	box_str = "images/main_menu/play_button.png"
	backgroung_str = "images/main_menu/menu_background.png"
	prefetchGraphics()

	// TODO Zasto mora lvl.addBox()? Ne radi bez toga!
	lvl.addBox(newSpecialBox(newV2(500, 200), newV2(750, 270)))
	box_str = "images/main_menu/see_all_levels.png"
	prefetchGraphics()
	lvl.addBox(newSpecialBox(newV2(500, 300), newV2(750, 370)))
}

func set_all_levels() {
	lvl.Instantiate("lala")
	all_levels = true
	is_main_menu = false

	backgroung_str = "images/main_menu/all_levels.png"
	prefetchGraphics()
}

// func IsMouseButtonPressed(mouseButton MouseButton) bool

func set_main_menu() {

	triangle_str = "images/level_1/triangle.png"
	box_str = "images/level_1/box.png"
	backgroung_str = "images/level_1/first_level_background.png"
	prefetchGraphics()
                
    	// levo
        lvl.addTriangle(newTriangle(vector2{400,500}, vector2{700,700}, "top-right"))
        lvl.addTriangle(newTriangle(vector2{200,400}, vector2{700,700}, "top-right"))
                
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))                
        lvl.addBox(newBox(newV2(300, 500), newV2(400, 600)))
        lvl.addBox(newBox(newV2(0, 400), newV2(100, 500)))
        lvl.addBox(newBox(newV2(100, 400), newV2(200, 500)))
                
        // desno
	triangle_str = "images/level_1/triangle_rotate.png"
	prefetchGraphics()
        lvl.addTriangle(newTriangle(vector2{900,500}, vector2{1200,700}, "top-left"))
        lvl.addTriangle(newTriangle(vector2{700,600}, vector2{1200,900}, "top-left"))
                
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
        lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))
        lvl.addBox(newBox(newV2(1000, 500), newV2(1100, 600)))
        lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))
        lvl.addBox(newBox(newV2(1000, 400), newV2(1100, 500)))
        lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
}

func set_forth_level() {
    
        lvl.Instantiate("lala")

	backgroung_str = "images/level_4/forth_level_background_bloor.png"

	box_str = "images/level_4/moon_box_1.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 500), newV2(600, 600)))
        lvl.addBox(newBox(newV2(600, 500), newV2(700, 600)))
        lvl.addBox(newBox(newV2(600, 400), newV2(700, 500)))

        lvl.addBox(newBox(newV2(200, 400), newV2(300, 500)))



	box_str = "images/level_4/moon_box_2.png"
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

}

func set_second_level() {

        lvl.Instantiate("lalaa")
	all_levels = false
	is_main_menu = false

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

}

func set_fifth_level() {

        lvl.Instantiate("lalaa")

	backgroung_str = "images/level_5/fifth_level_background.png"

	box_str = "images/level_5/desert_rock.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 500), newV2(600, 600)))
        lvl.addBox(newBox(newV2(600, 500), newV2(700, 600)))
        lvl.addBox(newBox(newV2(600, 400), newV2(700, 500)))

        lvl.addBox(newBox(newV2(200, 400), newV2(300, 500)))



	box_str = "images/level_5/desert_box.png"
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

}

