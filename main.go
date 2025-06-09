package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "metaops",
	Short: "MetaOps CLI tool for development operations",
	Long: `MetaOps is a CLI tool that helps developers manage development operations
for Metadiv Technology's backend framework and services.`,
}

func main() {
	// Add subcommands
	rootCmd.AddCommand(metaginCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
