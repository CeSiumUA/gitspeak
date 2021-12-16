create table languages(
    Id bigserial primary key,
    Language text,
    Size bigint,
    RepositoryId bigint references repositories (Id)
)