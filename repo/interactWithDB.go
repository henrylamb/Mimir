package repo

//GetDataFromTable This function will work as intended but tests might need to be run a few times due to the nature of how maps are randomly allocated when being looped through.
func (r *Repository) GetDataFromTable(newObject interface{}, whereQuery, selectQuery map[string]interface{}, tableName string) error {
	stringArray := MapToSlice(selectQuery)
	err := r.Db.Table(tableName).Select(stringArray).Where(whereQuery).Scan(newObject).Error
	return err
}

//GetSingleFromTable You can select many or a single object from the table depending on the searchQuery map size. Therefore, allowing you to search for either one more many values.
func (r *Repository) GetSingleFromTable(newObject interface{}, tableName, selectQuery string, searchQuery map[string]interface{}) error {
	result := r.Db.Table(tableName).Select(selectQuery).Where(searchQuery).Scan(newObject)
	return result.Error
}

func (r *Repository) CreateData(object interface{}, tableName string) error {
	//Change the password below --> method of account
	result := r.Db.Table(tableName).Create(object)
	if result.Error != nil {
		return result.Error
	}
	//can have an interaction with the cache below
	return nil
}

//UpdateData Will update the users account using the "where" conditional for the username and a converted account struct into a map.
func (r *Repository) UpdateData(object interface{}, tableName string, whereQuery map[string]interface{}, removeKey string) error {
	objectMap := ChangeToMap(object, removeKey)
	result := r.Db.Table(tableName).Where(whereQuery).Updates(objectMap)
	return result.Error
}

//DeleteData Deletes the account using the username as the identifier
func (r *Repository) DeleteData(whereQuery map[string]interface{}, deleteObject interface{}) error {
	result := r.Db.Where(whereQuery).Delete(deleteObject)
	return result.Error
}

//note for commit
