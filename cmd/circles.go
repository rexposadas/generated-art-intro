/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
	"github.com/spf13/cobra"
	"image/color"
	"math/rand"
	"time"
)

// circlesCmd represents the circles command
var circlesCmd = &cobra.Command{
	Use:   "circles",
	Short: "Generate circles",
	Long: `

Generate circles which makes you look like an artist!

`,
	Run: func(cmd *cobra.Command, args []string) {
		makeCircles()
	},
}

func makeCircles() {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}
	c := generativeart.NewCanva(800, 800)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewColorCircle2(30))
	c.ToPNG("circle.png")

}

func init() {
	rootCmd.AddCommand(circlesCmd)
}
