package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fs:=http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	
	http.HandleFunc("/data",dataServer)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func dataServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	b, err := ioutil.ReadFile("./data/data.json")

	if err != nil {
		log.Fatal(err)
	}

	raw := json.RawMessage(string(b))
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(raw, &objmap)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(objmap)
}
