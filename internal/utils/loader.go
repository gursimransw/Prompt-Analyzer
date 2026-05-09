package loader

//This package will be responsible for loading the prompt library json file into the detection logic.

import (
	"encoding/json"
	"os"

	"github.com/gursimransw/prompt-analyzer/internal/types"
)

func LoadPatterns(fileName string) (*types.PatternConfig, error) {

	//This functions simply takes file path as an argument, this is the filepath where the prompt library is saved.
	//THe function returns the pointer of PatterConfig struct , this struct is like a collection disctionary for all malicious prompt categories and their
	//respective patterns for detection. (Refer internal/types/types.go for PatternConfig struct)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	//Reading the file for file contents (Prompt Library JSON File in this case)

	var config types.PatternConfig
	//Variable declaration for storing JSON file content

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	//Here we are loading the contents from the "data", which is a string of bytes and we are storing that data in the address location of
	//config variable of type PatternConfig , that we declared above.

	return &config, nil //We are returning pointer of config here

}
