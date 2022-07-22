package cache

import (
	"fmt"
)

//CreateKeyAccount would be a utility function to turn a map and all the words that would be a condintioal for use within the key.
//would require pre-sorting of the input to ensure it works as intended
func CreateKeyAccount(keyMap map[string]interface{}) string {
	var username, email, dateOfBirth, gender, country, savedLibrary, active, startDate, subType string
	for key, value := range keyMap {
		switch key {
		case "username":
			username = value.(string)
		case "email":
			email = key
		case "dateOfBirth":
			dateOfBirth = key
		case "gender":
			gender = key
		case "country":
			country = key
		case "savedLibrary":
			savedLibrary = key
		case "active":
			active = key
		case "startDate":
			startDate = key
		case "subType":
			subType = key
		}
	}
	keyString := fmt.Sprint(username + email + dateOfBirth + gender + country + savedLibrary + active + startDate + subType)
	return keyString
}
