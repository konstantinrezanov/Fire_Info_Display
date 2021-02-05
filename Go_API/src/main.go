package main

import (
	"encoding/json"
	"ettuparser"
	"io/ioutil"
	"sync"
	"time"
	"weatherparser"
	//"log"
)

type store struct {
	Stations []station
	Weather  weatherparser.WeatherData
}

type station struct {
	ID     string
	Status ettuparser.Status
}

var wg sync.WaitGroup

func main() {

	var data store
	for {
		wg.Add(1)
		go writeStation(&data)
		wg.Add(1)
		go writeWeather(&data)
		wg.Wait()
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		_ = ioutil.WriteFile("/home/rezanov_konstantin/Documents/Programming/Projects/Fire_Info_Display/Go_Server/src/data/data.json", jsonData, 0644)
		time.Sleep(1*time.Minute)
	}
}

func writeWeather(data *store) {
	defer wg.Done()
	data.Weather = weatherparser.GetWeather()
}

func writeStation(data *store) {
	defer wg.Done()
	pstations := []string{"3442", "962313"}
	for _, n := range pstations {
		data.Stations = append(data.Stations, station{
			ID:     n,
			Status: ettuparser.ParseStation(n),
		})
	}
}
