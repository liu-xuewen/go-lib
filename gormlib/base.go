package gormlib

import (
	"github.com/liu-xuewen/go-lib/timelib"
	"time"
)

// Base Base
type Base struct {
	CreatedAt timelib.JsonTime  `gorm:"type:timestamp;index" json:"created_at"`
	UpdatedAt timelib.JsonTime  `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}
