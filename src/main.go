package main

import (
	"Go_API/ettuparser"
	"Go_API/ratePlot"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)

	http.HandleFunc("/data/station", stationServer)
	http.HandleFunc("/data/rate", rateServer)
	log.Fatal(http.ListenAndServe("192.168.1.240:8081", nil))
}

func stationServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	raw := json.RawMessage(string(ettuparser.StationSend()))
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(raw, &objmap)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(objmap)
}

func rateServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	rates := ratePlot.ParseRates(r.FormValue("base"), r.FormValue("symbol"))
	fmt.Fprintln(w, ratePlot.Plot(rates))
}
