// package main is the shit
package main

import (
	"github.com/bit101/tinfox/cmd"
)

/*
TODO:
	- v (?) verbose/expert modes
- update old configs
- template sections, i.e.
	- tinfox go would show go templates, tinfox js would show js templates, etc.
	- tinfox kp user defined, fav templates or whatever.
- ansi.Printf(theme.Header, ...) -> theme.Headerf(...), etc.
*/

func main() {
	cmd.Execute()
}
