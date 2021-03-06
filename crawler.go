package main

import (
	"fmt"
	"gitspeak/models"
	"gitspeak/storage"
	"net/http"
	"time"
)

func StartCrawler(writer storage.StorageWriter) {
	StartCrawlerFromId(0, writer)
}

func StartCrawlerFromId(startId int64, writer storage.StorageWriter) {
	httpClient := createhttpClient()
	for {
		request, err := createRepoRequest(startId)
		if err != nil {
			fmt.Printf("Error creating request: %s", err)
		}
		response, err := httpClient.Do(request)
		remains := response.Header.Get("X-RateLimit-Remaining")
		fmt.Printf("Requests remaining: %s \n", remains)
		if err != nil {
			fmt.Printf("Error doing request: %s", err)
			continue
		}
		if response.StatusCode == http.StatusForbidden {
			time.Sleep(10 * time.Minute)
			continue
		}
		repos, err := models.DeserializeArrayReader(&response.Body)
		if err != nil {
			fmt.Printf("Error reading body! %s", err)
		}
		response.Body.Close()
		lastGreatestId := startId
		for _, repo := range repos {
			languages := getLanguages(httpClient, repo.LanguagesUrl)
			if repo.Id > lastGreatestId {
				lastGreatestId = repo.Id
			}
			saveData(&repo, languages, writer)
		}
		startId = lastGreatestId
		if len(repos) != 100 {
			return
		}
	}
}

func getLanguages(httpClient *http.Client, languageUrl string) *models.RepoLanguagesSet {
	request, err := createLanguagesRequest(languageUrl)
	if err != nil {
		fmt.Printf("Error creating languages request! %s", err)
	}
	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Printf("Error doing language request: %s \n", err)
	}
	remains := response.Header.Get("X-RateLimit-Remaining")
	fmt.Printf("Requests remaining: %s \n", remains)
	if response.StatusCode == http.StatusForbidden {
		errorMessage, err := models.DeserializeErrorResponseFromBody(&response.Body)
		if err == nil && errorMessage.Message == "Repository access blocked" {
			langSet := make(models.RepoLanguagesSet, 0)
			fmt.Println("Repository access blocked")
			return &langSet
		}
		time.Sleep(10 * time.Minute)
		return getLanguages(httpClient, languageUrl)
	}
	languages, err := models.DeserializeLanguagesFromReader(&response.Body)
	if err != nil {
		fmt.Printf("Error getting languages from body! %s", err)
	}
	response.Body.Close()
	return languages
}

func createhttpClient() *http.Client {
	client := &http.Client{}
	return client
}

func saveData(repo *models.Repository, languages *models.RepoLanguagesSet, writer storage.StorageWriter) {
	repoDto := repo.ToDto()
	langDto := languages.ToDto(repo.Id)
	err := writer.Add(repoDto, langDto)
	if err != nil {
		fmt.Printf("Error adding data to repository! %s", err)
	}
}

func createRepoRequest(id int64) (*http.Request, error) {
	url := fmt.Sprintf("https://api.github.com/repositories?since=%d", id)
	token := GetToken()
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	return req, nil
}

func createLanguagesRequest(url string) (*http.Request, error) {
	token := GetToken()
	req, err := http.NewRequest(
		"GET", url, nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	return req, nil
}
