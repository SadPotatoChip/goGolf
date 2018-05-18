package main

func (c triangleCollider) bottomLeftTriangleCollision(b *ball) string{
	s:=""
	switch b.getMoveDirection() {
	case "downleft":
		if b.position.Y < c.Min.Y && b.collisonGhost.position.Y+b.size > c.Min.Y {
			s = "up"
		}
		if b.position.X+ b.size > c.Max.X && b.collisonGhost.position.X < c.Max.X{
			s="right"
		}
	case "downright":
		if b.position.Y < c.Min.Y && b.collisonGhost.position.Y+b.size > c.Min.Y {
			s = "up"
		}else{
			d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y},c)
			d2 := calcSide(vector2{b.collisonGhost.position.X + ballSize, b.collisonGhost.position.Y}, c)
			if d1!=d2 {
				if b.verticalSpeed==0{
					s="left"
				}else {
					b.angledBounce(c)
					s += "angled"
				}
			}
		}
	case "upleft":
		if b.position.X+ b.size > c.Max.X && b.collisonGhost.position.X < c.Max.X{
			s="right"
		}else{
			d1 := calcSide(vector2{b.position.X , b.position.Y+ballSize}, c)
			d2 := calcSide(vector2{b.collisonGhost.position.X + ballSize, b.collisonGhost.position.Y}, c)
			if d1!=d2 {
				b.angledBounce(c)
				s+="angled"
			}
		}
	case "upright":
		d1 := calcSide(vector2{b.position.X , b.position.Y+ballSize}, c)
		d2 := calcSide(vector2{b.collisonGhost.position.X + ballSize, b.collisonGhost.position.Y}, c)
		if d1!=d2 {
			b.angledBounce(c)
			s+="angled"
		}
	}
	return s
}

func (c triangleCollider) topLeftTriangleCollision(b *ball) string{
	s:=""
	switch b.getMoveDirection() {
	case "downleft":
		if b.position.X+ b.size > c.Max.X && b.collisonGhost.position.X < c.Max.X{
			s="right"
		}else{
			d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y}, c)
			d2 := calcSide(vector2{b.collisonGhost.position.X + ballSize, b.collisonGhost.position.Y+ballSize}, c)
			if d1!=d2 {
				if b.verticalSpeed==0{
					s="left"
				}else {
					b.angledBounce(c)
					s += "angled"
				}
			}
		}
	case "downright":
		d1 := calcSide(vector2{b.position.X , b.position.Y+ballSize}, c)
		d2 := calcSide(vector2{b.collisonGhost.position.X + ballSize, b.collisonGhost.position.Y + ballSize}, c)
		if d1!=d2 {
			b.angledBounce(c)
			s+="angled"
		}
	case "upleft":
		if b.position.Y+b.size > c.Max.Y && b.collisonGhost.position.Y < c.Max.Y {
			s = "down"
		}
		if b.position.X+ b.size > c.Max.X && b.collisonGhost.position.X < c.Max.X{
			s="right"
		}
	case "upright":
		if b.position.Y+b.size > c.Max.Y && b.collisonGhost.position.Y < c.Max.Y {
			s = "down"
		}else {
			d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y}, c)
			d2 := calcSide(vector2{b.collisonGhost.position.X + ballSize, b.collisonGhost.position.Y + ballSize}, c)
			if d1 != d2 {
				if b.verticalSpeed == 0 {
					s = "left"
				} else {
					b.angledBounce(c)
					s += "angled"
				}
			}
		}
	}
	return s
}

func (c triangleCollider) bottomRightTriangleCollision(b *ball) string{
	s:=""
	switch b.getMoveDirection() {
	case "downleft":
		if b.position.Y < c.Min.Y && b.collisonGhost.position.Y+b.size > c.Min.Y {
			s = "up"
		}else{
			d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y+ballSize}, c)
			d2 := calcSide(vector2{b.collisonGhost.position.X, b.collisonGhost.position.Y}, c)
			if d1!=d2 {
				if b.verticalSpeed==0{
					s="right"
				}else {
					b.angledBounce(c)
					s += "angled"
				}
			}
		}
	case "downright":
		if b.position.Y < c.Min.Y && b.collisonGhost.position.Y+b.size > c.Min.Y {
			s = "up"
		}
		if b.position.X< c.Min.X && b.collisonGhost.position.X+b.size>c.Min.X{
			s="left"
		}
	case "upleft":
		d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y+ballSize}, c)
		d2 := calcSide(vector2{b.collisonGhost.position.X, b.collisonGhost.position.Y}, c)
		if d1!=d2 {
			b.angledBounce(c)
			s+="angled"
		}
	case "upright":
		if b.position.X< c.Min.X && b.collisonGhost.position.X+b.size>c.Min.X{
			s="left"
		}else{
			d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y+ballSize}, c)
			d2 := calcSide(vector2{b.collisonGhost.position.X, b.collisonGhost.position.Y}, c)
			if d1!=d2 {
				if b.verticalSpeed==0{
					s="right"
				}else {
					b.angledBounce(c)
					s += "angled"
				}
			}
		}
	}
	return s
}

func (c triangleCollider) topRightTriangleCollision(b *ball) string{
	s:=""
	switch b.getMoveDirection() {
	case "downleft":
		d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y}, c)
		d2 := calcSide(vector2{b.collisonGhost.position.X, b.collisonGhost.position.Y+ballSize}, c)
		if d1!=d2 {
			b.angledBounce(c)
			s+="angled"
		}
	case "downright":
		if b.position.X< c.Min.X && b.collisonGhost.position.X+b.size>c.Min.X{
			s="left"
		}else{
			d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y}, c)
			d2 := calcSide(vector2{b.collisonGhost.position.X, b.collisonGhost.position.Y+ballSize}, c)
			if d1!=d2 {
				b.angledBounce(c)
				s+="angled"
			}
		}
	case "upleft":
		if b.position.Y+b.size > c.Max.Y && b.collisonGhost.position.Y < c.Max.Y {
			s = "down"
		}else{
			d1 := calcSide(vector2{b.position.X + ballSize, b.position.Y}, c)
			d2 := calcSide(vector2{b.collisonGhost.position.X, b.collisonGhost.position.Y+ballSize}, c)
			if d1!=d2 {
				b.angledBounce(c)
				s+="angled"
			}
		}
	case "upright":
		if b.position.X< c.Min.X && b.collisonGhost.position.X+b.size>c.Min.X{
			s="left"
		}
		if b.position.Y+b.size > c.Max.Y && b.collisonGhost.position.Y < c.Max.Y {
			s = "down"
		}
	}
	return s
}

func calcSide(p vector2, c triangleCollider) bool{
	var xa,xb,ya, yb float64
	switch c.Missing_part{
	case "bottom-left":
		xa=c.Min.X
		xb=c.Max.X
		ya=c.Min.Y
		yb=c.Max.Y
	case "top-left":
		xa=c.Min.X
		xb=c.Max.X
		ya=c.Max.Y
		yb=c.Min.Y
	case "bottom-right":
		xa=c.Min.X
		xb=c.Max.X
		ya=c.Max.Y
		yb=c.Min.Y
	case "top-right":
		xa=c.Min.X
		xb=c.Max.X
		ya=c.Min.Y
		yb=c.Max.Y
	}

	d:=(p.X-xa)*(yb-ya)-(p.Y-ya)*(xb-xa)
	if d>0{
		return true
	}else{
		return false
	}
}