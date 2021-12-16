package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	GitHubApiToken string
	DbName         string
	DbPassword     string
}

var loadedSettings *Settings

func loadSettings() error {
	fileBytes, err := os.ReadFile("settings.json")
	if err != nil {
		return err
	}
	loadedSettings = &Settings{}
	err = json.Unmarshal(fileBytes, loadedSettings)
	return err
}

func GetToken() string {
	if loadedSettings == nil {
		err := loadSettings()
		if err != nil {
			fmt.Printf("Error while loading settings: %s", err)
			return ""
		}
	}
	return loadedSettings.GitHubApiToken
}
