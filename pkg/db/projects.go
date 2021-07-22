package db

type Project struct {
	Model
	Name        string
	Description string
	GithubURL   string
	ProjectURL  string
	Public      bool
}
