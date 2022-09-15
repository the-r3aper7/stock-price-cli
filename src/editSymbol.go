package src

import (
	"os"

	"github.com/pterm/pterm"
)

func EditSymbol() string {
	jsonConfig := OpenConfigFile()

	ticker, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter symbol for editing ")

	if ticker == "" {
		return ""
	}

	index, ok := ContainSymbol(jsonConfig.Tickers, ticker)

	if !ok {
		return "Symbol does not exists."
	}

	updatedTicker, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Enter symbol to update ")
	pterm.Println()

	jsonConfig.Tickers[index] = updatedTicker

	data, _ := json.Marshal(&jsonConfig)
	os.WriteFile("config.json", data, os.ModePerm)

	return "Symbol Updated."
}
