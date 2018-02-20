package config

import(
  "errors"

  "gopkg.in/yaml.v2"
)

func Read(data []byte) (Config, error) {
  config := Config{}
  err := yaml.Unmarshal(data, &config)

  if err != nil {
    return config, errors.New("Could not parse web-deployer.yml")
  }

  return config, nil
}
