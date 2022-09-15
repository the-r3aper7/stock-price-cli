package src

import (
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/pterm/pterm"
)

const FileName = "config.json"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	Tickers []string `json:"tickers"`
}

func AddSymbol() string {
	jsonConfig := OpenConfigFile()

	ticker, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("Add symbol to list ")
	pterm.Println()

	if ticker == "" {
		return ""
	}

	_, ok := ContainSymbol(jsonConfig.Tickers, ticker)
	if ok {
		return "Symbol already exists."
	}

	jsonConfig.Tickers = append(jsonConfig.Tickers, ticker)

	data, _ := json.Marshal(&jsonConfig)
	os.WriteFile("config.json", data, os.ModePerm)

	return "Symbol Added."
}
