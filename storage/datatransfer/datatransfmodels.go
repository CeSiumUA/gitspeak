package datatransfer

type RepoDataTransfer struct {
	Id           int64
	Url          string
	LanguagesUrl string
}

type LanguageDataTransfer struct {
	Id           int64
	Language     string
	Size         int64
	RepositoryId int64
}
