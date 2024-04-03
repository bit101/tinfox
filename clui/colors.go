// Package clui has command line ui functions
package clui

import (
	"strings"

	"github.com/bit101/go-ansi"
)

// Theme defines the colors used to print various elements.
type Theme struct {
	Headers      ansi.AnsiColor
	Instructions ansi.AnsiColor
	Errors       ansi.AnsiColor
	Defaults     ansi.AnsiColor
}

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

// DefaultTheme is the default theme.
var DefaultTheme = Theme{
	Headers:      ansi.BoldGreen,
	Instructions: ansi.Yellow,
	Errors:       ansi.BoldRed,
	Defaults:     ansi.Blue,
}

// SetTheme sets the theme colors from the config.
func SetTheme(headers, instructions, errors, defaults string) {
	DefaultTheme.Headers = ColorMap[strings.ToLower(headers)]
	DefaultTheme.Instructions = ColorMap[strings.ToLower(instructions)]
	DefaultTheme.Errors = ColorMap[strings.ToLower(errors)]
	DefaultTheme.Defaults = ColorMap[strings.ToLower(defaults)]
}
