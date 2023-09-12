package providers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ammyy9908/authlibrary/interfaces"
)

type GitLabClient struct {
	Token string
}

func (gl *GitLabClient) GetProfile(username string) (interfaces.UserProfile, error) {
	apiUrl := fmt.Sprintf("https://gitlab.com/api/v4/users?username=%s", username)

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP request
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return interfaces.UserProfile{}, err
	}

	// Set headers or authentication if needed
	// For example, if GitLab API requires an access token:
	// req.Header.Add("Private-Token", "YOUR_ACCESS_TOKEN")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return interfaces.UserProfile{}, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return interfaces.UserProfile{}, fmt.Errorf("GitLab API request failed with status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into a UserProfile struct
	var userProfile interfaces.UserProfile
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&userProfile); err != nil {
		return interfaces.UserProfile{}, err
	}

	return userProfile, nil
}

func (gl *GitLabClient) CreateRepo(name string, private bool) (interfaces.RepoDetail, error) {
	// Implement using GitLab API
	return interfaces.RepoDetail{}, nil
}

func (gl *GitLabClient) ListRepos() ([]interfaces.RepoDetail, error) {
	// Implement using GitLab API
	return []interfaces.RepoDetail{}, nil
}

// Implement other methods
