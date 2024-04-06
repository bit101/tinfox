// package main is the shit
package main

import "github.com/bit101/tinpig2/templates"

/*
- cobra
- update old configs
- ansi.Printf(theme.Header, ...) -> theme.Headerf(...), etc.
*/

func main() {
	parser := templates.NewTemplateParser()
	parser.LoadAndParse()
}
