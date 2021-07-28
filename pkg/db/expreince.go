package db

import (
	"time"

	"github.com/lib/pq"
)

type Experience struct {
	Model
	Name      string
	Roles     pq.StringArray `gorm:"type:varchar(64)[]"`
	StartedAt time.Time
	EndedAt   time.Time
}
