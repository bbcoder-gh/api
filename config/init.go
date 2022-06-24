package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	Description string `json:"description,omitempty"`
	ServerPort  string `json:"server_port,omitempty"`
	Env         string `json:"env,omitempty"`
	DBHost      string `json:"db_host,omitempty"`
	DBUser      string `json:"db_user,omitempty"`
	DBPass      string `json:"db_pass,omitempty"`
	DBPort      string `json:"db_port,omitempty"`
	DBName      string `json:"db_name,omitempty"`
}

var DBString string
var Configuration *Config

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

	Configuration = new(Config)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &Configuration)

	DBString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", Configuration.DBUser, Configuration.DBPass, Configuration.DBHost, Configuration.DBPort, "authman")
	fmt.Println(DBString)
}

func init() {
	Configure()
}
