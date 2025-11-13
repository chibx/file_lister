package utils

import (
	"fmt"
	"lister/structs"
)

func StartScan(config *structs.Config) {
	dir := CheckDir(config.EntryPoint)

	// for _, data := range dir {
	// 	fmt.Println(data.Name())
	// }

	ftree := parseToFileTree(config, config.EntryPoint, dir, 0)

	fmt.Println(ftree)
}
