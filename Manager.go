package session

import (
	"fmt"
	"net/http"
	"time"
)

// Manager generates new sessions with the specified
// duration and within the given storage backend.
type Manager struct {
	Store Store

	// Session duration in seconds
	Duration int

	// SameSite attribute for cookies
	SameSite http.SameSite
}

// New creates a new session.
func (manager *Manager) New() *Session {
	sessionID := GenerateID()

	// Session data is not allowed to be an empty map.
	// Therefore we add the session ID itself as predefined data in "sid".
	// Additionally we'll save the creation date (UTC) in "created".
	sessionData := map[string]interface{}{
		"sid":     sessionID,
		"created": time.Now().UTC().Format(time.RFC3339),
	}

	session := New(sessionID, sessionData)
	err := manager.Store.Set(session.id, session)

	if err != nil {
		fmt.Println("Error saving session in the session store:", err)
	}

	return session
}

// Cookie creates a cookie for the given session.
func (manager *Manager) Cookie(session *Session) *http.Cookie {
	return &http.Cookie{
		Name:     "sid",
		Value:    session.ID(),
		HttpOnly: true,
		Secure:   true,
		MaxAge:   manager.Duration,
		Path:     "/",
		SameSite: manager.SameSite,
	}
}
