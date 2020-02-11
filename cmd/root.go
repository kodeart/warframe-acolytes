package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "acolytes",
	Short: "Warframe Acolyte Tracker",
	Long: `The application shows the name of the acolyte,
discovered location, the mission type with the
enemy faction, and the current health of the acolyte.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
