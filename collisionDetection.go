package main

import (
	//"image/color"
	"image/color"
)

type boxCollider struct {
	min, max, mid vector2
}


type triangleCollider struct {
    min, max, mid vector2
    missing_part string // moze da bude left, right...
}

func (c triangleCollider) isTriangleCollidingWithBall(b *ball) string{
    return ""
}

func (c boxCollider) isBoxCollidingWithBall(b *ball) string {
	s := ""

	if b.position.x+b.size > c.min.x && b.position.x < c.max.x{
		if b.verticalSpeed < 0 {
			if b.position.y < c.min.y && b.collisonGhost.position.y+b.size > c.min.y {
				s = "up"
			}
		} else {
			if b.position.y+b.size > c.max.y && b.collisonGhost.position.y < c.max.y {
				s = "down"
			}
		}
	}
	if b.position.y+b.size> c.min.y && b.position.y < c.max.y {
		if b.horisonatalSpeed>0{
			if b.position.x< c.min.x && b.collisonGhost.position.x+b.size>c.min.x{
				s="left"
			}
		}else{
			if b.position.x+ b.size > c.max.x && b.collisonGhost.position.x < c.max.x{
				s="right"
			}
		}
	}


	//Adjust for intersection
	switch s {
		case "up":
			b.position.y = c.min.y - player.size
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
		case "down":
			b.position.y = c.max.y
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
		case "right":
			b.position.x = c.max.x
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
		case "left":
			b.position.x = c.min.x - player.size
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
			}




	return s
}

///check all possible collisions
func (b *ball) checkForBallCollisions() string {
	candidates := getCandidateCollidersHorizontal(b)
	candidates = filterVetcial(b,candidates)

	//shows in red the boxes that are being checked for collision (replaces triangles with filled squares)
	//debugCollisionFilter(candidates)



	s:=""

	for _, boxy := range candidates {
                tmp := *boxy
                switch tmp.(type) {
                    case *triangle: s+=tmp.(*triangle).collider.isTriangleCollidingWithBall(b)
                    case *box: s+=tmp.(*box).collider.isBoxCollidingWithBall(b)
                }
	}


	return s
}

func getCandidateCollidersHorizontal(b *ball) []*shape {
	var collisionCandidateStartIndex int
	candidates := make([]*shape, 0)
	if b.horisonatalSpeed>=0 {
		for i := 0; i < lvl.nOfShapes; i++ {
			tmp := *lvl.maxSortedShapes[i]
			if b.position.x < tmp.getMax().x {
				collisionCandidateStartIndex = i
				break
			}
		}
		for i := collisionCandidateStartIndex; i < lvl.nOfShapes; i++ {
			tmp := *lvl.maxSortedShapes[i]
			candidates = append(candidates, &tmp)
		}
	}else{
		for i := 0; i < lvl.nOfShapes; i++ {
			tmp := *lvl.minSortedShapes[i]
			if b.position.x+b.size < tmp.getMin().x {
				collisionCandidateStartIndex = i
				break
			}
		}
		for i := 0; i < collisionCandidateStartIndex; i++ {
			tmp := *lvl.minSortedShapes[i]
			candidates = append(candidates, &tmp)
		}
	}

	return candidates
}


func filterVetcial(b *ball,candidates []*shape)[]*shape {
	l:= len(candidates)
	if b.verticalSpeed>0{
		for i:=0;i< l;i++ {
			tmp := *(candidates[i])
			var tmpMin =tmp.getMin()
			if b.position.y+b.size < tmpMin.y {
				copy(candidates[i:], candidates[i+1:])
				candidates[len(candidates)-1] = nil // or the zero value of T
				candidates = candidates[:len(candidates)-1]
				l=len(candidates)
			}
		}
	}else{
		for i:=0;i < l;i++ {
			tmp := *(candidates[i])
			var tmpMax =tmp.getMax()
			if b.position.y > tmpMax.y {
				copy(candidates[i:], candidates[i+1:])
				candidates[len(candidates)-1] = nil // or the zero value of T
				candidates = candidates[:len(candidates)-1]
				l=len(candidates)
			}
		}
	}


	return candidates
}

func debugCollisionFilter(candidates []*shape){
	for i := 0; i < lvl.nOfShapes; i++ {
		tmp:=*lvl.maxSortedShapes[i]
		tmp.getGraphic().Fill(color.White)
	}
	for i := 0; i < len(candidates); i++ {
		tmp:=*candidates[i]
		tmp.getGraphic().Fill(color.RGBA{255, 0, 0, 255})
	}
}