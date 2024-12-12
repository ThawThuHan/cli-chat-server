package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
	Port   int    `yaml:"port"`
}

func ReadConfig() Server {
	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("error opening YAML file: %v", err)
	}
	defer file.Close()

	var server Server
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&server); err != nil {
		log.Fatalf("error decoding YAML: %v", err)
	}

	return server
}
