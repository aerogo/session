package session

import "sync"

// Session ...
type Session struct {
	id       string
	data     map[string]interface{}
	lock     sync.RWMutex
	modified bool
}

// New creates a new session with the given ID and data.
func New(sid string, baseData map[string]interface{}) *Session {
	return &Session{
		id:   sid,
		data: baseData,
	}
}

// ID returns the session ID.
func (session *Session) ID() string {
	return session.id
}

// Get returns the value for the key in this session.
func (session *Session) Get(key string) interface{} {
	session.lock.RLock()
	value := session.data[key]
	session.lock.RUnlock()
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
	session.lock.Lock()

	if value == nil {
		delete(session.data, key)
	} else {
		session.data[key] = value
	}

	session.lock.Unlock()
	session.modified = true
}

// Modified indicates whether the session has been modified since it's been retrieved.
func (session *Session) Modified() bool {
	return session.modified
}

// Data returns a copy of the underlying session data.
func (session *Session) Data() map[string]interface{} {
	if session.data == nil {
		return nil
	}

	newMap := map[string]interface{}{}

	for key, value := range session.data {
		newMap[key] = value
	}

	return newMap
}
