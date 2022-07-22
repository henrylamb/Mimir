package repo

import (
	"encoding/json"
	"log"
)

//MapToSlice this function converts a map into a slice. This will usually convert the graphql params into a new slice.
func MapToSlice(params map[string]interface{}) []string {
	var newSlice []string
	for key, _ := range params {
		newSlice = append(newSlice, key)
	}
	return newSlice
}

//ChangeToMap Marshalling and unmarshalling from json is an expensive process. If possible come up with an alternative.
// what could be done is that the account model may not need to be used if a rework of the gorm instructions.
//this function marshals and unmarshal the data and then put it in a map to be returned.
//use keyValue as a point that you want to remove so that it does not get updated in the updates gorm function.
func ChangeToMap(object interface{}, keyValue string) map[string]interface{} {
	var newMap map[string]interface{}
	newRecord, err := json.Marshal(object)
	if err != nil {
		log.Print("failed to marshal json")
	}
	err = json.Unmarshal(newRecord, &newMap)
	if err != nil {
		log.Print(err)
	}

	delete(newMap, "CreatedAt")
	delete(newMap, "DeletedAt")
	delete(newMap, "UpdatedAt")
	delete(newMap, keyValue)

	for key, value := range newMap {
		if value == "" {
			delete(newMap, key)
		}
		if value == 0 {
			delete(newMap, key)
		}
		if value == nil {
			delete(newMap, key)
		}
		delete(newMap, "ID")
	}
	return newMap
}
