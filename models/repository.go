package models

import "encoding/json"

type Repository struct {
	Id           int64  `json:"id"`
	NodeId       string `json:"node_id"`
	HtmlUrl      string `json:"html_url"`
	ApiUrl       string `json:"url"`
	FullName     string `json:"full_name"`
	LanguagesUrl string `json:"languages_url"`
}

func DeserializeArrayString(rawjson string) ([]Repository, error) {
	repos := make([]Repository, 0)
	err := json.Unmarshal([]byte(rawjson), &repos)
	return repos, err
}

func DeserializeArrayBytes(rawjson []byte) ([]Repository, error) {
	repos := make([]Repository, 0)
	err := json.Unmarshal(rawjson, &repos)
	return repos, err
}

func (repository *Repository) Serialize() (string, error) {
	bytes, err := json.Marshal(*repository)
	return string(bytes), err
}
