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
	ServerName     string
	ServerPort     string
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
	err := preCheck()
	if err != nil {
		return ""
	}
	return loadedSettings.GitHubApiToken
}

func GetUserName() string {
	err := preCheck()
	if err != nil {
		return ""
	}
	return loadedSettings.DbName
}

func GetUserPassword() string {
	err := preCheck()
	if err != nil {
		return ""
	}
	return loadedSettings.DbPassword
}

func GetServerName() string {
	err := preCheck()
	if err != nil {
		return ""
	}
	return loadedSettings.ServerName
}

func GetServerPort() string {
	err := preCheck()
	if err != nil {
		return ""
	}
	return loadedSettings.ServerPort
}

func GetSettings() *Settings {
	err := preCheck()
	if err != nil {
		return nil
	}
	return loadedSettings
}

func preCheck() error {
	if loadedSettings == nil {
		err := loadSettings()
		if err != nil {
			fmt.Printf("Error while loading settings: %s", err)
			return err
		}
	}
	return nil
}
