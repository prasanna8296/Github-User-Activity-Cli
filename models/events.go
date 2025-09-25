package models

type GitHubEvents struct {
	Type  string `json:"type"`
	Actor Actor  `json:"actor"`
	Repo  Repo   `json:"repo"`
}
type Actor struct {
	Login string `json:"login"`
}
type Repo struct {
	Name string `json:"name"`
}
