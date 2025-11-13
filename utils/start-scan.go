package utils

import (
	"lister/structs"
)

func StartScan(config *structs.Config) *structs.FileTree {
	dir := CheckDir(config.EntryPoint)

	ftree := parseToFileTree(config, config.EntryPoint, dir, 0)

	// fmt.Println(ftree)

	return ftree
}
