package config

import (
	"io/ioutil"
)

func ReadFile(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return Config{}, err
	}

	return Read(data)
}
