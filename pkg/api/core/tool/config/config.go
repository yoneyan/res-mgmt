package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Controller Controller `json:"controller"`
	DB         DB         `json:"db"`
	Slack      []Slack    `json:"slack"`
	Genre      []Genre    `json:"genre"`
	User       []User     `json:"user"`
}

type Controller struct {
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type DB struct {
	Path string `json:"path"`
}

type Slack struct {
	WebHookUrl string `json:"url"`
	Channel    string `json:"channel"`
	Name       string `json:"name"`
}

type Genre struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

var Conf Config

func GetConfig(inputConfPath string) error {
	configPath := "./data.json"
	if inputConfPath != "" {
		configPath = inputConfPath
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	var data Config
	json.Unmarshal(file, &data)
	Conf = data
	return nil
}

//func GetConfig(inputConfPath string) (Config, error) {
//	configPath := "./data.json"
//	if inputConfPath != "" {
//		configPath = inputConfPath
//	}
//	file, err := ioutil.ReadFile(configPath)
//	if err != nil {
//		return Config{}, err
//	}
//	var data Config
//	json.Unmarshal(file, &data)
//	return data, nil
//}
