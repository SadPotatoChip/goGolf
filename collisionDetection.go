package main

import (
	"fmt"
)

type boxCollider struct {
	min, max vector2
}

func (c boxCollider) isCollidingWithPoint(position vector2) string {
	if position.x+player.size >= c.min.x && position.x <= c.max.x && position.y+player.size >= c.min.y && position.y <= c.max.y {
		fmt.Print("collision!")
		if player.position.y <= c.min.y {
			//adjust for collision leaway
			player.position.y = c.min.y - player.size
			player.opts.GeoM.Reset()
			player.opts.GeoM.Translate(player.position.x, player.position.y)
			return "up"
		}
		if player.position.y+player.size >= c.max.y {
			player.position.y = c.max.y + player.size
			player.opts.GeoM.Reset()
			player.opts.GeoM.Translate(player.position.x, player.position.y)
			return "down"
		}
	}

	//ADD others
	return ""
}

func (b *ball) checkForBallCollisions() string {
	for i := 0; i < lvl.nOfBoxes; i++ {
		var dir = ""
		dir = lvl.boxes[i].collider.isCollidingWithPoint(newV2(b.position.x, b.position.y))
		if dir != "" {
			return dir
		}
	}
	return ""
}
