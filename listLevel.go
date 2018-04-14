package main

//contains obsolete code with a list implemetation of the lvl.boxes

/*
import (
	"math"
	"container/list"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)



import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"container/list"
)

const maxLevelObjects int = 50

type level struct {
	boxes    *list.List
	nOfBoxes int
}

type box struct {
	collider boxCollider
	graphic  *ebiten.Image
	opts     *ebiten.DrawImageOptions
}

func (l *level) Instantiate(  ) {
	l.nOfBoxes = 0
	l.boxes = list.New()
	l.boxes.Init()
	ground := newBox(newV2(0, screenHeight-10), newV2(screenWidth, screenHeight+90))
	l.addBox(ground)
}

func (l *level) addBox(b *box) {
	var mark *list.Element = lvl.boxes.Front()
	if lvl.nOfBoxes>1{
		for ; mark.Value.(*box).collider.min.x < b.collider.min.x && mark.Next()!=nil;
		mark = mark.Next(){
		}

		l.boxes.InsertBefore(b, mark)
	}else{
		l.boxes.PushBack(b)
	}

	//debug
	for curr:=lvl.boxes.Front(); curr!=nil; curr = curr.Next() {
		fmt.Printf("%f ",curr.Value.(*box).collider.min.x)
	}
	fmt.Printf("\n")


	l.nOfBoxes++
}






package main

import (
	"math"
	"container/list"
	"image/color"
)

type boxCollider struct {
	min, max, mid vector2
}

//HACK
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

	return s
}

///wrapper for checking all the collisions for all the boxes
func (b *ball) checkForBallCollisions() string {
	var originalLevel=*lvl.boxes
	var collisionCandidates *list.List=lvl.boxes

	if b.horisonatalSpeed>0{
		collisionCandidates=removeBoxesOnLeftSide(collisionCandidates, b)
	}else{
		collisionCandidates=removeBoxesOnRightSide(collisionCandidates, b)
	}

	*lvl.boxes=originalLevel

	for current:=lvl.boxes.Front();current!=nil;current=current.Next(){
		current.Value.(*box).graphic.Fill(color.White)
	}

	for current:=collisionCandidates.Front();current!=nil;current=current.Next(){
		current.Value.(*box).graphic.Fill(color.RGBA{255,0,0,255})
	}

	return ""
}

func removeBoxesOnLeftSide(collisionCandidates *list.List, b *ball) *list.List{
	for current:=collisionCandidates.Front();
		current.Value.(*box).collider.min.x<b.position.x && current.Next()!=nil;{
		current=current.Next()
		collisionCandidates.Remove(collisionCandidates.Front())
	}
	return collisionCandidates
}

func removeBoxesOnRightSide(collisionCandidates *list.List, b *ball) *list.List{
	var current=collisionCandidates.Front()
	for ;
		current.Value.(*box).collider.min.x<b.position.x+b.size && current.Next()!=nil;{
		current=current.Next()
	}
	for current!=nil {
		tmp:=current.Next()
		collisionCandidates.Remove(current)
		current=tmp
	}
	return collisionCandidates
}


*/







