/*
 * u ovom fajlu su funkcije za kreiranje nivoa
 * (5 nivoa)
 *
*/

package main

func back_to_mm() {
	box_str = "images/main_menu/next.png"
	prefetchGraphics()
	lvl.add_uninteractable_image(newSpecialBox(newV2(1100, 10), newV2(1120, 110)))
}

func set_first_level() {
	var w, h float64 = float64(screenWidth), float64(screenHeight)
	level_num = 1
	lvl.Instantiate(1*w/12.0, 1*h/6-ballSize)
	all_levels = false
	is_main_menu = false
	songPaths[0] = "music_2.mp3"
	if selectedSong == 0 {
		audioPlayMainTrack()
	}

	gravityStrenght = 0.1
	airFrictionStrenght = 0.002
	groundFrictionStrenght = 0.03

	box_str = "images/level_1/ice.png"
	backgroung_str = "images/level_1/first_level_background.png"
	triangle_str = "images/level_1/triangle.png"
	prefetchGraphics()

	for i := 0; i < 11; i++ {
		lvl.addBox(newBox(vector2{float64(i) * w / 12.0, 1 * h / 6},
			vector2{float64(i+1) * w / 12, 2 * h / 6}))
	}
	lvl.addTriangle(newTriangle(vector2{11 * w / 12.0, 3 * h / 6},
		vector2{w, 4 * h / 6}, "top-left"))
	for i := 2; i < 12; i++ {
		if i != 7 {
			lvl.addBox(newBox(vector2{float64(i) * w / 12.0, 4 * h / 6},
				vector2{float64(i+1) * w / 12, 5 * h / 6}))
		}
	}
	for i := 0; i < 12; i++ {
		lvl.addBox(newBox(vector2{float64(i) * w / 12.0, 5 * h / 6},
			vector2{float64(i+1) * w / 12, 6 * h / 6}))
	}
	lvl.addTriangle(newTriangle(vector2{11 * w / 12.0, 0},
		vector2{w, 1 * w / 12}, "bottom-left"))

	lvl.hole = newHole(vector2{3 * w / 32, (5 * h / 6) - 10},
		vector2{2 * w / 12, 5 * h / 6})

	back_to_mm()

}

func set_second_level() {
	level_num = 2
	lvl.Instantiate(150, 470) // pocetna poz loptice
	all_levels = false
	is_main_menu = false
	songPaths[0] = "music_3.mp3"
	if selectedSong == 0 {
		audioPlayMainTrack()
	}

	gravityStrenght = 0.1
	airFrictionStrenght = 0.002
	groundFrictionStrenght = 0.05

	backgroung_str = "images/level_2/second_level_background.png"

	box_str = "images/level_2/rock.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 5, h / 6 * 5), newV2(w / 12 * 600, h / 6 * 600)))
	lvl.addBox(newBox(newV2(w / 12 * 600, h / 6 * 500), newV2(w / 12 *700, h / 6 * 600)))
	lvl.addBox(newBox(newV2(w / 12 * 600, h / 6 * 400), newV2(w / 12 *700, h / 6 * 500)))

	lvl.addBox(newBox(newV2(w / 12 * 0, h / 6 * 400), newV2(w / 12 *100, h / 6 * 500)))

	box_str = "images/level_2/bottom_pipe.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 *350, h / 6 * 200), newV2(w / 12 *450, h / 6 * 300)))

	box_str = "images/level_2/top_pipe.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 *350, h / 6 * 100), newV2(w / 12 *450, h / 6 * 200)))

	box_str = "images/level_2/top_pipe.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 *350, h / 6 * 0), newV2(w / 12 *450, h / 6 * 100)))

	box_str = "images/level_2/bottom_pipe.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 *500, h / 6 * 100), newV2(w / 12 *600, h / 6 * 200)))

	box_str = "images/level_2/top_pipe.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 *500, h / 6 * 0), newV2(w / 12 *600, h / 6 * 100)))

	box_str = "images/level_2/wood1.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 0, h / 6 * 5), newV2(w / 12 *1, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 1, h / 6 * 5), newV2(w / 12 *2, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 2, h / 6 * 5), newV2(w / 12 *3, h / 6 * 6)))

	lvl.addBox(newBox(newV2(w / 12 * 7, h / 6 * 5), newV2(w / 12 *8, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 8, h / 6 * 5), newV2(w / 12 *9, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 9, h / 6 * 5), newV2(w / 12 *10, h / 6 * 6)))

	lvl.addBox(newBox(newV2(w / 12 * 11, h / 6 * 3), newV2(w / 12 *1200, h / 6 * 4)))
	lvl.addBox(newBox(newV2(w / 12 * 11, h / 6 * 4), newV2(w / 12 *1200, h / 6 * 5)))
	lvl.addBox(newBox(newV2(w / 12 * 11, h / 6 * 5), newV2(w / 12 *1200, h / 6 * 6)))

	lvl.addBox(newBox(newV2(w / 12 * 8, h / 6 * 3), newV2(w / 12 *9, h / 6 * 4)))
	lvl.addBox(newBox(newV2(w / 12 * 7.5, h / 6 * 4), newV2(w / 12 *8.5, h / 6 * 5)))
	lvl.addBox(newBox(newV2(w / 12 * 8.5, h / 6 * 4), newV2(w / 12 *9.5, h / 6 * 5)))

	lvl.hole = newHole(newV2(w / 12 * 10.3, 5.9), newV2(w / 12 *10.7, 6))

	back_to_mm()
}

func set_third_level() {
	level_num = 3
	var w, h float64 = float64(screenWidth), float64(screenHeight)
	lvl.Instantiate(w/32, h/2)
	all_levels = false
	is_main_menu = false
	songPaths[0] = "music_4.mp3"
	if selectedSong == 0 {
		audioPlayMainTrack()
	}

	gravityStrenght = 0.1
	airFrictionStrenght = 0.002
	groundFrictionStrenght = 0.05

	backgroung_str = "images/level_3/third_level_background.png"

	box_str = "images/level_3/city_bucket_1.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 4, h / 6 * 5), newV2(w / 12 *5, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 4, h / 6 * 4), newV2(w / 12 *5, h / 6 * 5)))
	lvl.addBox(newBox(newV2(w / 12 *4, h / 6 * 3), newV2(w / 12 *5, h / 6 * 4)))
	lvl.addBox(newBox(newV2(w / 12 * 4, h / 6 * 2), newV2(w / 12 *5, h / 6 * 3)))

	box_str = "images/level_3/city_bucket_4.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 7, h / 6 * 5), newV2(w / 12 *8, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 7, h / 6 * 4), newV2(w / 12 *8, h / 6 * 5)))

	box_str = "images/level_3/city_box_3.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 0, h / 6 * 3.3), newV2(w / 12 *1, h / 6 * 4)))
	lvl.addBox(newBox(newV2(w / 12 * 1, h / 6 * 3.3), newV2(w / 12 *2, h / 6 * 4)))
	lvl.addBox(newBox(newV2(w / 12 * 2, h / 6 * 3.3), newV2(w / 12 *3, h / 6 * 4)))

	box_str = "images/level_3/city_box_2.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 8, h / 6 * 0.5), newV2(w / 12 *9, h / 6 * 15)))
	lvl.addBox(newBox(newV2(w / 12 * 9, h / 6 * 0.5), newV2(w / 12 *10, h / 6 * 15)))

	lvl.addBox(newBox(newV2(w / 12 * 9, h / 6 * 1.5), newV2(w / 12 *10, h / 6 * 2.5)))
	lvl.addBox(newBox(newV2(w / 12 * 7, h / 6 * 0), newV2(w / 12 *8, h / 6 * 1)))
	lvl.addBox(newBox(newV2(w / 12 * 10, h / 6 * 1), newV2(w / 12 *11, h / 6 * 2)))

	lvl.hole = newHole(newV2(w / 12 * 10.3, h / 6 * 5.9), newV2(w / 12 *10.7, h / 6 * 6))

	back_to_mm()
}

func set_forth_level() {
	var w, h float64 = float64(screenWidth), float64(screenHeight)
	level_num = 4
	lvl.Instantiate(w/32, h-ballSize)
	all_levels = false
	is_main_menu = false
	songPaths[0] = "music_5.mp3"
	if selectedSong == 0 {
		audioPlayMainTrack()
	}

	gravityStrenght = 0.04
	airFrictionStrenght = 0.002
	groundFrictionStrenght = 0.05

	backgroung_str = "images/level_4/forth_level_background.png"
	box_str = "images/level_4/moon_box_1.png"
	triangle_str = "images/level_4/triangle.png"
	prefetchGraphics()
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			if !((i == 4 && j == 1) || (i == 2 && j == 1) || (i == 4 && j == 2) || (i == 3 && j == 0)) {
				lvl.addBox(newBox(vector2{float64(i+1) * w / 12.0, (5 - float64(j)) * h / 6},
					vector2{float64(i+2) * w / 12.0, (5 - float64(j-1)) * h / 6}))
			}
		}
		lvl.addTriangle(newTriangle(vector2{float64(i+1) * w / 12.0, (5 - float64(i)) * h / 6},
			vector2{float64(i+2) * w / 12.0, (6 - float64(i)) * h / 6}, "top-left"))
	}

	lvl.addBox(newBox(vector2{13 * w / 24.0, 7 * h / 12},
		vector2{15 * w / 24.0, 9 * h / 12}))
	lvl.addBox(newBox(vector2{17 * w / 24.0, 11 * h / 12},
		vector2{19 * w / 24.0, 13 * h / 12}))
	lvl.addBox(newBox(vector2{21 * w / 24.0, 9 * h / 12},
		vector2{23 * w / 24.0, 11 * h / 12}))
	lvl.addBox(newBox(vector2{18.5 * w / 24.0, 3.5 * h / 12},
		vector2{20.5 * w / 24.0, 5.5 * h / 12}))

	lvl.hole = newHole(vector2{5.3 * w / 12.0, (5*h - 15) / 6},
		vector2{5.8 * w / 12.0, 6 * h / 6})

	back_to_mm()

}

func set_fifth_level() {
	level_num = 5
	lvl.Instantiate(50, 470)
	all_levels = false
	is_main_menu = false

	songPaths[0] = "music_6.mp3"
	if selectedSong == 0 {
		audioPlayMainTrack()
	}

	gravityStrenght = 0.1
	airFrictionStrenght = 0.002
	groundFrictionStrenght = 0.05

	backgroung_str = "images/level_5/fifth_level_background.png"

	box_str = "images/level_5/desert_box_1.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 5, h / 6 * 5), newV2(w / 12 *6, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 6, h / 6 * 5), newV2(w / 12 *7, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 6, h / 6 * 4), newV2(w / 12 *7, h / 6 * 5)))

	lvl.addBox(newBox(newV2(w / 12 * 950, 450), newV2(1050, h / 6 * 550)))

	box_str = "images/level_5/desert_box_1_rotate.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 2, h / 6 * 0), newV2(w / 12 *3, h / 6 * 1)))
	lvl.addBox(newBox(newV2(w / 12 * 2, h / 6 * 1), newV2(w / 12 *3, h / 6 * 2)))
	lvl.addBox(newBox(newV2(w / 12 * 2, h / 6 * 2), newV2(w / 12 *3, h / 6 * 3)))
	lvl.addBox(newBox(newV2(w / 12 * 2, h / 6 * 4), newV2(w / 12 *3, h / 6 * 5)))

	box_str = "images/level_5/desert_box_2.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(w / 12 * 0, h / 6 * 5), newV2(w / 12 *1, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 1, h / 6 * 5), newV2(w / 12 *2, h / 6 * 6)))
	lvl.addBox(newBox(newV2(w / 12 * 2, h / 6 * 5), newV2(w / 12 *3, h / 6 * 6)))

	lvl.addBox(newBox(newV2(w / 12 * 7, h / 6 * 5), newV2(w / 12 *8, h / 6 * 6)))

	lvl.addBox(newBox(newV2(w / 12 * 11, h / 6 * 3), newV2(w / 12 *12, h / 6 * 4)))
	lvl.addBox(newBox(newV2(w / 12 * 11, h / 6 * 4), newV2(w / 12 *12, h / 6 * 5)))
	lvl.addBox(newBox(newV2(w / 12 * 11, h / 6 * 5), newV2(w / 12 *12, h / 6 * 6)))

	lvl.addBox(newBox(newV2(w / 12 * 8, h / 6 * 3), newV2(w / 12 *9, h / 6 * 4)))
	lvl.addBox(newBox(newV2(w / 12 * 7.5, h / 6 * 4), newV2(w / 12 *8.5, h / 6 * 5))) // TODO

	lvl.addBox(newBox(newV2(w / 12 * 9, h / 6 * 0), newV2(w / 12 *10, h / 6 * 1)))
	lvl.addBox(newBox(newV2(w / 12 * 10, h / 6 * 1), newV2(w / 12 *11, h / 6 * 2)))

	lvl.addBox(newBox(newV2(w / 12 * 5, h / 6 * 0), newV2(w / 12 *6, h / 6 * 1)))
	lvl.addBox(newBox(newV2(w / 12 * 6, h / 6 * 1), newV2(w / 12 *7, h / 6 * 2)))

	lvl.hole = newHole(newV2(w / 12 * 9.8, h / 6 * 4.45), newV2(w / 12 *10.3, h / 6 * 5)) // TODO

	back_to_mm()
}
