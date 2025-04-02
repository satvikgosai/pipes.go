package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var speedStr string
var speed int

var rootCmd = &cobra.Command{
	Use:   "pipes",
	Short: "Pipes screen saver CLI application",
	Long:  `Pipes is a screen saver CLI application built using xTerm and Cobra in Go.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		parsedSpeed, err := strconv.Atoi(speedStr)
		if err != nil {
			return fmt.Errorf("invalid value for --speed: must be an integer")
		}
		if parsedSpeed < 0 || parsedSpeed > 100 {
			return fmt.Errorf("invalid value for --speed: must be from 0 to 100")
		}
		speed = parsedSpeed
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(speed)
	},
}

func Execute() error {
	rootCmd.Flags().StringVarP(&speedStr, "speed", "s", "50", "Set the pipes speed")
	return rootCmd.Execute()
}
