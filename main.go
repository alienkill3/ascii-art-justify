package main

// main.go is the entry point of the ASCII Art Color application.
// It orchestrates the complete workflow: parsing CLI arguments, validating input,
// loading the selected ASCII art font, building a font map, detecting terminal width,
// rendering the ASCII art with specified styling (color, alignment, substring coloring),
// and writing the output to either stdout or a file. It handles all error cases with
// appropriate error messages to stderr.

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/src"
)

type hexFlag int

func main() {
	cfg, err := src.ParseCLIArgs(os.Args[1:])
	// if cfg.Input == "" && err != nil {
	// 	// if len(os.Args) > 1 {
	// 	// 	fmt.Fprintln(os.Stderr, "Usage: go run . [FLAGS] [SUBSTRING] [STRING] [BANNER]\n\nEX: go run . --color=red so something standard")
	// 	// 	os.Exit(1)
	// 	// }
	// 	return
	// }
	if err != nil {
		// If no arguments were provided at all, silently exit
		if err == src.ErrUsage && cfg.Input == "" && len(os.Args) < 2 {
			return
		}
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if !src.IsValidAscii(cfg.Input) {
		fmt.Fprintln(os.Stderr, "Error: found non ascii character")
		os.Exit(1)
	}

	fontData := src.GetReader(cfg.Font)
	fontMap, err := src.BuildFontMap(fontData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing font map")
		os.Exit(1)
	}

	ws, err := src.IoctlGetWinsize(int(os.Stdout.Fd()))
	width := uint16(80) // default width
	if err == nil {
		width = ws.Cols
	}

	// Render once, then choose output target (stdout or file).
	input := strings.ReplaceAll(cfg.Input, "\\n", "\n")
	content := src.RenderInputString(input, fontMap, cfg.Options, width)

	if cfg.OutputFile != "" {
		err = os.WriteFile(cfg.OutputFile, []byte(content), 0o644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error writing output file")
			os.Exit(1)

		}
		return
	}
	fmt.Print(content)
}
