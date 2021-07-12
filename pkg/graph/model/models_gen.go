// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Blog struct {
	Link          string `json:"link"`
	DateOfRelease string `json:"dateOfRelease"`
}

type Name struct {
	First    string `json:"first"`
	Middle   string `json:"middle"`
	Last     string `json:"last"`
	Username string `json:"username"`
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
	Name             string `json:"name"`
	Description      string `json:"description"`
	CreatedAt        string `json:"createdAt"`
	SkillsCategoryID string `json:"skillsCategoryID"`
	UpdatedAt        string `json:"updatedAt"`
}

type SkillInput struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	SkillsCategoryID string `json:"skillsCategoryID"`
}

type SkillsCategory struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type SkillsCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Social struct {
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
