package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Really basic version holder
const Version = "0.0.1"

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Application version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("App version: %v\n", Version)
	},
}

func init() {
	rootCommand.AddCommand(versionCommand)
}
