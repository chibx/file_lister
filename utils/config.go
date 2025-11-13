package utils

import (
	"fmt"
	"lister/structs"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Args = map[string]string

func getcwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		PrintAndExit("Error occurred while getting working directory", 1)
	}

	return cwd
}

func CreateConfig(parsedArg Args) *structs.Config {
	var cwd = getcwd()
	var output = parseOutput(parsedArg)
	var sleep = parseDelay(parsedArg)
	var includeFiles = parseIncludes(parsedArg)
	var entryName = parseEntry(parsedArg)
	var depth = parseDepth(parsedArg)
	var dumpAs = parseFormat(parsedArg)

	return &structs.Config{
		Output:       output, // -o
		Sleep:        sleep,  // --delay
		EntryPoint:   entryName,
		IncludeFiles: includeFiles, // --fl
		DumpAs:       dumpAs,       // -f
		MaxDepth:     depth,        // -d
		Cwd:          cwd,
	}
}

func parseFormat(parsedArg Args) string {
	var dumpAs = defaultOutFormat

	if format, ok := parsedArg["-f"]; ok {
		if slices.Contains(SupportedFormats, format) {
			dumpAs = format
		}
	}

	return dumpAs
}

func parseOutput(parsedArg Args) string {
	var output string = defaultOutput
	if outFile, ok := parsedArg["-o"]; ok {
		if !filepath.IsAbs(outFile) {
			outFile = filepath.Join(getcwd(), outFile)
		}
		output = filepath.Clean(outFile)
	}

	return output
}

func parseDelay(parsedArg Args) int {
	var sleep = defaultSleep
	if del, ok := parsedArg["--delay"]; ok {
		delay, err := strconv.Atoi(del)
		if err != nil {
			fmt.Printf("Expected a number for the delay --delay! Defaulting to %d", defaultSleep)
		}

		sleep = int(math.Max(float64(delay), 0))
	}

	return sleep
}

func parseIncludes(parsedArg Args) []int {
	var includeFiles = defaultFileLevel

	if _inc, ok := parsedArg["--fl"]; ok {
		include, err := ParseToIntList(_inc)

		if err != nil {
			fmt.Println("Error:", err.Error(),
				"due to issue with input provided to --fl, I got", _inc, "instead\n",
				"Continuing with default...")

			include = make([]int, 0) // Empty slice of levels to include
		}

		if slices.Contains(include, -1) && len(include) > 1 {
			include = defaultFileLevel // Avoid mixing -1 with other levels
		}

		includeFiles = include

	}

	return includeFiles
}

func parseEntry(parsedArg Args) string {
	var entryName = ""

	if e, ok := parsedArg["0"]; ok {
		entryName = filepath.Clean(strings.TrimSpace(e))

		if !filepath.IsAbs(entryName) {
			entryName = filepath.Join(getcwd(), entryName)
		}

		if len(entryName) == 0 {
			PrintAndExit("Invalid filepath suggested as the entry", 1)
		}
	} else {
		PrintAndExit("No Entry Folder Specified!", 1)
	}

	return entryName
}

func parseDepth(parsedArg Args) int {
	var depth = defaultDepth

	if _dep, ok := parsedArg["-d"]; ok {
		dep, err := strconv.Atoi(_dep)
		if err != nil {
			fmt.Printf("Expected a number for the depth -d! Defaulting to %d", defaultSleep)
		}

		if dep < -1 {
			PrintAndExit("Accepted -d (Depth) values must be from -1 up", 1)
		}

		depth = dep
	}

	return depth
}
