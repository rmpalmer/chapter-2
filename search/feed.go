package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	//open the file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// schedule the file to be closed
	defer file.Close()

	// decode the file into a slice of pointers
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
