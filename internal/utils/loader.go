package loader

import (
	"encoding/json"
	"os"

	"github.com/gursimransw/prompt-analyzer/internal/types"
)

func LoadPatterns(fileName string) (*types.PatternConfig, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var config types.PatternConfig

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}
