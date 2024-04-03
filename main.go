// package main is the shit
package main

import (
	"tinpig/templates"
)

/*
- DefaultTheme => ActiveTheme
- Clean up theme element names
- cli standard stuff
*/

func main() {
	parser := templates.NewTemplateParser()
	parser.LoadAndParse()
}
