package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

const controlerStartAngle, controlerStartPower float64 = 0, 1
const defaultAngleLerp, defaultpowerLerp float64 = math.Pi / 48, 0.1
const defaultMaxPower float64 = 20
const defaultIndicatorDistanceFromBall float64 = 10

var hitKeyIsDown = false

type controler struct {
	angle float64
	power float64

	_angleLerp float64
	_powerLerp float64
	_maxPower  float64

	powerUpKey, powerDownKey ebiten.Key
	angleUpKey, angleDownKey ebiten.Key
	hitKey                   ebiten.Key
	indicator                *angleIndicator

	//makes fentching position of ball cleaner in some places
	parent *ball
}

type angleIndicator struct {
	position         vector2
	graphic          *ebiten.Image
	opts             *ebiten.DrawImageOptions
	distanceFromBall float64
}

func makeControler(b *ball) *controler {
	c := new(controler)
	c.angle = controlerStartAngle
	c.power = controlerStartPower
	c._maxPower = defaultMaxPower
	c._angleLerp = defaultAngleLerp
	c._powerLerp = defaultpowerLerp
	c.powerUpKey = ebiten.KeyUp
	c.powerDownKey = ebiten.KeyDown
	c.angleUpKey = ebiten.KeyLeft
	c.angleDownKey = ebiten.KeyRight
	c.hitKey = ebiten.KeySpace

	c.parent = b

	c.makeIndicator()

	return c
}

func (c *controler) makeIndicator() {
	tmp := new(angleIndicator)
	tmp.graphic, _ = ebiten.NewImage(2, 2, ebiten.FilterNearest)
	tmp.graphic.Fill(color.White)
	tmp.opts = &ebiten.DrawImageOptions{}
	tmp.distanceFromBall = defaultIndicatorDistanceFromBall

	//HACK - not bound to starting angle (assumes it is 0)
	tmp.position.x = c.parent.position.x - tmp.distanceFromBall
	tmp.position.y = c.parent.position.y + 1
	tmp.opts.GeoM.Translate(tmp.position.x, tmp.position.y)

	c.indicator = tmp
}

func (c *controler) changeAngle(dir float64) {
	c.angle += c._angleLerp * dir

	//indicator orbit
	/*	x, y := math.Cos(c._angleLerp*dir), math.Sin(c._angleLerp*dir)
		c.indicator.position.x += x
		c.indicator.position.y += y
		c.indicator.opts.GeoM.Translate(x, y)*/

	if c.angle == 2*math.Pi {
		c.angle = 0
	}
	if c.angle == -2*math.Pi {
		c.angle = 0
	}
}

func (c *controler) changePower(dir float64) {
	if c.power < c._maxPower && dir > 0 {
		c.power += c._powerLerp
	}
	if c.power > defaultpowerLerp && dir < 0 {
		c.power -= c._powerLerp
	}
}

func hitKeyDown(hitKey ebiten.Key) bool {
	if ebiten.IsKeyPressed(hitKey) {
		if hitKeyIsDown == false {
			hitKeyIsDown = true
			return true
		}
		return false
	}
	hitKeyIsDown = false
	return false
}
