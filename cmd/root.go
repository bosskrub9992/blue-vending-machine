package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "blue vending machine",
	Short: "a simple vending machine",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
