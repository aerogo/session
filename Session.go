package session

import "sync"

// ession represents the session information
// in a single request & response context.
type Session struct {
	id       string
	data     sync.Map
	modified bool
}

// New creates a new session with the given ID and data.
func New(sid string, baseData map[string]interface{}) *Session {
	session := &Session{
		id: sid,
	}

	for key, value := range baseData {
		session.data.Store(key, value)
	}

	return session
}

// ID returns the session ID.
func (session *Session) ID() string {
	return session.id
}

// Get returns the value for the key in this session.
func (session *Session) Get(key string) interface{} {
	value, _ := session.data.Load(key)
	return value
}

// GetString returns the string value for the key in this session.
func (session *Session) GetString(key string) string {
	value := session.Get(key)

	if value != nil {
		str, ok := value.(string)

		if ok {
			return str
		}

		return ""
	}

	return ""
}

// Set sets the value for the key in this session.
func (session *Session) Set(key string, value interface{}) {
	session.data.Store(key, value)
	session.modified = true
}

// Delete deletes the the key/value entry in this session.
func (session *Session) Delete(key string) {
	session.data.Delete(key)
	session.modified = true
}

// Modified indicates whether the session has been modified since it's been retrieved.
func (session *Session) Modified() bool {
	return session.modified
}

// Data returns a copy of the underlying session data.
func (session *Session) Data() map[string]interface{} {
	newMap := map[string]interface{}{}

	session.data.Range(func(key, value interface{}) bool {
		newMap[key.(string)] = value
		return true
	})

	return newMap
}
