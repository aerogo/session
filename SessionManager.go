package session

// Manager ...
type Manager struct {
	Store Store

	// Session duration in seconds
	Duration int
}

// New creates a new session.
func (manager *Manager) New() *Session {
	sessionID := GenerateID()
	sessionData := make(map[string]interface{})
	session := New(sessionID, sessionData)
	manager.Store.Set(session.id, session)
	return session
}
