package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var Context []string

func lookup(a int, b int) (ret int) {
	ret = (a + b) % b
	return
}

func Print(message string) {

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

	for index, s := range Context {
		colorIndex := lookup(index, length)
		color.Set(colors[colorIndex])
		if len(Context) > 1 {
			fmt.Print(s + "/")
		} else {
			fmt.Print(s)
		}
	}

	color.Set(color.FgWhite)
	fmt.Println("]", message)

}

func Printf(message string, v ...any) {
	formatted := fmt.Sprintf(message, v...)
	Print((formatted))
}

func PrintErr(message string) {

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

	for index, s := range Context {
		colorIndex := lookup(index, length)
		color.Set(colors[colorIndex])
		if len(Context) > 1 {
			fmt.Print(s + "/")
		} else {
			fmt.Print(s)
		}
	}

	color.Set(color.FgWhite)
	color.Set(color.FgRed)
	fmt.Print("] ")

	color.Set(color.FgRed)
	fmt.Println(message)
	os.Exit(1)

}

func PrintErrf(message string, v ...any) {
	formatted := fmt.Sprintf(message, v...)
	PrintErr((formatted))
}
