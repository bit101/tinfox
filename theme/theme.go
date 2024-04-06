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

//////////////////////////////
// Header
//////////////////////////////

// PrintHeaderf prints a formatted string with the header style.
func PrintHeaderf(formatString string, vars ...any) {
	ansi.Printf(Header, formatString, vars...)
}

// PrintHeaderln prints a formatted string with the header style.
func PrintHeaderln(value ...any) {
	ansi.Println(Header, value...)
}

// PrintHeader prints a formatted string with the header style.
func PrintHeader(value ...any) {
	ansi.Print(Header, value...)
}

//////////////////////////////
// Instruction
//////////////////////////////

// PrintInstructionf prints a formatted string with the instruction style.
func PrintInstructionf(formatString string, vars ...any) {
	ansi.Printf(Instruction, formatString, vars...)
}

// PrintInstructionln prints a formatted string with the instruction style.
func PrintInstructionln(value ...any) {
	ansi.Println(Instruction, value...)
}

// PrintInstruction prints a formatted string with the instruction style.
func PrintInstruction(value ...any) {
	ansi.Print(Instruction, value...)
}

//////////////////////////////
// Error
//////////////////////////////

// PrintErrorf prints a formatted string with the error style.
func PrintErrorf(formatString string, vars ...any) {
	ansi.Printf(Error, formatString, vars...)
}

// PrintErrorln prints a formatted string with the error style.
func PrintErrorln(value ...any) {
	ansi.Println(Error, value...)
}

// PrintError prints a formatted string with the error style.
func PrintError(value ...any) {
	ansi.Print(Error, value...)
}

//////////////////////////////
// Default
//////////////////////////////

// PrintDefaultf prints a formatted string with the default style.
func PrintDefaultf(formatString string, vars ...any) {
	ansi.Printf(Default, formatString, vars...)
}

// PrintDefaultln prints a formatted string with the default style.
func PrintDefaultln(value ...any) {
	ansi.Println(Default, value...)
}

// PrintDefault prints a formatted string with the default style.
func PrintDefault(value ...any) {
	ansi.Print(Default, value...)
}
