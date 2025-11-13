package main

import (
	"lister/emitter"
	"lister/utils"
	"os"
)

func main() {
	trueArgs := os.Args[1:]
	// fmt.Println(trueArgs)
	parsedArg := utils.ParseValues(trueArgs)
	// fmt.Println(parsedArg)

	config := utils.CreateConfig(parsedArg)

	// fmt.Println(config)

	ftree := utils.StartScan(config)

	emitter.EmitOutput(config, ftree)
}
