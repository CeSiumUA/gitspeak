package models

import "encoding/json"

type RepoLanguagesSet map[string]int64

func DeserializeLanguagesFromStr(jsonStr string) (*RepoLanguagesSet, error) {
	langsSet := make(RepoLanguagesSet, 0)
	err := json.Unmarshal([]byte(jsonStr), &langsSet)
	return &langsSet, err
}

func DeserializeLanguagesFromBytes(jsonBytes []byte) (*RepoLanguagesSet, error) {
	langsSet := make(RepoLanguagesSet, 0)
	err := json.Unmarshal(jsonBytes, &langsSet)
	return &langsSet, err
}
