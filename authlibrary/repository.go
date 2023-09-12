
package authlibrary

type RepoDetail struct {
    // Define fields for repository details
}

type Repository interface {
    CreateRepo(name string, private bool) (RepoDetail, error)
    ListRepos() ([]RepoDetail, error)
    // ... other repo-related methods
}
