// Package cmd has the tinfox commands
package cmd

import (
	"github.com/bit101/tinfox/templates"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available templates",
	Long:  `List all available templates`,
	Run: func(cmd *cobra.Command, args []string) {
		parser := templates.NewTemplateParser()
		parser.DisplayList()
	},
}
