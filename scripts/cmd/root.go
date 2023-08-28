package cmd

import (
	"fmt"
	"os"

	"github.com/App/desktop/scripts/cmd/version"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(version.VersionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "meltctl-desktop",
	Short: "Utility scripts for developing App Desktop",
	Long:  "scripts such as checking version, updating version etc.",

	Run: func(cmd *cobra.Command, args []string) {
		color.HiCyan("MELTCTL DESKTOP")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
