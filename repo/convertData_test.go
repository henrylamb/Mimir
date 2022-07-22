package repo

import (
	"github.com/stretchr/testify/require"
	"gorm.io/datatypes"
	"testing"
)

type testObject struct {
	Username     string            `json:"username"`
	Email        string            `json:"email"`
	Password     string            `json:"password"`
	DateOfBirth  string            `json:"dateOfBirth"`
	Gender       string            `json:"gender"`
	Country      string            `json:"country"`
	SavedLibrary datatypes.JSONMap `json:"savedLibrary"`
	Active       string            `json:"active"`
	StartDate    string            `json:"startDate"`
	SubType      string            `json:"subType"`
}

func TestAccount_ChangeToMap(t *testing.T) {

	username := "Henry"
	password := "stonks"
	gender := "male"
	country := "UK"

	testAccount := testObject{
		Username: username,
		Password: password,
		Gender:   gender,
		Country:  country,
	}

	compareMap := map[string]interface{}{
		"password": password,
		"gender":   gender,
		"country":  country,
	}

	testMap := ChangeToMap(&testAccount, "username")

	require.Equal(t, compareMap, testMap, "These maps should be the same. Though may not work due to the nature of maps")
}
