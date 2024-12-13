package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

func ReadConfig() Server {
	var server Server
	filePath := getFilePathByOperatingSystem()
	checkFileAndCreateIfNotExist(filePath, &server)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error opening YAML file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&server); err != nil {
		log.Fatalf("error decoding YAML: %v", err)
	}

	return server
}

func checkFileAndCreateIfNotExist(filePath string, server *Server) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("config.yml not found! Creating a new one...")
		server.Name = "Cli-Chat"
		server.IP = "0.0.0.0"
		server.Port = 5000

		yamlData, err := yaml.Marshal(&server)
		if err != nil {
			log.Fatalf("erro creating YAML: %v", err)
		}

		err = os.WriteFile(filePath, yamlData, 0644)
		if err != nil {
			log.Fatalf("error writing YAML file: %v in %s", err, filePath)
		}
		log.Println("config.yml created in ", filePath, " successfully!")
	}
}

func getFilePathByOperatingSystem() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error getting home directory: %v", err)
	}

	var filePath string
	if runtime.GOOS == "windows" {
		filePath = filepath.Join(homeDir, "AppData", "Local", "Cli-Chat", "config.yml")
	} else {
		filePath = filepath.Join(homeDir, ".Cli-Chat", "config.yml")
	}

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		log.Fatalf("error creating config directory: %v", err)
	}

	return filePath
}
