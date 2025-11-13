package text

import "lister/structs"

func FileTreeToText(config *structs.Config, tree *structs.FileTree) string {
	outFile := config.Output + ".txt"

	return outFile
}
