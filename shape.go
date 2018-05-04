package main

import "github.com/hajimehoshi/ebiten"

type shape interface{
	getMin() vector2
	getMax() vector2
	getGraphic() *ebiten.Image
	getOpts() *ebiten.DrawImageOptions
}

type box struct {
	Collider boxCollider
	Graphic  *ebiten.Image
	Opts     *ebiten.DrawImageOptions
}

type triangle struct {
	Collider triangleCollider
	Graphic  *ebiten.Image
	Opts     *ebiten.DrawImageOptions
}

type uninteractableImage struct {
	Graphic *ebiten.Image
	Opts    *ebiten.DrawImageOptions
}

func (img *uninteractableImage)draw(screen *ebiten.Image){
	screen.DrawImage(img.Graphic, img.Opts)
}

func (t triangle)getMin() vector2{
	return t.Collider.Min;
}

func (t triangle)getMax() vector2{
	return t.Collider.Max;
}

func (b box)getMin() vector2{
	return b.Collider.Min;
}

func (b box)getMax() vector2{
	return b.Collider.Max;
}

func (t triangle)getGraphic() *ebiten.Image{
	return t.Graphic
}

func (b box)getGraphic() *ebiten.Image{
	return b.Graphic
}

func (t triangle)getOpts() *ebiten.DrawImageOptions{
	return t.Opts
}

func (b box)getOpts() *ebiten.DrawImageOptions{
	return b.Opts
}
