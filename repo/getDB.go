package repo

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

type Repo interface {
	GetDataFromTable(newObject interface{}, whereQuery, selectQuery map[string]interface{}, tableName string) error
	GetSingleFromTable(newObject interface{}, tableName, selectQuery string, searchQuery map[string]interface{}) error
	CreateData(object interface{}, tableName string) error
	UpdateData(object interface{}, tableName string, whereQuery map[string]interface{}, whereSingle string) error
	DeleteData(whereQuery map[string]interface{}, deleteObject interface{}) error
}

//dbConn this value must be assigned by the create database connection so that the below function may work.
var dbConn *gorm.DB

// GetDatabaseConnection this may not work as intended with the change in the struct type.
func GetDatabaseConnection() (Repo, error) {
	sqlDB, err := dbConn.DB()
	repo := &Repository{
		Db: dbConn,
	}
	if err != nil {
		return repo, err
	}
	if err = sqlDB.Ping(); err != nil {
		return repo, err
	}

	return repo, nil
}
