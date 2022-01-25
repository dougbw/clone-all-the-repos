package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func lookup(a int, b int) (ret int) {
	ret = (a + b) % b
	return
}

func PrintLogMessage(context []string, message string) {

	colors := []color.Attribute{
		color.FgBlue,
		color.FgGreen,
		color.FgHiBlue,
		color.FgMagenta,
		color.FgCyan,
		color.FgYellow,
		color.FgHiGreen,
		color.FgHiYellow,
		color.FgHiMagenta,
		color.FgHiCyan,
	}
	length := len(colors)

	fmt.Print("[")

	for index, s := range context {
		colorIndex := lookup(index, length)
		color.Set(colors[colorIndex])
		fmt.Print(s + "/")
	}

	color.Set(color.FgWhite)
	fmt.Println("]", message)

}

func PrintErrMessage(context []string, message string) {

	colors := []color.Attribute{
		color.FgBlue,
		color.FgGreen,
		color.FgHiBlue,
		color.FgMagenta,
		color.FgCyan,
		color.FgYellow,
		color.FgHiGreen,
		color.FgHiYellow,
		color.FgHiMagenta,
		color.FgHiCyan,
	}
	length := len(colors)

	fmt.Print("[")

	for index, s := range context {
		colorIndex := lookup(index, length)
		color.Set(colors[colorIndex])
		fmt.Print(s + "/")
	}

	color.Set(color.FgWhite)
	fmt.Print("] ")

	color.Set(color.FgRed)
	fmt.Println(message)
	os.Exit(1)

}
