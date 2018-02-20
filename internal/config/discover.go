package config

import(
  "path"
)

func Discover(dir string) (Config, error) {
  return ReadFile(path.Join(dir, "web-deployer.yml"))
}
