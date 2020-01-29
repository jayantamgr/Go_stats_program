package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open the json data
	dataToParse, err := os.Open("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users. json")
	defer dataToParse.Close()

	byteValue, _ := ioutil.ReadAll(dataToParse)
	type History struct {
		Squirrel bool
		Events   []string
	}
	var his []History
	err = json.Unmarshal(byteValue, &his)
	if err != nil {
		fmt.Println(err)
	}

	var activities []string

	for i := 0; i < len(his); i++ {
		events := his[i].Events
		for j := 0; j < len(events); j++ {
			found := Find(activities, events[j])
			if !found {
				activities = append(activities, events[j])
			}

		}
	}

	var stats = make(map[string]map[string]int)

	countYY := 0
	countYN := 0
	countNY := 0
	countNN := 0

	for _, values := range activities {
		stats[values] = make(map[string]int)
		stats[values]["squirrel Yes "+values+" Yes"] = 0
		stats[values]["squirrel Yes "+values+" No"] = 0
		stats[values]["squirrel No "+values+" Yes"] = 0
		stats[values]["squirrel No "+values+" No"] = 0

		for items := 0; items < len(his); items++ {
			if his[items].Squirrel && Find(his[items].Events, values) {
				countYY++
				stats[values]["squirrel Yes "+values+" Yes"] = countYY
			}
			if his[items].Squirrel && !Find(his[items].Events, values) {
				countYN++
				stats[values]["squirrel Yes "+values+" No"] = countYN
			}
			if !his[items].Squirrel && Find(his[items].Events, values) {
				countNY++
				stats[values]["squirrel No "+values+" Yes"] = countNY
			}
			if !his[items].Squirrel && !Find(his[items].Events, values) {
				countNN++
				stats[values]["squirrel No "+values+" No"] = countNN
			}
		}
		countYY = 0
		countYN = 0
		countNY = 0
		countNN = 0
	}

	var result = make(map[string]int)

	for key, val := range stats {
		result[key] = 0
		fmt.Println(key, val)
	}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
