package GO_API

import (
	"encoding/json"
	"Go_API/src/ettuparser"
	"sync"
	"Go_API/src/weatherparser"
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

func Ret() string {

	for {
		var data store
		wg.Add(1)
		go writeStation(&data)
		wg.Add(1)
		go writeWeather(&data)
		wg.Wait()
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		return string(jsonData)
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
