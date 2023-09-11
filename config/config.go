package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	decoder := json.NewDecoder(file)
	config = &configStruct{}
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}