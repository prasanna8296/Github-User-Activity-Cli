package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github-cli/models"
	"net/http"
)

func FetchUserActivity(username string) ([]models.GitHubEvents, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("failed to fetch the api")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {

		return nil, fmt.Errorf("github api error. %s", resp.Status)

	}
	var events []models.GitHubEvents

	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, errors.New("failed to parse the response")

	}

	return events, nil

}
