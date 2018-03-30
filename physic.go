package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

const maxSpeed float64 = 40
const gravityStrenght float64 = 0.05
const airFrictionStrenght float64 = 0.002
const groundFrictionStrenght float64 = 0.05
const inputAngleLerp float64 = math.Pi / 6

type ball struct {
	position                        vector2
	graphic                         *ebiten.Image
	opts                            *ebiten.DrawImageOptions
	controls                        *controler
	verticalSpeed, horisonatalSpeed float64
	isGrounded                      bool
}

func makeBall(x, y float64) *ball {
	tmp := new(ball)
	tmp.graphic, _ = ebiten.NewImage(5, 5, ebiten.FilterNearest)
	tmp.graphic.Fill(color.White)
	tmp.opts = &ebiten.DrawImageOptions{}
	tmp.position.x = x
	tmp.position.y = y
	tmp.opts.GeoM.Translate(x, y)
	tmp.verticalSpeed, tmp.horisonatalSpeed = 0, 0
	tmp.controls = makeControler(tmp)
	tmp.isGrounded = true

	return tmp
}

func (b *ball) applyNaturalForces() {
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
	b.isGrounded = false
	b.horisonatalSpeed += power * math.Cos(angle)
	b.verticalSpeed += power * math.Sin(angle)
}

func (b *ball) move() {
	if b.horisonatalSpeed < 0.1 && b.horisonatalSpeed > -0.1 {
		b.horisonatalSpeed = 0
	}
	//HACK
	if b.position.y > screenHeight-10 && b.isGrounded == false {
		b.position.y = screenHeight - 10
		b.opts.GeoM.Reset()
		b.opts.GeoM.Translate(b.position.x, b.position.y)
		b.verticalBounce()
	} else {
		b.opts.GeoM.Translate(0, -b.verticalSpeed)
	}
	b.opts.GeoM.Translate(b.horisonatalSpeed, 0)

	b.position.x += b.horisonatalSpeed
	b.position.y -= b.verticalSpeed
}

func (b *ball) verticalBounce() {
	fmt.Printf("%f\n", b.verticalSpeed)
	if b.verticalSpeed < 0.3 && b.verticalSpeed > -0.3 {
		fmt.Printf("stop\n")
		b.isGrounded = true
		b.verticalSpeed = 0
	} else {
		b.verticalSpeed = -b.verticalSpeed * 0.5
	}
}
