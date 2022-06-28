package assets

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	Nodes        = LoadDefault("assets/nodes.json")
	Providers    = LoadDefault("assets/providers.json")
	RefinedNodes = LoadDefault("assets/refined_nodes.json")
)

func LoadDefault(file string) *DefaultJson { // Load a file using DefaultJson as the type
	jsonFile, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var response DefaultJson

	json.Unmarshal(byteValue, &response)
	defer jsonFile.Close()

	return &response
}
