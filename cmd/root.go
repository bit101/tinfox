// Package cmd has the tinfox commands
package cmd

import (
	"fmt"
	"os"

	"github.com/bit101/tinfox/templates"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tinfox",
	Short: "tinfox builds custom projects based on project templates.",
	Long:  `tinfox builds custom projects based on project templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		parser := templates.NewTemplateParser()
		parser.LoadAndParse()
	},
}

// Execute runs the app.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
