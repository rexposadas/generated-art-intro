package rn

import (
	"github.com/google/uuid"

	"strings"
)

// ID provides a random string
func ID() string {
	s := uuid.NewString()
	parts := strings.Split(s, "-")

	// The first part of the UUID is random enough for us.
	return parts[0]
}
