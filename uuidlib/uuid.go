package uuidlib

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

// UUID UUID
func UUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(),"-","")
}

