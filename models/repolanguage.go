package models

import (
	"encoding/json"
	"gitspeak/storage/datatransfer"
	"io"
)

type RepoLanguagesSet map[string]int64

func DeserializeLanguagesFromStr(jsonStr string) (*RepoLanguagesSet, error) {
	langsSet := make(RepoLanguagesSet, 0)
	err := json.Unmarshal([]byte(jsonStr), &langsSet)
	return &langsSet, err
}

func DeserializeLanguagesFromReader(reader *io.ReadCloser) (*RepoLanguagesSet, error) {
	langsSet := make(RepoLanguagesSet, 0)
	decoder := json.NewDecoder(*reader)
	err := decoder.Decode(&langsSet)
	return &langsSet, err
}

func DeserializeLanguagesFromBytes(jsonBytes []byte) (*RepoLanguagesSet, error) {
	langsSet := make(RepoLanguagesSet, 0)
	err := json.Unmarshal(jsonBytes, &langsSet)
	return &langsSet, err
}

func (langSet *RepoLanguagesSet) ToDto(repositoryId int64) *[]datatransfer.LanguageDataTransfer {
	dtos := make([]datatransfer.LanguageDataTransfer, 0)
	for index, size := range *langSet {
		dto := datatransfer.LanguageDataTransfer{
			Language:     index,
			Size:         size,
			RepositoryId: repositoryId,
		}
		dtos = append(dtos, dto)
	}
	return &dtos
}
