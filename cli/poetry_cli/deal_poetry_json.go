package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func ParseJsonPoe(fileLoc string) string {

	// Construct Poetry struct for json unmarshal
	type Poetry struct {
		Author string
		Paragraphs []string
		Rhythmic string
		Title string
	}

	jsonData, err := ioutil.ReadFile(fileLoc)
	ChkErr(err)

	jsonByte := []byte(string(jsonData))

	var jsonSlice []Poetry

	err = json.Unmarshal(jsonByte, &jsonSlice)
	ChkErr(err)

	// Generate random number for picking randomly
	seed := rand.NewSource(time.Now().UnixNano())
	rand0999 := rand.New(seed).Intn(1000)

	// Re-rand while there aren't enough element in jsonSlice, (such as the last one file).
	for rand0999 >= len(jsonSlice) {
		rand0999 = rand.New(seed).Intn(1000)
	}
	randPoe := jsonSlice[rand0999]

	// Declare a slice to store poetry string
	poeSlice := []string {}

	// Build poetry slice
	poeSlice = append(poeSlice, "\n")

	// Add title for poetry, rhythmic for songci
	poeSlice = append(poeSlice, randPoe.Title)
	poeSlice = append(poeSlice, randPoe.Rhythmic)
	poeSlice = append(poeSlice, " - ")
	poeSlice = append(poeSlice, randPoe.Author + "    ")
	poeSlice = append(poeSlice, "\n")
	for ord := range randPoe.Paragraphs {
		poeSlice = append(poeSlice, "\n")
		poeSlice = append(poeSlice, randPoe.Paragraphs[ord])
		poeSlice = append(poeSlice, "\n")
	}

	// Turn poetry slice to string
	poeStr := strings.Join(poeSlice, "")

	return poeStr
}
