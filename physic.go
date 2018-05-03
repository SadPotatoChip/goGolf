package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

const maxSpeed float64 = 40
const gravityStrenght float64 = 0.05
const airFrictionStrenght float64 = 0.002
const groundFrictionStrenght float64 = 0.05
const inputAngleLerp float64 = math.Pi / 6

const ballSize float64 = 5

type ball struct {
	position                        vector2
	size                            float64
	graphic                         *ebiten.Image
	opts                            *ebiten.DrawImageOptions
	controls                        *controler
	verticalSpeed, horisonatalSpeed float64
	isGrounded                      bool
	isGhost							bool
	collisonGhost					*ball
}

func makeBall(x, y float64,isGhost bool) *ball {
	tmp := new(ball)
	tmp.size = ballSize
	tmp.position.X = x
	tmp.position.Y = y
	tmp.graphic, _ = ebiten.NewImage(int(ballSize), int(ballSize), ebiten.FilterNearest)
	tmp.isGhost=isGhost
	if isGhost==false {
		tmp.graphic.Fill(color.RGBA{255,165,0, 255})
		tmp.collisonGhost = makeBall(x,y,true)
		tmp.controls = makeControler(tmp)
	}else{
		tmp.graphic.Fill(color.RGBA{0, 0, 255, 0})
	}
	tmp.opts = &ebiten.DrawImageOptions{}
	tmp.opts.GeoM.Translate(x, y)
	tmp.verticalSpeed, tmp.horisonatalSpeed = 0, 0
	tmp.isGrounded = true


	return tmp
}

func (b *ball) resetGhostPosition(){
	b.collisonGhost.position.X = b.position.X
	b.collisonGhost.position.Y = b.position.Y
	b.collisonGhost.opts=&ebiten.DrawImageOptions{}
	b.collisonGhost.opts.GeoM.Translate(b.position.X, b.position.Y)
	b.collisonGhost.isGrounded = b.isGrounded
	b.collisonGhost.move()
}

func (b *ball) applyNaturalForces() {
	if b.isGhost==false {
		player.collisonGhost.applyNaturalForces()
	}
	if b.isGrounded == false {
		if b.verticalSpeed > -maxSpeed {
			b.verticalSpeed -= gravityStrenght
		}
		if b.horisonatalSpeed > 0 {
			b.horisonatalSpeed -= airFrictionStrenght
		}
		if b.horisonatalSpeed < 0 {
			b.horisonatalSpeed += airFrictionStrenght
		}
	} else {
		b.verticalSpeed = 0
		if b.horisonatalSpeed > 0 {
			b.horisonatalSpeed -= groundFrictionStrenght
		}
		if b.horisonatalSpeed < 0 {
			b.horisonatalSpeed += groundFrictionStrenght
		}
	}

}

func (b *ball) hit(angle, power float64) {
	if b.isGhost==false {
		player.collisonGhost.hit(angle, power)
	}
	b.isGrounded = false
	b.horisonatalSpeed += power * math.Cos(angle)
	b.verticalSpeed += power * math.Sin(angle)
}

func (b *ball) move() {
	if b.horisonatalSpeed < 0.1 && b.horisonatalSpeed > -0.1 {
		b.horisonatalSpeed = 0
	}

	if b.isGhost==false {
		player.collisonGhost.move()
		if b.isGrounded == false {
			collisionDirection := b.checkForBallCollisions()
			processBounces(collisionDirection,b.collisonGhost)
			processBounces(collisionDirection,b)

		}
	}

	b.opts.GeoM.Translate(b.horisonatalSpeed, 0)
	b.position.X += b.horisonatalSpeed
	b.position.Y -= b.verticalSpeed
}

func processBounces(collisionDirection string, b *ball){
	if collisionDirection != "" {
		if strings.Contains(collisionDirection, "up") {
			b.upwardBounce()
		}
		if strings.Contains(collisionDirection, "down") {
			b.downwardBounce()
		}
		if strings.Contains(collisionDirection, "left") || strings.Contains(collisionDirection, "right") || strings.Contains(collisionDirection, "horisontal") {

			b.horizontalBounce()
		}
		if b.isGhost==false{
			b.resetGhostPosition()
		}
	} else {
		b.opts.GeoM.Translate(0, -b.verticalSpeed)
	}
}


func (b *ball) upwardBounce() {
	if b.verticalSpeed <= 0 {
		//fmt.Printf("%f\n", b.verticalSpeed)
		if b.verticalSpeed < 0.3 && b.verticalSpeed > -0.3 {
			fmt.Printf("stop\n")
			b.isGrounded = true
			b.verticalSpeed = 0
		} else {
			b.verticalSpeed = -b.verticalSpeed * 0.5
		}
	}
}

func (b *ball) downwardBounce() {
	b.verticalSpeed = -b.verticalSpeed * 0.5
}

func (b *ball) horizontalBounce() {
	b.horisonatalSpeed = -b.horisonatalSpeed * 0.5
}
