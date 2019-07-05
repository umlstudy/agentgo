package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"github.com/mitchellh/colorstring"
)

func test7() {
	{
		// ansi.Print("\x1b[2J")
		for i := 0; i < 30; i++ {
			colorstring.Fprintln(ansi.NewAnsiStdout(), "[green]mitchellh/colorstring-----------------------------")
		}
		colorstring.Fprintln(ansi.NewAnsiStdout(), "[yellow]yellow mitchellh/colorstring-----------------------------")
		ansi.CursorUp(1)
		//ansi.EraseInLine(1)
	}
	{
		colorstring.Println("[blue]Hello [red]World!")
		colorstring.Println(" World!")
		fmt.Println(" World!")

	}
	{
		color.Cyan("Prints text in cyan.")

		// a newline will be appended automatically
		color.Blue("Prints %s in blue.", "text")

		// More default foreground colors..
		color.Red("We have red")
		color.Yellow("Yellow color too!")
		color.Magenta("And many others ..")

		// Hi-intensity colors
		color.HiGreen("Bright green color.")
		color.HiBlack("Bright black means gray..")
		color.HiWhite("Shiny white color!")

		// Create a new color object
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Println("Prints cyan text with an underline.")

		// Or just add them to New()
		d := color.New(color.FgCyan, color.Bold)
		d.Printf("This prints bold cyan %s\n", "too!.")

		// Mix up foreground and background colors, create new mixes!
		red := color.New(color.FgRed)

		boldRed := red.Add(color.Bold)
		boldRed.Println("This will print text in bold red.")

		whiteBackground := red.Add(color.BgWhite)
		whiteBackground.Println("Red text with White background.")

		// Use your own io.Writer output
		color.New(color.FgBlue).Fprintln(os.Stdout, "blue color!")

		blue := color.New(color.FgBlue)
		blue.Fprint(os.Stdout, "This will print text in blue.")
	}

	{
		// Create a custom print function for convenient
		red := color.New(color.FgRed).PrintfFunc()
		red("warning")
		red("error: %s", fmt.Errorf("aaa"))

		// Mix up multiple attributes
		notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
		notice("don't forget this...")
	}

	{
		// blue := color.New(FgBlue).FprintfFunc()
		// blue(os.Stdout, "important notice: %s", stars)

		// Mix up with multiple attributes
		success := color.New(color.Bold, color.FgGreen).FprintlnFunc()
		success(os.Stdout, "don't forget this...")
	}
}
