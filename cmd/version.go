package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// vars injected with ldflags at build time (this can be done automatically by goreleaser)
	version = "unknown"
	commit  = "unknown"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   `version`,
	Short: "Prints badrobot version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version %s\ngit commit %s\n", version, commit)
	},
}
