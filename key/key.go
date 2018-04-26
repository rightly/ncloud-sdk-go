package key

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type KeyConfig struct {
	Key struct {
		ApiKey    string `json:"apiKey"`
		AccessKey string `json:"accessKey"`
		SecretKey string `json:"secretKey"`
	}
}

func Load(path string) KeyConfig {
	var config KeyConfig
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(f, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}