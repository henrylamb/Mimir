package repo

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

//The below two tests appear to work but due to issues with how gorm structures an update it is leading to errors.
//Thus meaning that the first two tests results are meaningless.
/*
func TestRepository_CreateData(t *testing.T) {
	assert.New(t)
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	postgresDBTest := postgres.New(postgres.Config{Conn: db})
	gdb, err := gorm.Open(postgresDBTest, &gorm.Config{}) // open gorm db
	if err != nil {
		t.Fatal("Failed to create gorm DB")
	}
	mock.MatchExpectationsInOrder(false)

	username := "hamRam"
	email := "hamRam@gmail.com"
	gender := "male"
	password := "password"
	active := "false"
	dateOfBirth := "12/12/12"
	country := "UK"
	savedLibrary := map[string]interface{}{"name of book": "book one"}
	startDate := "01/01/13"
	subType := "monthly"

	mock.ExpectBegin()
	mock.ExpectQuery(
		`INSERT INTO "account" (.+)`).
		WithArgs(sqlmock.AnyArg(), username, email, password, dateOfBirth, gender, country, savedLibrary, active, startDate, subType)
	mock.ExpectCommit()

	repo := Repository{
		Db: gdb,
	}

	newAccount := testObject{
		Username:     username,
		Password:     password,
		Gender:       gender,
		Active:       active,
		Email:        email,
		Country:      country,
		SavedLibrary: savedLibrary,
		StartDate:    startDate,
		SubType:      subType,
		DateOfBirth:  dateOfBirth,
	}
	repo.CreateData(&newAccount, "account")
	require.NoError(t, err, "No error should occur whilst creating this value")
}

func TestRepository_UpdateData(t *testing.T) {
	assert.New(t)
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	postgresDBTest := postgres.New(postgres.Config{Conn: db})
	gdb, err2 := gorm.Open(postgresDBTest, &gorm.Config{}) // open gorm db
	if err2 != nil {
		t.Fatal("Failed to create gorm DB")
	}
	mock.MatchExpectationsInOrder(false)

	username := "hamRam"
	email := "hamRam@gmail.com"
	gender := "male"
	password := "password"
	active := "false"

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta(`UPDATE "account" SET "email"=$1,"password"=$2,"gender"=$3,"active"=$4 WHERE "username" = $5`)).
		WithArgs(email, password, gender, active, username).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	repo := Repository{
		Db: gdb,
	}

	newAccount := testObject{
		Username: username,
		Password: password,
		Gender:   gender,
		Active:   active,
		Email:    email,
	}
	whereQuery := map[string]interface{}{
		"username": username,
	}

	err = repo.UpdateData(&newAccount, "account", whereQuery, "username")
	require.NoError(t, err, "If this fails the update function failed to work as intended.")
}


*/
func TestRepository_GetDataFromTable(t *testing.T) {
	// database and initial setup
	assert.New(t)
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	postgresDBTest := postgres.New(postgres.Config{Conn: db})
	gdb, err := gorm.Open(postgresDBTest, &gorm.Config{}) // open gorm db
	if err != nil {
		t.Fatal("Failed to create gorm DB")
	}
	mock.MatchExpectationsInOrder(false)

	username := "hamRam"
	email := "hamRam@gmail.com"
	gender := "male"
	mapSelectQuery := map[string]interface{}{"gender": gender}
	mock.ExpectQuery(
		`SELECT gender FROM "account" WHERE "username" = $1`).WithArgs(username).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username", "email", "password", "gender"}).AddRow("one", username, email, "password", gender))

	repo := Repository{
		Db: gdb,
	}

	newAccount := testObject{}
	whereQuery := map[string]interface{}{
		"username": username,
	}

	repo.GetDataFromTable(&newAccount, whereQuery, mapSelectQuery, "account")
	require.Equal(t, gender, newAccount.Gender, "Should be equal. The presence of maps may fuck things up.")
}

func TestRepository_DeleteData(t *testing.T) {
	assert.New(t)
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	postgresDBTest := postgres.New(postgres.Config{Conn: db})
	gdb, err2 := gorm.Open(postgresDBTest, &gorm.Config{}) // open gorm db
	if err2 != nil {
		t.Fatal("Failed to create gorm DB")
	}
	username := "hamRam"

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "test_objects" WHERE "username" = $1`).
		WithArgs(username).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	repo := Repository{
		Db: gdb,
	}
	whereQuery := map[string]interface{}{
		"username": username,
	}
	newAccount := testObject{}
	err = repo.DeleteData(whereQuery, newAccount)
	require.NoError(t, err)

}
