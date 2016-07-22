package unpackage

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name         string
	Version      string
	Source       []string
	License      string
	Description  string
	Conflicts    []string
	Requires     []string
	Setenv       map[string]string
	Unsetenv     []string
	Append_path  map[string][]string
	Prepend_path map[string][]string
	Setup        []string
	Build        []string
	Install      []string
	Load         []string
	Unload       []string
}

func Parse(b []byte) *Config {
	c := Config{}
	err := yaml.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return &c
}
