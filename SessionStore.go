package session

// Store ...
type Store interface {
	Get(string) (*Session, error)
	Set(string, *Session) error
}
