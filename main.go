package main

import (
	"fmt"
	"gitspeak/storage"
)

func main() {
	fmt.Println("Starting GitHub scraper...")
	conn, err := createDataBaseConnection()
	if err != nil {
		fmt.Printf("Error creating database connection! %s", err)
		return
	}
	lastId, err := conn.GetLastId()
	fmt.Printf("Got lastId from DB: %d", lastId)
	if err != nil {
		fmt.Printf("Error reading from database connection! %s", err)
		return
	}
	StartCrawlerFromId(lastId, conn)
}

func createDataBaseConnection() (*storage.DatabaseStorage, error) {
	settings := GetSettings()
	return storage.CreatePostgresConnection(settings.DbName, settings.DbPassword, settings.ServerName, settings.ServerPort)
}
