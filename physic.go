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

const bounceSpeedReductionFactor float64=0.5

const numOfIndicatorGhosts=4
const ghostDistance=10

const ballSize float64 = 10

var useGroundFriction=false

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
	indicatorGhost					[numOfIndicatorGhosts]*ball
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
		tmp.setIndicators()
	}else{
		tmp.graphic.Fill(color.RGBA{0, 0, 255, 120})
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
	b.collisonGhost.verticalSpeed=b.verticalSpeed
	b.collisonGhost.horisonatalSpeed=b.horisonatalSpeed
	b.collisonGhost.move()
}

func (b *ball) applyNaturalForces() {
	if b.isGhost==false {
		player.collisonGhost.applyNaturalForces()
	}
	if b.isGrounded ==false{
		if b.verticalSpeed > -maxSpeed {
			b.verticalSpeed -= gravityStrenght
		}
	}else{
		b.verticalSpeed = 0
	}

	if useGroundFriction == false {
		if b.horisonatalSpeed > 0 {
			b.horisonatalSpeed -= airFrictionStrenght
		}
		if b.horisonatalSpeed < 0 {
			b.horisonatalSpeed += airFrictionStrenght
		}
	} else {

		if b.horisonatalSpeed > 0 {
			b.horisonatalSpeed -= groundFrictionStrenght
		}
		if b.horisonatalSpeed < 0 {
			b.horisonatalSpeed += groundFrictionStrenght
		}
	}

}

func (b *ball) hit(angle, power float64) {
	useGroundFriction=false
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
		if b.isGrounded && b.isGhost==false{
			b.setIndicators()
		}
	}

	if b.isGhost==false {
		player.collisonGhost.move()
		if b.verticalSpeed!=0 || b.horisonatalSpeed!=0 {
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
			b.verticalBounce()
		}
		if strings.Contains(collisionDirection, "down") {
			b.verticalBounce()
		}
		if strings.Contains(collisionDirection, "left") || strings.Contains(collisionDirection, "right") || strings.Contains(collisionDirection, "horisontal") {

			b.horizontalBounce()
		}
		if b.isGhost==false{
			fmt.Println(collisionDirection)
			b.resetGhostPosition()
		}
	} else {
		b.opts.GeoM.Translate(0, -b.verticalSpeed)
	}
}

func (b *ball) verticalBounce(){
		if b.isGhost==false && b.verticalSpeed < 0.3 && b.verticalSpeed > -0.3 {
			useGroundFriction=true
			if b.horisonatalSpeed==0 {
				b.isGrounded = true
				b.verticalSpeed = 0
				if b.isGhost == false {
					b.setIndicators()
				}
			}else{
				b.verticalSpeed = -b.verticalSpeed * bounceSpeedReductionFactor
			}
		} else {
			b.verticalSpeed = -b.verticalSpeed * bounceSpeedReductionFactor
		}
}

func (b *ball) horizontalBounce() {
	b.horisonatalSpeed = -b.horisonatalSpeed * bounceSpeedReductionFactor
}

func (b *ball) angledBounce(c triangleCollider){
	totalSpeed:=b.horisonatalSpeed+b.verticalSpeed
	var c1 float64=(c.Max.Y-c.Min.Y)/(c.Max.X-c.Min.X)
	var c2 float64=b.verticalSpeed/b.horisonatalSpeed

	relativeDeflectionAngle:=math.Atan2(c1-c2 , 1+(c1*c2))
	surfaceAngle:=math.Tan(c1)
	absoluteReflectionAngle:=-(relativeDeflectionAngle+surfaceAngle)

	b.horisonatalSpeed =  math.Cos(absoluteReflectionAngle)
	b.verticalSpeed =  math.Sin(absoluteReflectionAngle)
	fmt.Println(b.horisonatalSpeed)
	fmt.Println(b.verticalSpeed)

	b.horisonatalSpeed *=totalSpeed
	b.verticalSpeed *=totalSpeed
}

func (b *ball) setIndicators() {
	if b.isGhost {
		fmt.Println("tried to set indicators for ghost")
	}else{
		x,y:=b.position.X,b.position.Y
		b.indicatorGhost[0]=makeBall(x,y,true)
		b.indicatorGhost[0].hit(b.controls.angle, b.controls.power)
		for i:=0;i< len(b.indicatorGhost);i++ {
			for j := 0; j < ghostDistance; j++ {
				b.indicatorGhost[i].applyNaturalForces()
				b.indicatorGhost[i].move()
				processBounces("", b.indicatorGhost[i])
			}
			if i<len(b.indicatorGhost)-1{
				x, y = b.indicatorGhost[i].position.X, b.indicatorGhost[i].position.Y
				b.indicatorGhost[i+1] = makeBall(x, y, true)
				b.indicatorGhost[i+1].isGrounded=false
				b.indicatorGhost[i+1].verticalSpeed = b.indicatorGhost[i].verticalSpeed
				b.indicatorGhost[i+1].horisonatalSpeed = b.indicatorGhost[i].horisonatalSpeed
			}

		}


	}
}

func (b *ball) getMoveDirection() string{
	s:=""
	if b.verticalSpeed>0{
		s+="up"
	}else{
		s+="down"
	}
	if b.horisonatalSpeed>0{
		s+="right"
	}else{
		s+="left"
	}
	return s
}













