package storage

import "gitspeak/storage/datatransfer"

type StorageWriter interface {
	Add(repo *datatransfer.RepoDataTransfer, languages *[]datatransfer.LanguageDataTransfer) error
}
