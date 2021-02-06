package downdetector

import (
	"encoding/json"
	"os/exec"
	"strings"
)

type status struct {
	Name string
	Up   bool
}

type statuses struct {
	List []status
}

func Downdetector(input []string) string {
	var data statuses

	for i := 0; i < len(input); i++ {
		data.List = append(data.List, ping(input[i]))
	}

	jsonData, _ := json.MarshalIndent(data, "", " ")

	return string(jsonData)
}

func ping(url string) status {
	var endpoint status
	out, _ := exec.Command("ping", url, "-c 1", "-w 5").Output()
	endpoint.Name = url
	if strings.Contains(string(out), "Destination Host Unreachable") {
		endpoint.Up = false
	} else {
		endpoint.Up = true
	}

	return endpoint
}