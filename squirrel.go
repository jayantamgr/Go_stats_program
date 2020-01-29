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
		stats[values]["activityYesEventYes"] = 0
		stats[values]["activityYesEventNo"] = 0
		stats[values]["activityNoEventYes"] = 0
		stats[values]["activityNoEventNo"] = 0

		for items := 0; items < len(his); items++ {
			if his[items].Squirrel && Find(his[items].Events, values) {
				countYY++
				stats[values]["activityYesEventYes"] = countYY
			}
			if his[items].Squirrel && !Find(his[items].Events, values) {
				countYN++
				stats[values]["activityYesEventNo"] = countYN
			}
			if !his[items].Squirrel && Find(his[items].Events, values) {
				countNY++
				stats[values]["activityNoEventYes"] = countNY
			}
			if !his[items].Squirrel && !Find(his[items].Events, values) {
				countNN++
				stats[values]["activityNoEventNo"] = countNN
			}
		}
		countYY = 0
		countYN = 0
		countNY = 0
		countNN = 0
	}

	fmt.Println(stats)
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
