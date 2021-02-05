package GO_API

import (
	"Go_API/src/ettuparser"
	"encoding/json"
	//"log"
)

type station struct {
	ID     string
	Status ettuparser.Status
}
type stations struct {
	List []station
}

func StationSend() string {
	var data stations
	pstations := []string{"3442", "962313"}
	for _, n := range pstations {
		data.List= append(data.List, station{
			ID:     n,
			Status: ettuparser.ParseStation(n),
		})
	}
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonData)
}