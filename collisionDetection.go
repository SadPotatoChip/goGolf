package main

import (
	//"image/color"
	"image/color"
)

type boxCollider struct {
	Min, Max, Mid vector2
}


type triangleCollider struct {
    Min, Max, Mid vector2
    Missing_part string // moze da bude left, right...
}

func (c triangleCollider) isTriangleCollidingWithBall(b *ball) string{
    return ""
}

func (c boxCollider) isBoxCollidingWithBall(b *ball) string {
	s := ""

	if b.position.X+b.size > c.Min.X && b.position.X < c.Max.X{
		if b.verticalSpeed < 0 {
			if b.position.Y < c.Min.Y && b.collisonGhost.position.Y+b.size > c.Min.Y {
				s = "up"
			}
		} else {
			if b.position.Y+b.size > c.Max.Y && b.collisonGhost.position.Y < c.Max.Y {
				s = "down"
			}
		}
	}
	if b.position.Y+b.size> c.Min.Y && b.position.Y < c.Max.Y {
		if b.horisonatalSpeed>0{
			if b.position.X< c.Min.X && b.collisonGhost.position.X+b.size>c.Min.X{
				s="left"
			}
		}else{
			if b.position.X+ b.size > c.Min.X && b.collisonGhost.position.X < c.Max.X{
				s="right"
			}
		}
	}


	//Adjust for intersection
	switch s {
		case "up":
			b.position.Y = c.Min.Y - player.size
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.X, player.position.Y)
		case "down":
			b.position.Y = c.Max.Y
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.X, player.position.Y)
		case "right":
			b.position.X = c.Max.X
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.X, player.position.Y)
		case "left":
			b.position.X = c.Min.X - player.size
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.X, player.position.Y)
			}




	return s
}

///check all possible collisions
func (b *ball) checkForBallCollisions() string {
	candidates := getCandidateCollidersHorizontal(b)
	candidates = filterVertcial(b,candidates)

	//shows in red the boxes that are being checked for collision (replaces triangles with filled squares)
	//debugCollisionFilter(candidates)

	s:=""

	for _, boxy := range candidates {
                tmp := *boxy
                switch tmp.(type) {
                    case *triangle: s+=tmp.(*triangle).Collider.isTriangleCollidingWithBall(b)
                    case *box: s+=tmp.(*box).Collider.isBoxCollidingWithBall(b)
                }
	}
	return s
}

func getCandidateCollidersHorizontal(b *ball) []*shape {
	var collisionCandidateStartIndex int
	candidates := make([]*shape, 0)
	if b.horisonatalSpeed>=0 {
		for i := 0; i < lvl.NumOfShapes; i++ {
			tmp := lvl.MaxSortedShapes[i]
			if b.position.X < tmp.getMax().X {
				collisionCandidateStartIndex = i
				break
			}
		}
		for i := collisionCandidateStartIndex; i < lvl.NumOfShapes; i++ {
			tmp := lvl.MaxSortedShapes[i]
			candidates = append(candidates, &tmp)
		}
	}else{
		for i := 0; i < lvl.NumOfShapes; i++ {
			tmp := lvl.MinSortedShapes[i]
			if b.position.X+b.size < tmp.getMin().X {
				collisionCandidateStartIndex = i
				break
			}
		}
		for i := 0; i < collisionCandidateStartIndex; i++ {
			tmp := lvl.MinSortedShapes[i]
			candidates = append(candidates, &tmp)
		}
	}

	return candidates
}


func filterVertcial(b *ball,candidates []*shape)[]*shape {
	l:= len(candidates)
	if b.verticalSpeed>0{
		for i:=0;i< l;i++ {
			tmp := *(candidates[i])
			var tmpMin =tmp.getMin()
			if b.position.Y+b.size < tmpMin.Y {
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
			if b.position.Y > tmpMax.Y {
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
	for i := 0; i < lvl.NumOfShapes; i++ {
		tmp:=lvl.MaxSortedShapes[i]
		tmp.getGraphic().Fill(color.White)
	}
	for i := 0; i < len(candidates); i++ {
		tmp:=*candidates[i]
		tmp.getGraphic().Fill(color.RGBA{255, 0, 0, 255})
	}
}