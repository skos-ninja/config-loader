package config

import (
	"encoding/json"
)

func setJSONConfig(configStr string, config interface{}) error {
	if configStr == "" {
		return nil
	}

	err := json.Unmarshal([]byte(configStr), config)
	if err != nil {
		return err
	}

	return nil
}
