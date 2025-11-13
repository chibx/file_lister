package utils

import (
	"errors"
	"fmt"
	"lister/structs"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

var SupportedFormats = []string{"text", "json"}

func PrintAndExit(error string, code int) {
	fmt.Println(error)
	os.Exit(code)
}

// Parses keys duhhhh
func ParseValues(args []string) map[string]string {
	params := make(map[string]string)

	for i := 0; i < len(args); i++ {
		curItem := args[i]
		if strings.HasPrefix(curItem, PREFIX) {
			if i == len(args)-1 {
				params[curItem] = ""
			} else {
				val := args[i+1]
				// Checks whether a dangling param was defined i.e a boolean param
				if strings.HasPrefix(val, PREFIX) {
					params[curItem] = ""
					continue
				}
				params[curItem] = val
				i++
				continue
			}
		} else {
			// This handles non-param values but only ensures there is one in this case
			if a, ok := params["0"]; ok {
				PrintAndExit(fmt.Sprint("Invalid parameter", params["0"], "\nI got", a), 1)
			}

			params["0"] = curItem
		}
	}

	return params
}

func MapHas[T comparable](bag map[T]any, key T) bool {
	_, ok := bag[key]

	return ok
}

func ParseToIntList(input string) ([]int, error) {
	list := strings.Split(input, ",")
	intList := make([]int, 1, 5)

	for _, val := range list {
		tmp, err := strconv.Atoi(val)
		if err != nil {
			return intList, fmt.Errorf("Found a non numeric value")
		}

		intList = append(intList, tmp)
	}

	return intList, nil
}

func CheckDir(name string) []os.DirEntry {
	dir, err := os.ReadDir(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			PrintAndExit("Specified Entrypoint does not exist!", 1)
		} else {
			PrintAndExit(err.Error(), 1)
		}
	}

	return dir
}

func parseToFileTree(config *structs.Config, curDir string, dir []os.DirEntry, depth int) *structs.FileTree {
	ft := new(structs.FileTree)
	ft.Name = filepath.Base(curDir)

	for _, data := range dir {
		if data.IsDir() {
			dirToEnter := filepath.Join(curDir, data.Name())
			_dir, err := os.ReadDir(dirToEnter)
			if err != nil {
				PrintAndExit(fmt.Sprint("Error occurred while reading into directory", dirToEnter), 1)
			}
			// time.Sleep(time.Second)
			if config.MaxDepth == -1 || depth <= config.MaxDepth {
				time.Sleep(time.Millisecond * time.Duration(config.Sleep))
				ft.Folders = append(ft.Folders, *parseToFileTree(config, dirToEnter, _dir, depth+1))
			}
		} else {
			if slices.Contains(config.IncludeFiles, -1) {
				ft.Files = append(ft.Files, data.Name())
			} else if slices.Contains(config.IncludeFiles, depth) {
				ft.Files = append(ft.Files, data.Name())
			}
		}
	}

	return ft
}
