package src

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ContainSymbol(tickers []string, ticker string) (int, bool) {
	for i, value := range tickers {
		if strings.EqualFold(strings.ToLower(value), strings.ToLower(ticker)) {
			return i, true
		}
	}
	return -1, false
}

func OpenConfigFile() Config {
	file, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	jsonConfig := Config{}
	json.Unmarshal(file, &jsonConfig)

	log.Println("Current symbols: ", strings.Join(jsonConfig.Tickers, ", "))

	return jsonConfig
}

var clear map[string]func()

func init() {
	clear = make(map[string]func()) 
	clear["linux"] = func() {
		cmd := exec.Command("clear") 
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") 
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {                         
		value()
	} else {
		log.Println("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
