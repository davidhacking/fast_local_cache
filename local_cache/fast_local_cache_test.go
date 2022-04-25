package local_cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFastLocalCache(t *testing.T) {
	var db FastLocalCaching = New()
	err := db.Init("./db")
	assert.Nil(t, err)
}
