package db

import "time"

type Secret struct {
	Model
	Role      Role
	Token     string
	ExpiresIn time.Time
}

type Role int

const (
	GUEST Role = iota
	MENTAINER
	ADMIN
)
