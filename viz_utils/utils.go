package viz_utils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func ResizeImage(src *ebiten.Image, width, height int) *ebiten.Image {
	target := ebiten.NewImage(width, height)
	sx := float64(width) / float64(src.Bounds().Dx())
	sy := float64(height) / float64(src.Bounds().Dy())
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(sx, sy)
	target.DrawImage(src, opts)
	return target
}
