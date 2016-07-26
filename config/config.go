package config

import (
	"gopkg.in/yaml.v2"
	"io"
	"errors"
	"io/ioutil"
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

func Parse(b []byte) (c *Config, err error) {
	c = &Config{}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		err = errors.New("Could not parse config, reason: " + err.Error())
	}
	return
}

func ReadConfig(r *io.Reader) (c *Config, err error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		err = errors.New("Could not read config, reason: " + err.Error())
		return
	}
	c, err = Parse(raw)
	return
}