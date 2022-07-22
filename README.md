# Mimir

This is an opinonated library of helpful graphQL API functionality.

To download use the following command: 

go get "github.com/henrylamb/Mimir"

## Contents

- JWT helper section
- Redis cache helper section 
- GORM database interaction

### JWT helper section 

Within this section there are functions that will parse and generate a token. This will allow for users of to check and maintain authentication whenever a request is made to an API. 

### Redis cache helper

First this folder holds a file which aims to help you create a redis key in the order of which the keyCondintions are inputed. This will allow you to convert the map of values that were updated or the object that was updated into a string which will allow a key to be created. 

The second folder holds a function that will connect to the redis database depending to the inputed address values and a test value which can be inputed once. Or this is the object of which is an example for the types of objects that the values in redis will hold. --> clarify

The other two functions are helper functions which are related to the output of the connection function. These functions will set and get redis cache values. 

### GORM database interaction

This folder of functionality is split into three main parts. Which comprise of converting data, getting a DB connection and interacting with the database. 
For converting data there is a way to convert a map into a slice. This function will be particually helpful with converting graphQL params for a select query. 
The next function changes an obejct into a map along with a value which you want to remove from the map. This would be needed so that when updating a value in the database which is an identitify is updated, as not to waste resoruces. 

The next file in this folder holds a function that will get the established connection if the global variable dbConn is assigned. This will allow other functions to interact with the database with a shared connection. There is also an interface which contains all the helper methods for interacting with the database. 

The final file contains all the methods which will interact with the database directly using GORM.

