package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	speedStr string
	cfg      *Config
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

		cfg = NewConfig(parsedSpeed)
		if parsedSpeed < cfg.MinSpeed || parsedSpeed > cfg.MaxSpeed {
			return fmt.Errorf("invalid value for --speed: must be from %d to %d", cfg.MinSpeed, cfg.MaxSpeed)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cfg)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&speedStr, "speed", "s", "50", "Set the pipes speed (0-100)")
}

func Execute() error {
	return rootCmd.Execute()
}
