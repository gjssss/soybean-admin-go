package system

import (
	"time"
)

type TimeRecord struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
