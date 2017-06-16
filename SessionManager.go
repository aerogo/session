package session

// Manager ...
type Manager struct {
	Store Store
}

// New creates a new session.
func (manager *Manager) New() *Session {
	sessionID := GenerateUUID()

	// Session data is not allowed to be empty.
	// Therefore we are adding the session ID as dummy data.
	sessionData := map[string]interface{}{
		"sid": sessionID,
	}

	session := New(sessionID, sessionData)
	manager.Store.Set(session.id, session)
	return session
}
