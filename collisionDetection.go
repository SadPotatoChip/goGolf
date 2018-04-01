package main

import (
	"math"
)

type boxCollider struct {
	min, max, mid vector2
}

//TODO swap out player size and fix argument passing
func (c boxCollider) isCollidingWithBall(b *ball) string {
	s := ""
	var ledgeCandidate float64 = 0
	if b.position.x+b.size >= c.min.x && b.position.x <= c.max.x && b.position.y+b.size >= c.min.y && b.position.y <= c.max.y {
		if b.position.y <= c.mid.y {
			s = "up"
			ledgeCandidate = c.min.y
		} else {
			s = "down"
			ledgeCandidate = c.max.y
		}
		candidateDistance := math.Abs(ledgeCandidate - b.position.y)
		if candidateDistance > math.Abs(b.position.x-c.max.x) && b.horisonatalSpeed < 0 {
			s = "right"
		}
		if candidateDistance > math.Abs(b.position.x-c.min.x) && b.horisonatalSpeed > 0 {
			s = "left"
		}

		//Adjust for intersection
		switch s {
		case "up":
			player.position.y = c.min.y - player.size
			player.opts.GeoM.Reset()
			player.opts.GeoM.Translate(player.position.x, player.position.y)
		case "down":
			player.position.y = c.max.y + player.size
			player.opts.GeoM.Reset()
			player.opts.GeoM.Translate(player.position.x, player.position.y)
		case "right":
			player.position.x = c.max.x + player.size
			player.opts.GeoM.Reset()
			player.opts.GeoM.Translate(player.position.x, player.position.y)
		case "left":
			player.position.x = c.min.x - player.size
			player.opts.GeoM.Reset()
			player.opts.GeoM.Translate(player.position.x, player.position.y)

		}

	}

	//ADD others
	return s
}

func (b *ball) checkForBallCollisions() string {
	for i := 0; i < lvl.nOfBoxes; i++ {
		var dir = ""
		dir = lvl.boxes[i].collider.isCollidingWithBall(b)
		if dir != "" {
			return dir
		}
	}
	return ""
}
