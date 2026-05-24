package loader

//This package will be responsible for loading the prompt library json file into the detection logic.

import (
	"encoding/json"
	"os"

	"github.com/gursimransw/prompt-analyzer/internal/types"
)

func LoadDetectionRules(fileName string) (*[]types.DetectionRule, error) {

	//This functions simply takes file path as an argument, this is the filepath where the detection rule library is saved.
	//The function returns the pointer of array of DetectionRule struct , this struct is like a structure for a detection rule which contains both detection logic and metadata
	//(Refer internal/types/types.go for DetectionRule struct)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	//Reading the file for file contents (Detection Rules JSON File in this case)

	var detectionRules []types.DetectionRule
	//Variable declaration for storing JSON file content

	err = json.Unmarshal(data, &detectionRules)
	if err != nil {
		return nil, err
	}
	//Here we are loading the contents from the "data", which is a string of bytes and we are storing that data in the address location of
	//config variable of type (array of DetectionRule) , that we declared above.

	return &detectionRules, nil //We are returning pointer of detectionRule array here

}

func LoadPolicyConfig(fileName string) (*types.PolicyConfig, error) {

	//Just like above, this function loads the Policy Configuration into the program, this policy configuration
	//consists of both severity configuration and actions configuration. Here we are loading the whole policy into
	//Variable of type PolicyConfig

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	//Reading the file for file contents (Prompt Library JSON File in this case)

	var policyConfig types.PolicyConfig
	//Variable declaration for storing JSON file content

	err = json.Unmarshal(data, &policyConfig)
	if err != nil {
		return nil, err
	}
	return &policyConfig, nil //We are returning pointer of policyConfig here

}
