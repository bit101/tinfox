// Package theme controls colors
package theme

import (
	"strings"

	"github.com/bit101/go-ansi"
)

// Theme elements that define the colors used to print various elements.
var (
	Header      ansi.AnsiColor
	Instruction ansi.AnsiColor
	Error       ansi.AnsiColor
	Default     ansi.AnsiColor
)

// ColorMap helps convert strings to ansi colors.
var ColorMap = map[string]ansi.AnsiColor{}

func init() {
	ColorMap["black"] = ansi.Black
	ColorMap["boldblack"] = ansi.BoldBlack
	ColorMap["red"] = ansi.Red
	ColorMap["boldred"] = ansi.BoldRed
	ColorMap["green"] = ansi.Green
	ColorMap["boldgreen"] = ansi.BoldGreen
	ColorMap["yellow"] = ansi.Yellow
	ColorMap["boldyellow"] = ansi.BoldYellow
	ColorMap["blue"] = ansi.Blue
	ColorMap["boldblue"] = ansi.BoldBlue
	ColorMap["purple"] = ansi.Purple
	ColorMap["boldpurple"] = ansi.BoldPurple
	ColorMap["cyan"] = ansi.Cyan
	ColorMap["boldcyan"] = ansi.BoldCyan
	ColorMap["white"] = ansi.White
	ColorMap["boldwhite"] = ansi.BoldWhite
}

// SetTheme sets the theme colors from the config.
func SetTheme(headers, instructions, errors, defaults string) {
	Header = ColorMap[strings.ToLower(headers)]
	Instruction = ColorMap[strings.ToLower(instructions)]
	Error = ColorMap[strings.ToLower(errors)]
	Default = ColorMap[strings.ToLower(defaults)]
}
