package gormlib

import (
	"github.com/liu-xuewen/go-lib/timelib"
)

// Base Base
type Base struct {
	CreatedAt timelib.JsonTime  `gorm:"type:timestamp NULL;default:NULL;index" json:"created_at"`
	UpdatedAt timelib.JsonTime  `gorm:"type:timestamp NULL;default:NULL" json:"updated_at"`
	DeletedAt *timelib.JsonTime `gorm:"type:timestamp NULL;default:NULL;index"`
}
