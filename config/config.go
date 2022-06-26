package config

import (
	"encoding/json"
	"fmt"       // used to print errors majorly.
	"io/ioutil" // it will be used to help us read our config.json file.
	"os"
)

var (
	Token      string // To store value of Token from config.json .
	BotPrefix  string // To store value of BotPrefix from config.json.
	PrivateKey string
	config     *configStruct // To store value extracted from config.json.
)

type configStruct struct {
	Token      string `json:"Token"`
	BotPrefix  string `json:"BotPrefix"`
	PrivateKey string `json:"PrivateKey"`
}

func ReadConfig() error {
	if len(os.Getenv("PrivateKey")) < 20 {
		return ReadFileConfig()
	}
	Token = os.Getenv("Token")
	BotPrefix = os.Getenv("BotPrefix")
	PrivateKey = os.Getenv("PrivateKey")
	// If there isn't any error we will return nil.
	return nil
}

func ReadFileConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json") // ioutil package's ReadFile method which we read config.json and return it's value we will then store it in file variable and if an error ocurrs it will be stored in err .
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// After storing value in config variable we will access it and storing it in our declared variables .
	Token = config.Token
	BotPrefix = config.BotPrefix
	PrivateKey = config.PrivateKey
	// If there isn't any error we will return nil.
	return nil
}
