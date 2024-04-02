// package main is the shit
package main

import (
	"tinpig/templates"
)

/*
- make modules into structs with methods
- cli standard stuff
- colors in config
*/

func main() {
	parser := templates.NewTemplateParser()
	parser.GetTemplateChoice()
	parser.GetProjectDir()
	parser.DefineTokens()
	parser.CreateProject()
	parser.ShowSuccess()
}
