package session

// Store ...
type Store interface {
	Get(string) *Session
	Set(string, *Session)
}
