package src

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

var boldWithBgBlack = &tablewriter.Colors{tablewriter.BgBlackColor, tablewriter.Bold}
var boldWithBgGreen = &tablewriter.Colors{tablewriter.BgGreenColor, tablewriter.Bold}
var boldWithBgGRed = &tablewriter.Colors{tablewriter.BgRedColor, tablewriter.Bold}
var bold = &tablewriter.Colors{tablewriter.Bold}

func TableRender() {
	var start time.Time

	jsonConfig := OpenConfigFile()
	tickers := jsonConfig.Tickers

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Ticker", "Currency", "Price", "Change", "% Change", "Volume"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeaderColor(*boldWithBgBlack, *boldWithBgBlack, *boldWithBgBlack, *boldWithBgBlack, *boldWithBgBlack, *boldWithBgBlack)

	for {
		start = time.Now()
		jsonData := GetData(tickers)
		table.SetCaption(true, fmt.Sprintf("fetched data in %.2fs, last updated %v, Ctrl + C to exit", time.Since(start).Seconds(), time.Now().Format("15:04:05")))
		for _, data := range *jsonData {
			if data.ErrDescription != "" {
				colData := []string{"", "", fmt.Sprintf("Symbol not found: %s", data.Data.Symbol), "", "", ""}
				table.Rich(colData, []tablewriter.Colors{{}, {}, {tablewriter.FgRedColor, tablewriter.Bold}, {}, {}, {}})
			} else if data.Data.Change > 0 {
				// Green Color
				colData := []string{data.Data.Symbol, data.Data.Currency, strconv.FormatFloat(data.Data.Price, 'f', -1, 64), strconv.FormatFloat(data.Data.Change, 'f', -1, 64), strconv.FormatFloat(data.Data.PerChange, 'f', -1, 64), strconv.Itoa(data.Data.Volume)}
				table.Rich(colData, []tablewriter.Colors{*bold, *bold, *boldWithBgGreen, *boldWithBgGreen, *boldWithBgGreen, *bold})
			} else if data.Data.Change < 0 {
				// Red Color
				colData := []string{data.Data.Symbol, data.Data.Currency, strconv.FormatFloat(data.Data.Price, 'f', -1, 64), strconv.FormatFloat(data.Data.Change, 'f', -1, 64), strconv.FormatFloat(data.Data.PerChange, 'f', -1, 64), strconv.Itoa(data.Data.Volume)}
				table.Rich(colData, []tablewriter.Colors{*bold, *bold, *boldWithBgGRed, *boldWithBgGRed, *boldWithBgGRed, *bold})
			} else {
				//Grey Color
				colData := []string{data.Data.Symbol, data.Data.Currency, strconv.FormatFloat(data.Data.Price, 'f', -1, 64), strconv.FormatFloat(data.Data.Change, 'f', -1, 64), strconv.FormatFloat(data.Data.PerChange, 'f', -1, 64), strconv.Itoa(data.Data.Volume)}
				table.Rich(colData, []tablewriter.Colors{})
			}
		}

		ClearScreen()
		table.Render()

		table.ClearRows()

		time.Sleep(10 * time.Second)
	}
}
