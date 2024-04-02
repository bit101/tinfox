// package main is the shit
package main

import (
	"tinpig/templates"
)

/*
- cli standard stuff
- colors in config
*/

func main() {
	parser := templates.NewTemplateParser()
	parser.LoadAndParse()
}
