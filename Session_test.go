package session_test

import (
	"testing"

	"github.com/aerogo/session"
	"github.com/akyoto/assert"
)

func Test(t *testing.T) {
	s := session.New("1", map[string]interface{}{
		"key": "value",
	})

	assert.Equal(t, s.Get("key").(string), "value")
	assert.Equal(t, s.GetString("key"), "value")
	assert.Equal(t, s.Modified(), false)

	s.Set("key", "new value")
	assert.Equal(t, s.GetString("key"), "new value")
	assert.Equal(t, s.Data()["key"], "new value")
	assert.Equal(t, s.Modified(), true)

	s.Delete("key")
	assert.Equal(t, s.Get("key"), nil)
}
