// Package clui has command line ui functions
package clui

import (
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
	ColorMap["Black"] = ansi.Black
	ColorMap["BoldBlack"] = ansi.BoldBlack
	ColorMap["Red"] = ansi.Red
	ColorMap["BoldRed"] = ansi.BoldRed
	ColorMap["Green"] = ansi.Green
	ColorMap["BoldGreen"] = ansi.BoldGreen
	ColorMap["Yellow"] = ansi.Yellow
	ColorMap["BoldYellow"] = ansi.BoldYellow
	ColorMap["Blue"] = ansi.Blue
	ColorMap["BoldBlue"] = ansi.BoldBlue
	ColorMap["Purple"] = ansi.Purple
	ColorMap["BoldPurple"] = ansi.BoldPurple
	ColorMap["Cyan"] = ansi.Cyan
	ColorMap["BoldCyan"] = ansi.BoldCyan
	ColorMap["White"] = ansi.White
	ColorMap["BoldWhite"] = ansi.BoldWhite
}

// DefaultTheme is the default theme.
var DefaultTheme = NewTheme(ansi.BoldGreen, ansi.Yellow, ansi.BoldRed, ansi.Blue)

// NewTheme creates a new theme with the given colors.
func NewTheme(headers, instructions, errors, defaults ansi.AnsiColor) Theme {
	return Theme{
		headers,
		instructions,
		errors,
		defaults,
	}
}

// SetTheme sets the theme colors from the config.
func SetTheme(headers, instructions, errors, defaults string) {
	DefaultTheme.Headers = ColorMap[headers]
	DefaultTheme.Instructions = ColorMap[instructions]
	DefaultTheme.Errors = ColorMap[errors]
	DefaultTheme.Defaults = ColorMap[defaults]
}
