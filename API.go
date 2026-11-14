package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type DailyData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type Ticker struct {
	MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		OutputSize    string `json:"4. Output Size"`
		TimeZone      string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeriesDaily map[string]DailyData `json:"Time Series (Daily)"`
}

func cmdUseTicker(state *GameState) {
	scanner := bufio.NewScanner(os.Stdin)
	oldScreen()
	fmt.Print("***")
	scanner.Scan()
	symbol := cleanInput(scanner.Text())
	ticker, err := shareTicker(symbol)
	if err != nil {
		fmt.Printf("Could not find that share: %s", err)
	}
	if len(ticker.TimeSeriesDaily[ticker.MetaData.LastRefreshed].Close) == 0 {
		fmt.Println("Could not find share.")
	} else {
		symbolString := strings.ToUpper(symbol)
		fmt.Printf("The current share price of %s is %s.\n", symbolString, ticker.TimeSeriesDaily[ticker.MetaData.LastRefreshed].Close)
	}
}

func shareTicker(symbol string) (Ticker, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=WZTB9USXA7ZPYNG0", symbol)
	res, err := http.Get(url)
	if err != nil {
		return Ticker{}, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Ticker{}, err
	}
	var ticker Ticker
	if err := json.Unmarshal(data, &ticker); err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}
