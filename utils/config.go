package utils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Config struct {
	Sleep      int
	Output     string // Without extension
	EntryPoint string
	// If files should be included, -1 = all levels, 0 = none, 1 = depth one, and so on....
	//
	// This depends on the depth level of course
	IncludeFiles []int
	// How many folders deep
	//
	// Defaults to -1 for all. 0 returns nothing or just files if the IncludeFiles option is true
	Depth  int
	DumpAs string
}

func CreateConfig(parsedArg map[string]string) *Config {
	var output string = defaultOutput
	var sleep = defaultSleep
	var includeFiles = defaultFileLevel
	var entryName = ""
	var dumpAs = defaultOutFormat
	var depth = defaultDepth

	if outFile, ok := parsedArg["-o"]; ok {
		output = outFile
	}

	if del, ok := parsedArg["--delay"]; ok {
		delay, err := strconv.Atoi(del)
		if err != nil {
			fmt.Printf("Expected a number for the delay --delay! Defaulting to %d", defaultSleep)
		}

		sleep = int(math.Max(float64(delay), 0))
	}

	if e, ok := parsedArg["0"]; ok {
		entryName = filepath.Clean(strings.TrimSpace(e))
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error occurred while accessing working directory")
			os.Exit(1)
		}

		if !filepath.IsAbs(entryName) {
			entryName = filepath.Join(cwd, entryName)
		}

		if len(entryName) == 0 {
			fmt.Println("Invalid filepath suggested as the entry")
			os.Exit(1)
		}
	} else {
		fmt.Println("No Entry Folder Specified!")
		os.Exit(1)
	}

	if _inc, ok := parsedArg["--include-files"]; ok {
		include, err := ParseToIntList(_inc)

		if err != nil {
			fmt.Println("Error:", err.Error(), "due to issue with input provided to --include-files, I got", _inc, "instead")
		}

		if slices.Contains(include, 0) {
			include = defaultFileLevel
		}

		includeFiles = include

	}

	if _dep, ok := parsedArg["-d"]; ok {
		dep, err := strconv.Atoi(_dep)
		if err != nil {
			fmt.Printf("Expected a number for the depth -d! Defaulting to %d", defaultSleep)
		}

		if dep < -1 {
			fmt.Println("Accepted -d (Depth) values must be from -1 up")
		}

		depth = dep
	}

	return &Config{
		Output:       output,
		Sleep:        sleep,
		EntryPoint:   entryName,
		IncludeFiles: includeFiles,
		DumpAs:       dumpAs,
		Depth:        depth,
	}
}
