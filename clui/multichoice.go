// Package clui has ui functions
package clui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bit101/go-ansi"
)

// MultiChoice presents a multichoice menu
func MultiChoice(choices []string, instructions string) (int, string) {
	var choice int
	count := len(choices)
	errStr := ""

	for i := 0; i < count+4; i++ {
		fmt.Println()
	}
	ansi.MoveUp(count + 4)
	ansi.Save()

	reader := bufio.NewReader(os.Stdin)

	ok := false
	for ok == false {
		outputMultiChoice(choices, instructions, errStr)
		ok = true
		errStr = ""

		// check raw input
		input, err := reader.ReadString('\n')
		if err != nil {
			errStr = fmt.Sprint("Unable to read the response.\n")
			ok = false
		}
		input = strings.TrimSuffix(input, "\n")

		// quit?
		if strings.ToLower(input) == "q" {
			os.Exit(0)
		}

		// parse int
		choice64, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			errStr = fmt.Sprintf("Choose between 1 and %d or 'q'\n", count)
			ok = false
		}

		// check choices
		choice = int(choice64)
		if ok && (choice < 1 || choice > count) {
			errStr = fmt.Sprintf("Choose between 1 and %d or 'q'\n", count)
			ok = false
		}

	}
	result := choices[choice-1]
	ansi.Restore()
	ansi.ClearToEnd()
	return choice - 1, result
}

func outputMultiChoice(choices []string, instructions, errStr string) {
	ansi.Restore()
	ansi.ClearToEnd()
	if errStr != "" {
		ansi.Print(DefaultTheme.Errors, errStr)
	}
	ansi.Println(DefaultTheme.Headers, instructions, "\r")

	for i := 0; i < len(choices); i++ {
		fmt.Printf("%d. %s\r\n", i+1, choices[i])
	}
	fmt.Println("q. Quit")
	ansi.Print(DefaultTheme.Instructions, "Choice: ")

}
