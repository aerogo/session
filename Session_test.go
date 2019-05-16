package session_test

import (
	"testing"

	"github.com/aerogo/session"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := session.New("1", map[string]interface{}{
		"key": "value",
	})

	assert.Equal(t, s.Get("key").(string), "value")
	assert.Equal(t, s.GetString("key"), "value")
	assert.False(t, s.Modified())

	s.Set("key", "new value")
	assert.Equal(t, s.GetString("key"), "new value")
	assert.Equal(t, s.Data()["key"], "new value")
	assert.True(t, s.Modified())

	s.Delete("key")
	assert.Equal(t, s.Get("key"), nil)
}
