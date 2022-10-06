package circles

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/rexposadas/generated-art-intro/util/rn"
	"image/color"
	"math/rand"
	"time"
)

func Gradient() {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		rn.Color(),
		rn.Color(),
		rn.Color(),
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}
	c := generativeart.NewCanva(800, 800)
	c.SetBackground(rn.Color())
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewColorCircle2(30))
	c.ToPNG("circle.png")

}
