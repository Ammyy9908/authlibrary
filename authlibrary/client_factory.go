package authlibrary

import (
	"errors"

	"github.com/ammyy9908/authlibrary/interfaces"
	"github.com/ammyy9908/authlibrary/providers"
)

type ClientFactory struct{}

func (cf *ClientFactory) CreateClient(provider string, token string) (Client, error) {
	var client Client
	if provider == "github" {
		client = &providers.GitHubClient{Token: token}
	} else if provider == "gitlab" {
		client = &providers.GitLabClient{Token: token}
	} else {
		return nil, errors.New("unknown provider")
	}

	return client, nil
}

type Client interface {
	interfaces.UserRepository
	interfaces.Repository
	interfaces.CollaboratorRepository
	interfaces.PullRequestRepository
}
