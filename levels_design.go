package main

import (
	//"image"
	//"fmt"
	//"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
)



func set_first_level() {
        lvl.Instantiate("lalaa")
	all_levels = false
	is_main_menu = false
	level_num = 1


	//triangle_str = "images/level_1/snow_triangle.png"
	box_str = "images/level_1/ice.png"
	backgroung_str = "images/level_1/first_level_background.png"
	triangle_str = "images/level_1/snow_triangle.png"
	prefetchGraphics()


	lvl.addTriangle(newTriangle(vector2{screenWidth/2,screenHeight/3},
		vector2{screenWidth/2+200,screenHeight/3+200},
		"bottom-left"))
                
    	// levo
        //lvl.addTriangle(newTriangle(vector2{400,500}, vector2{700,700}, "top-right"))
        //lvl.addTriangle(newTriangle(vector2{200,400}, vector2{700,700}, "top-right"))
                
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(screenWidth / 12, 5 * screenHeight / 6), newV2(200, 600)))
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
	box_str = "images/main_menu/next.png"
	prefetchGraphics()
	lvl.add_uninteractable_image(newSpecialBox(newV2(1100, 10), newV2(1120, 110)))

	lvl.hole = newHole(newV2(500, 200), newV2(600, 300))

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

        lvl.addBox(newBox(newV2(0, 400), newV2(100, 500)))


	box_str = "images/level_2/bottom_pipe.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(350, 200), newV2(450, 300)))

	box_str = "images/level_2/top_pipe.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(350, 100), newV2(450, 200)))

	box_str = "images/level_2/top_pipe.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(350, 0), newV2(450, 100)))


	box_str = "images/level_2/bottom_pipe.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 100), newV2(600, 200)))

	box_str = "images/level_2/top_pipe.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 0), newV2(600, 100)))


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
	box_str = "images/main_menu/next.png"
	prefetchGraphics()
	lvl.add_uninteractable_image(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

	lvl.hole = newHole(newV2(1030, 590), newV2(1070, 600))
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
        lvl.addBox(newBox(newV2(400, 200), newV2(600, 300)))

	box_str = "images/level_3/city_bucket_4.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(700, 500), newV2(800, 600)))
        lvl.addBox(newBox(newV2(700, 400), newV2(800, 500)))


	box_str = "images/level_3/city_box_3.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(0, 300), newV2(100, 400)))
        lvl.addBox(newBox(newV2(100, 300), newV2(200, 400)))
        lvl.addBox(newBox(newV2(200, 300), newV2(300, 400)))

	box_str = "images/level_3/city_box_2.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(800, 50), newV2(900, 150)))
        lvl.addBox(newBox(newV2(900, 50), newV2(1000, 150)))

        lvl.addBox(newBox(newV2(900, 150), newV2(1000, 250)))
        lvl.addBox(newBox(newV2(700, 0), newV2(800, 100)))
        lvl.addBox(newBox(newV2(1000, 100), newV2(1100, 200)))

	// next level
	box_str = "images/main_menu/next.png"
	prefetchGraphics()
	lvl.add_uninteractable_image(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

	lvl.hole = newHole(newV2(1030, 590), newV2(1070, 600))
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
	box_str = "images/main_menu/next.png"
	prefetchGraphics()
	lvl.add_uninteractable_image(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

	lvl.hole = newHole(newV2(1030, 590), newV2(1070, 600))
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

	lvl.addBox(newBox(newV2(950, 450), newV2(1050, 550)))

	box_str = "images/level_5/desert_box_1_rotate.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(200, 0), newV2(300, 100)))
	lvl.addBox(newBox(newV2(200, 100), newV2(300, 200)))
	lvl.addBox(newBox(newV2(200, 200), newV2(300, 300)))
        lvl.addBox(newBox(newV2(200, 400), newV2(300, 500)))


	box_str = "images/level_5/desert_box_2.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))

	lvl.addBox(newBox(newV2(700, 500), newV2(800, 600)))

	lvl.addBox(newBox(newV2(1100, 300), newV2(1200, 400)))
	lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
	lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))

	lvl.addBox(newBox(newV2(800, 300), newV2(900, 400)))
	lvl.addBox(newBox(newV2(750, 400), newV2(850, 500)))

	lvl.addBox(newBox(newV2(900, 0), newV2(1000, 100)))
	lvl.addBox(newBox(newV2(1000, 100), newV2(1100, 200)))

	lvl.addBox(newBox(newV2(500, 0), newV2(600, 100)))
	lvl.addBox(newBox(newV2(600, 100), newV2(700, 200)))

	// next level
	box_str = "images/main_menu/next.png"
	prefetchGraphics()
	lvl.add_uninteractable_image(newSpecialBox(newV2(1100, 10), newV2(1200, 110)))

	lvl.hole = newHole(newV2(980, 450), newV2(1120, 500))
}

