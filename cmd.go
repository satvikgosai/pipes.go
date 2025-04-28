package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	speedStr string
	speedInt int
	themeStr string
	styleStr string
)

var (
	themeArgs = strings.Join(slices.Collect(maps.Keys(ColorThemes)), ", ")
	styleArgs = strings.Join(slices.Collect(maps.Keys(PipeStyles)), ", ")
)

var rootCmd = &cobra.Command{
	Use:   "pipes",
	Short: "Pipes screen saver CLI application",
	Long: `Pipes is a screen saver CLI application that creates an animated pipe maze effect.
The animation speed can be adjusted from 0 (slowest) to 100 (fastest).`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		parsedSpeed, err := strconv.Atoi(speedStr)
		if err != nil {
			return fmt.Errorf("invalid value for --speed: must be an integer")
		}

		if parsedSpeed < MinSpeed || parsedSpeed > MaxSpeed {
			return fmt.Errorf("invalid value for --speed: must be from %d to %d", MinSpeed, MaxSpeed)
		}

		speedInt = parsedSpeed

		themeStr = strings.ToLower(themeStr)
		if _, ok := ColorThemes[themeStr]; !ok {
			return fmt.Errorf("invalid value for --theme: must be one of (%s)", themeArgs)
		}

		styleStr = strings.ToLower(styleStr)
		if _, ok := PipeStyles[styleStr]; !ok {
			return fmt.Errorf("invalid value for --style: must be one of (%s)", styleArgs)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&speedStr, "speed", "s", "50", fmt.Sprintf("Set the pipes speed (%d-%d)", MinSpeed, MaxSpeed))
	rootCmd.Flags().StringVarP(&themeStr, "theme", "t", "default", fmt.Sprintf("Set the theme (%s)", themeArgs))
	rootCmd.Flags().StringVarP(&styleStr, "style", "l", "default", fmt.Sprintf("Set the style (%s)", styleArgs))
}

func Execute() error {
	return rootCmd.Execute()
}
