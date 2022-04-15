package common

import (
	"encoding/json"
	"os"
	"time"
)

const CONFIGURATION_FILE_PATH = "./config.json"
const CONFIGURATION_DEFAULT_PORT = 7788

type Configuration struct {
	Port          int `json:"port"`
	LastQueryTime int `json:"time"`
}

func (p *Configuration) Save() error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = os.WriteFile(CONFIGURATION_FILE_PATH, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetConfiguration() (*Configuration, error) {
	configuration := &Configuration{
		Port:          CONFIGURATION_DEFAULT_PORT,
		LastQueryTime: int(time.Now().UnixMilli()),
	}

	data, err := os.ReadFile(CONFIGURATION_FILE_PATH)
	if os.IsNotExist(err) {
		err = configuration.Save()
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		// Parse byte data.
		json.Unmarshal(data, configuration)
	}

	return configuration, nil
}
