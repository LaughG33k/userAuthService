package config

import (
	"io/ioutil"

	"github.com/goccy/go-json"
)

func Load() (AppConfig, error) {

	appCfg := map[string]AppConfig{"app": AppConfig{}}

	bytes, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return AppConfig{}, err
	}

	if err := json.Unmarshal(bytes, &appCfg); err != nil {
		return AppConfig{}, err
	}

	return appCfg["app"], nil

}
