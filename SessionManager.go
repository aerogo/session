package session

import "time"

// Manager ...
type Manager struct {
	Store Store

	// Session duration in seconds
	Duration int
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
	manager.Store.Set(session.id, session)
	return session
}
