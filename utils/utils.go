package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
				fmt.Println("Invalid parameter", params["0"], "\nI got", a)
				os.Exit(1)
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
			fmt.Println("Specified Entrypoint does not exist!")
		} else {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	return dir
}
