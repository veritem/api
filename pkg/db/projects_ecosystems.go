package db

type ProjectEcosystem struct {
	Model
	Name     string
	Projects []Project
}
