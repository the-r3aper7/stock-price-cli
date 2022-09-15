package main

import (
	"errors"
	"log"
	"os"

	"github.com/pterm/pterm"
	"github.com/the-r3aper7/stock-price-cli/src"
)

func main() {
	if _, err := os.Stat("config.json"); errors.Is(err, os.ErrNotExist) {
		os.Create("config.json")
	}

	var options []string
	options = append(options, "START Session")
	options = append(options, "ADD Symbol")
	options = append(options, "EDIT Symbol")
	options = append(options, "REMOVE Symbol")
	options = append(options, "Exit")

	for {
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

		switch selectedOption {
		case "START Session":
			src.TableRender()
		case "ADD Symbol":
			status := src.AddSymbol()
			log.Println(status)
		case "EDIT Symbol":
			status := src.EditSymbol()
			log.Println(status)
		case "REMOVE Symbol":
			status := src.DeleteSymbol()
			log.Println(status)
		case "Exit":
			log.Println("Exiting stock-price-cli...")
			os.Exit(0)
		}
	}
}
