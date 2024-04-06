// package main is the shit
package main

import "github.com/bit101/tinpig2/templates"

/*
TODO:
- cobra
	- h
	- v (?) verbose/expert modes
- update old configs
- change executable name to just tinpig (not tinpig2)
- template sections, i.e.
	- tinpig go would show go templates, tinpig js would show js templates, etc.
	- tinpig kp user defined, fav templates or whatever.
- ansi.Printf(theme.Header, ...) -> theme.Headerf(...), etc.
*/

func main() {
	parser := templates.NewTemplateParser()
	parser.LoadAndParse()
}
