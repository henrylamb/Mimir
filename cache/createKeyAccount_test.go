package cache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateKey(t *testing.T) {
	id := 1
	username := "hamRam"
	country := "UK"
	newMap := map[string]interface{}{
		"id":       id,
		"username": username,
		"country":  country,
	}

	finalString := username + "country"

	newKey := CreateKeyAccount(newMap)

	require.Equal(t, finalString, newKey)
}
