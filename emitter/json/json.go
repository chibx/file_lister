package json

import (
	"encoding/json"
	"fmt"
	"lister/structs"
	"lister/utils"
	"os"
)

func FileTreeToJSON(config *structs.Config, tree *structs.FileTree) string {
	outFile := config.Output + ".json"

	file, err := os.Create(outFile)
	if err != nil {
		utils.PrintAndExit(fmt.Sprint("Error:", err.Error()), 1)
	}

	enc := json.NewEncoder(file)
	enc.SetEscapeHTML(false)
	err = enc.Encode(tree)

	if err != nil {
		utils.PrintAndExit("Error writing output to "+outFile, 1)
	}

	file.Close()

	return outFile
}
