package storage

import (
	"database/sql"
	"fmt"
	"gitspeak/storage/datatransfer"

	_ "github.com/lib/pq"
)

type DatabaseStorage struct {
	Connection *sql.DB
}

func CreatePostgresConnection(name, password, servername, port string) (*DatabaseStorage, error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=disable", name, password, servername, port)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &DatabaseStorage{
		Connection: db,
	}, nil
}

func (storage DatabaseStorage) Add(repo *datatransfer.RepoDataTransfer, languages *[]datatransfer.LanguageDataTransfer) error {
	_, err := storage.Connection.Exec("insert into repositories (id, url, languagesurl) values ($1, $2, $3)",
		repo.Id, repo.Url, repo.LanguagesUrl)
	if err != nil {
		return err
	}
	for _, lan := range *languages {
		_, err := storage.Connection.Exec("insert into languages (language, size, repositoryid) values ($1, $2, $3)",
			lan.Language, lan.Size, lan.RepositoryId)
		if err != nil {
			return err
		}
	}
	return err
}

func (storage *DatabaseStorage) GetLastId() (int64, error) {
	row := storage.Connection.QueryRow("SELECT coalesce(max(id), 0) FROM public.repositories")
	var index int64
	err := row.Scan(&index)
	return index, err
}
