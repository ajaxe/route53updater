package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadConfiguration(filepath string) (*config, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s, %v", filepath, err)
	}
	cfg := config{}
	err = json.Unmarshal(b, &cfg)

	return &cfg, nil
}
