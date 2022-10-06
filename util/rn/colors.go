package rn

import (
	"github.com/jdxyw/generativeart/common"
	"image/color"
	"math/rand"
)

// Color returns a randomized color from a set.
// This is to provide difference colors to the generated art.
func Color() color.RGBA {

	// Use the colors from the imported library.
	list := []color.RGBA{
		common.Black,
		common.Azure,
		common.Plum,
		common.Orange,
		common.MediumAquamarine,
		common.PaleTurquoise,
	}

	return list[rand.Intn(len(list))]
}