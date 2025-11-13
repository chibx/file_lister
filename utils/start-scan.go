package utils

import "fmt"

func StartScan(config *Config) {
	dir := CheckDir(config.EntryPoint)

	for _, data := range dir {
		fmt.Println(data.Name())
	}
}
