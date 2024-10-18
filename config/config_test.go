package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Setenv("PORT", "3333")

	got, err := NewConfig()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}

	assert.Equal(t, 3333, got.Port)
	assert.Equal(t, "dev", got.Env)
}
