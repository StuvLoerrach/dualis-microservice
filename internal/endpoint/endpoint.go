package endpoint

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"dhbw-loerrach.de/dualis/microservice/internal"
)

var tokenSecret []byte
var dbHost string
var dbUsername string
var dbPassword string
var dbName string

var db *sql.DB

// LoadService reads the settings.json file
func LoadService() {
	file, err := ioutil.ReadFile("internal/settings.json")

	if err != nil {
		log.Fatal(err)
	}

	LoadServiceFromJSON(file)
}

// LoadServiceFromJSON unmarshals settings JSON
func LoadServiceFromJSON(file []byte) {
	data := internal.Settings{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Printf("Couldn't unmarshal settings JSON: %v", err)
	}
	tokenSecret, err = base64.StdEncoding.DecodeString(data.TokenSecret)

	if err != nil {
		log.Printf("Missing secret: %v", err)
	}

	dbHost = data.DbHost
	dbUsername = data.DbUsername
	dbPassword = data.DbPassword
	dbName = data.DbName
}

func CreateDBClient() {
	var err error

	connection := fmt.Sprintf("%s:%s@%s/%s", dbUsername, dbPassword, dbHost, dbName)

	db, err = sql.Open("mysql", connection)
	if err != nil {
		log.Printf("Failed to connect to database!")
	}
}
