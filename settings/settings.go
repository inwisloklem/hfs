package settings

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Path          string
	ProfileNumber int
}

func Decode(tomlData string) (Config, error) {
	var config Config

	if _, err := toml.Decode(tomlData, &config); err != nil {
		return config, errors.New("can't decode toml")
	}
	return config, nil
}

func ReadConfigFile(filename string) (Config, error) {
	var config Config

	content, err := os.ReadFile(filename)
	if err != nil {
		return config, errors.New("can't read config file")
	}
	config, err = Decode(string(content))
	if err != nil {
		return config, err
	}
	return config, nil
}
