package cmd

import (
	"github.com/rexposadas/generated-art-intro/util/circles"
	"github.com/spf13/cobra"
)

// circlesCmd represents the circles command
var circlesCmd = &cobra.Command{
	Use:   "circles",
	Short: "Generate circles",
	Long: `

Generate circles which makes you look like an artist!

`,
	Run: func(cmd *cobra.Command, args []string) {
		circles.Gradient()
	},
}

func init() {
	rootCmd.AddCommand(circlesCmd)
}
