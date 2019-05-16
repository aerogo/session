package session

// Store describes the interface for a session store.
type Store interface {
	Get(string) (*Session, error)
	Set(string, *Session) error
	Delete(string)
}
