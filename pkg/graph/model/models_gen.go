// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Blog struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	Summary     string `json:"summary"`
	LastUpdated string `json:"lastUpdated"`
}

type CreateExperienceInput struct {
	Name      string   `json:"name"`
	StartedAt string   `json:"startedAt"`
	EndedAt   *string  `json:"endedAt"`
	Roles     []string `json:"roles"`
}

type CreateProjectInput struct {
	Name        string `json:"name"`
	CategoryID  string `json:"categoryID"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
	GithubURL   string `json:"githubUrl"`
	ProjectURL  string `json:"projectUrl"`
}

type Experience struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	StartedAt string   `json:"startedAt"`
	EndedAt   *string  `json:"endedAt"`
	Roles     []string `json:"roles"`
	CreatedAt string   `json:"CreatedAt"`
	UpdatedAt string   `json:"UpdatedAt"`
}

type Name struct {
	First    string `json:"first"`
	Middle   string `json:"middle"`
	Last     string `json:"last"`
	Username string `json:"username"`
}

type OpenSource struct {
	OpenedPr                int    `json:"openedPr"`
	StarsReceived           int    `json:"starsReceived"`
	IssuesSubmitted         int    `json:"issuesSubmitted"`
	RepositoriesContributed int    `json:"repositoriesContributed"`
	Repositories            int    `json:"repositories"`
	Started                 string `json:"started"`
	CommitContributions     string `json:"commitContributions"`
}

type Project struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Logo        *string `json:"logo"`
	IsPublic    bool    `json:"isPublic"`
	GithubURL   string  `json:"githubUrl"`
	ProjectURL  string  `json:"projectUrl"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type ProjectEcoInput struct {
	Name string `json:"name"`
}

type ProjectEcosystem struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Projects  []*Project `json:"projects"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}

type ScretInput struct {
	Role      Role    `json:"role"`
	ExpiresIn *string `json:"expiresIn"`
}

type Secret struct {
	ID        string `json:"id"`
	Role      Role   `json:"role"`
	Token     string `json:"token"`
	ExpiresIn string `json:"expiresIn"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type Skill struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type SkillInput struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	SkillsCategoryID string `json:"skillsCategoryID"`
}

type SkillsCategory struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
	Skills      []*Skill `json:"skills"`
}

type SkillsCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Social struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type Role string

const (
	RoleGuest     Role = "GUEST"
	RoleMentainer Role = "MENTAINER"
	RoleAdmin     Role = "ADMIN"
)

var AllRole = []Role{
	RoleGuest,
	RoleMentainer,
	RoleAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleGuest, RoleMentainer, RoleAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
