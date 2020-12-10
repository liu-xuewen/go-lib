package gormlib

import (
	"github.com/liu-xuewen/go-lib/timelib"
	"time"
)

// Base Base
type Base struct {
	CreatedAt timelib.JsonTime  `sql:"index"`
	UpdatedAt timelib.JsonTime
	DeletedAt *time.Time `sql:"index"`
}
