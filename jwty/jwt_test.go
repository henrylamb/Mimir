package jwty

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	username := "John Smith"
	secretKey := []byte("secret")
	token, err := GenerateToken(username, "username", secretKey)
	if err != nil {
		t.Logf("generating a JWT failed: %s", err)
	}
	t.Logf("The token that is produced is: %s", token)
	require.NotEmpty(t, token, "This token value should not be empty. See above for the token that was produced")
}

func TestParseToken(t *testing.T) {
	username := "Jane Doe"
	secretKey := []byte("secret")
	token, err := GenerateToken(username, "username", secretKey)
	if err != nil {
		t.Logf("generating a JWT failed: %s", err)
		t.Error(err)
	}
	user, err := ParseToken(token, "username", []byte("secret"))
	if err != nil {
		t.Logf("failed to parse token: %s", err)
		t.Error(err)
	}
	t.Logf("The user value from parsing the token is: %s", user)
	require.Equal(t, user, username, "The user values should be the same.")
}
