// package main is the shit
package main

import (
	"tinpig/templates"
)

/*
- cli standard stuff
*/

func main() {
	parser := templates.NewTemplateParser()
	parser.LoadAndParse()
}
