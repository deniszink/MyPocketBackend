package settings

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var environments = map[string]string{
	"production": "api.jwt.auth/settings/prod.json",
	"preproduction": "/Users/denis/Programming/Golang/src/backend//api.jwt.auth/settings/pre.json",
	//"tests": "../../settings/tests.json",
}

type Settings struct {
	PrivateKeyPath string
	PublicKeyPath string
	JWTExpirationDelta int
}

var settings Settings = Settings{}
var env = "preproduction"

func Init(){
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string){

	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}