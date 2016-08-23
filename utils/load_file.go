package utils

import (
	"io/ioutil"
	"os"

	"github.com/gogap/env_json"
)

func LoadFile(path string, v interface{}) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	envJSON := env_json.NewEnvJson(env_json.ENV_JSON_KEY, env_json.ENV_JSON_EXT)
	err = envJSON.Unmarshal(data, &v)
	return
}
