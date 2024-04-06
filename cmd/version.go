// Package cmd has the tinfox commands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.1.5"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tinfox",
	Long:  `Print the version number of tinfox`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tinfox project creator %s\n", version)
	},
}
