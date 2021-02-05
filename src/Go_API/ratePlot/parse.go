package ratePlot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type RateData struct {
	Days   []string
	Values []float64
}

type rateResponse struct {
	Date string 
	Rates map[string]float64 `json:"rates"`
}

func ParseRates(base, symbol string) RateData {
	var data RateData
	now := time.Now()
	url := fmt.Sprintf("https://api.exchangeratesapi.io/%v-%v-%v?base=%s&symbols=%s", 
		int(now.Year()), int(now.Month()), int(now.Day()),base, symbol)
	parse(url,symbol, &data)
	for i := 0; i < 6; i++ {
		now=now.AddDate(0,0,-1)
		url := fmt.Sprintf("https://api.exchangeratesapi.io/%v-%v-%v?base=%s&symbols=%s", 
		int(now.Year()), int(now.Month()), int(now.Day()),base, symbol)
		parse(url,symbol, &data)
	}

	return data
}

func parse(url,symbol string, data *RateData) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	var response rateResponse

	json.Unmarshal(body, &response)
	data.Days=append(data.Days,response.Date)
	data.Values = append(data.Values, response.Rates[symbol])
}
