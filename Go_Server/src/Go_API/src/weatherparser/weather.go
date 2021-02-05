package weatherparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const key = "90a74e7934e549966c49076196f53888"

type WeatherData struct {
	Temp string
	FeelsLike string
	Pressure string
}
func GetWeather() WeatherData {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Ekaterinburg&units=metric&appid=" + key)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]map[string]interface{}
	json.Unmarshal(body, &data)

	ret:=WeatherData{
		Temp: fmt.Sprintf("%v", data["main"]["temp"]),
		FeelsLike: fmt.Sprintf("%v", data["main"]["feels_like"]),
		Pressure: fmt.Sprintf("%v", data["main"]["pressure"]),
	}
	
	return ret
}
