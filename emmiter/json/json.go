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
	data, err := json.Marshal(tree)

	if err != nil {
		utils.PrintAndExit("Error outputting to json format", 1)
	}

	file, err := os.Create(outFile)
	if err != nil {
		utils.PrintAndExit(fmt.Sprint("Error:", err.Error()), 1)
	}

	_, err = file.Write(data)
	if err != nil {
		// fmt.Print
		utils.PrintAndExit("Error writing output to "+outFile, 1)
	}

	file.Close()

	return outFile
}
