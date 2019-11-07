package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config is the required data for program to run
type Config struct {
	bytes []byte
	SMTP struct {
		Server     string
		Port       int
		Sender     string
		Recipients []string
	}
	Directory struct {
		Name string
	}
}

// New gets ./config.yaml from local dir and parses yaml to Config type
func New() (*Config, error) {
	bytes, err := loadFile();
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	conf, err := parseYamlToConfig(bytes)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conf, nil
}

func loadFile() ([]byte, error) {
	f, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return f, nil
}

// Decodes yaml bytes to Config struct
func parseYamlToConfig(data []byte) (*Config, error) {
	var conf Config
	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &conf, nil
}
