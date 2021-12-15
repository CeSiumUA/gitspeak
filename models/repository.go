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

func DeserializeArrayString(rawjson string) []Repository {
	repos := make([]Repository, 0)
	json.Unmarshal([]byte(rawjson), &repos)
	return repos
}

func DeserializeArrayBytes(rawjson []byte) []Repository {
	repos := make([]Repository, 0)
	json.Unmarshal(rawjson, &repos)
	return repos
}

func (repository *Repository) Serialize() (string, error) {
	bytes, err := json.Marshal(*repository)
	return string(bytes), err
}
