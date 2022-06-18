package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Env         string `json:"env"`
	DBHost      string `json:"db_host"`
	DBUser      string `json:"db_user"`
	DBPass      string `json:"db_pass"`
	DBPort      string `json:"db_port"`
	DBName      string `json:"db_name"`
}

var DBString string

func Configure() {

	var jsonFile *os.File
	var err error

	authmanConfigFilePath := os.Getenv("AUTHMAN_CONFIG_FILE")
	if authmanConfigFilePath == "" {
		authmanConfigFilePath = "./config/config.json"
	}

	jsonFile, err = os.Open(authmanConfigFilePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var _config Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &_config)

	DBString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", _config.DBUser, _config.DBPass, _config.DBHost, _config.DBPort, "authman")
	fmt.Println(DBString)
}

func init() {
	Configure()
}
