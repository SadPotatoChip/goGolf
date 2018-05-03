package main

import (
	"image"
)

/*

TODO

1. promeni brojeve u %
2. dodaj jos nivoa
3. dodaj meni na pocetku

(ovo sam sebi napisala, nemoj ti da radis - hvala)

*/

var backgroundImage uninteractableImage

var triangleGraphic image.Image
var boxGraphic image.Image

var current_level = 0
var levelIsInstantiating = true

var triangle_str = "images/triangle.png"
var box_str = "images/box.png"
var backgroung_str = "images/first_level_background.png"

// textures
// var texture_array

func set_first_level() {
                
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
        lvl.addTriangle(newTriangle(vector2{900,500}, vector2{1200,700}, "top-left"))
        lvl.addTriangle(newTriangle(vector2{700,600}, vector2{1200,900}, "top-left"))
                
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
        lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))
        lvl.addBox(newBox(newV2(1000, 500), newV2(1100, 600)))
        lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))
        lvl.addBox(newBox(newV2(1000, 400), newV2(1100, 500)))
        lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
}

func set_second_level() {
    
        lvl.Instantiate("lala")

	backgroung_str = "images/second_level_background.png"

	box_str = "images/rock.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(500, 500), newV2(600, 600)))
        lvl.addBox(newBox(newV2(600, 500), newV2(700, 600)))
        lvl.addBox(newBox(newV2(600, 400), newV2(700, 500)))

	box_str = "images/wood1.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(0, 500), newV2(100, 600)))
        lvl.addBox(newBox(newV2(100, 500), newV2(200, 600)))
        lvl.addBox(newBox(newV2(200, 500), newV2(300, 600)))

	box_str = "images/rock.png"
	prefetchGraphics()
        lvl.addBox(newBox(newV2(200, 400), newV2(300, 500)))

	box_str = "images/wood1.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(700, 500), newV2(800, 600)))
        lvl.addBox(newBox(newV2(800, 500), newV2(900, 600)))
	lvl.addBox(newBox(newV2(900, 500), newV2(1000, 600)))
	//lvl.addBox(newBox(newV2(1000, 500), newV2(1100, 600)))

	box_str = "images/wood1.png"
	prefetchGraphics()
	lvl.addBox(newBox(newV2(1100, 300), newV2(1200, 400)))
	lvl.addBox(newBox(newV2(1100, 400), newV2(1200, 500)))
	lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))

	lvl.addBox(newBox(newV2(800, 300), newV2(900, 400)))
	lvl.addBox(newBox(newV2(750, 400), newV2(850, 500)))
	lvl.addBox(newBox(newV2(850, 400), newV2(950, 500)))

}

func set_third_level() {

        lvl.Instantiate("lalaa")

        lvl.addBox(newBox(newV2(1000, 500), newV2(1100, 600)))
        lvl.addBox(newBox(newV2(1100, 500), newV2(1200, 600)))
}

