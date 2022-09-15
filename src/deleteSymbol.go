package src

import (
	"os"

	"github.com/pterm/pterm"
)

func DeleteSymbol() string {
	jsonConfig := OpenConfigFile()

	ticker, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter symbol to remove ")

	if ticker == "" {
		return ""
	}

	index, ok := ContainSymbol(jsonConfig.Tickers, ticker)

	if !ok {
		return "Symbol does not exists."
	}

	jsonConfig.Tickers = remove(jsonConfig.Tickers, index)

	data, _ := json.Marshal(&jsonConfig)
	os.WriteFile("config.json", data, os.ModePerm)

	return "Symbol Removed."
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
