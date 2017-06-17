package session

import (
	"github.com/google/uuid"
)

// sessionIDLength is a constant defining the length of session IDs.
const sessionIDLength = 36

// GenerateID generates a unique ID.
func GenerateID() string {
	return uuid.New().String()
}

// IsValidID returns whether the string is a valid session ID or not.
func IsValidID(sid string) bool {
	return len(sid) == sessionIDLength
}
