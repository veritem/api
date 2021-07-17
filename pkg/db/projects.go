package db

type Project struct {
	Model
	Name        string
	Description string
	GithubUrl   string
	ProjectUrl  string
}
