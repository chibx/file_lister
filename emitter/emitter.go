package emitter

import (
	"fmt"
	"lister/emitter/json"
	"lister/emitter/text"
	"lister/structs"
	"lister/utils"
)

func EmitOutput(config *structs.Config, ftree *structs.FileTree) {
	format := config.DumpAs
	outFile := ""

	switch format {
	case "text":
		outFile = text.FileTreeToText(config, ftree)
	case "json":
		outFile = json.FileTreeToJSON(config, ftree)
	default:
		utils.PrintAndExit(fmt.Sprint("Unsupported format", format, "passed as argument to the -f flag"), 1)
	}

	fmt.Println("Data has been written to file with path:", outFile)
}
