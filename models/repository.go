package models

import (
	"encoding/json"
	"gitspeak/storage/datatransfer"
	"io"
)

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

func DeserializeArrayReader(reader *io.ReadCloser) ([]Repository, error) {
	repos := make([]Repository, 0)
	decoder := json.NewDecoder(*reader)
	err := decoder.Decode(&repos)
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

func (repo *Repository) ToDto() *datatransfer.RepoDataTransfer {
	return &datatransfer.RepoDataTransfer{
		Id:           repo.Id,
		Url:          repo.ApiUrl,
		LanguagesUrl: repo.LanguagesUrl,
	}
}
