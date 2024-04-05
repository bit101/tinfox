// package main is the shit
package main

import (
	"tinpig/templates"
)

/*
- cobra
- ansi.Printf(theme.Header, ...) -> theme.Headerf(...), etc.
*/

func main() {
	parser := templates.NewTemplateParser()
	parser.LoadAndParse()
}
