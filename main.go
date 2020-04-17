package main

import (
	"fbreqs/cmd"
	"fbreqs/random"
	"log"
	"os"
)

func main() {
	// Exibe erros corretamente
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()

	// Init random (seed rand)
	random.Init()

	// Execute cobra command
	cmd.Execute()
}
