package config

import (
	"encoding/json"
	"fmt"       // used to print errors majorly.
	"io/ioutil" // it will be used to help us read our config.json file.
	"math/big"
	"os"

	"github.com/decendgame/bot/smartcontract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Token       string // To store value of Token from config.json .
	BotPrefix   string // To store value of BotPrefix from config.json.
	PrivateKey  string
	ChainID     *big.Int
	RPCServer   string
	NFTAddress  common.Address
	NFTContract *smartcontract.Nft
	ETHClient   *ethclient.Client
	config      *configStruct // To store value extracted from config.json.
)

type configStruct struct {
	Token      string `json:"Token"`
	BotPrefix  string `json:"BotPrefix"`
	PrivateKey string `json:"PrivateKey"`
	ChainID    int    `json:"ChainID"`
	RPCServer  string `json:"RPCServer"`
	NFTAddress string `json:"NFTAddress"`
}

func ReadConfig() error {
	var err error
	if len(os.Getenv("PrivateKey")) < 20 {
		return ReadFileConfig()
	}
	Token = os.Getenv("Token")
	BotPrefix = os.Getenv("BotPrefix")
	PrivateKey = os.Getenv("PrivateKey")
	tmp := new(big.Int)
	tmp, _ = tmp.SetString(os.Getenv("PrivateKey"), 10)
	ChainID = tmp
	RPCServer = os.Getenv("RPCServer")
	NFTAddress = common.HexToAddress(os.Getenv("NFTAddress"))
	err = LoadETHClient()
	if err != nil {
		return err
	}
	err = LoadNFT()
	if err != nil {
		return err
	}

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
	ChainID = big.NewInt(int64(config.ChainID))
	RPCServer = config.RPCServer
	NFTAddress = common.HexToAddress(config.NFTAddress)
	err = LoadETHClient()
	if err != nil {
		return err
	}
	err = LoadNFT()
	if err != nil {
		return err
	}
	// If there isn't any error we will return nil.
	return nil
}

func LoadETHClient() (err error) {
	err = nil
	ETHClient, err = ethclient.Dial(RPCServer)
	if err != nil {
		fmt.Printf("Error connecting to Blockchain: %s\n", err.Error())
		return
	}
	return
}

func LoadNFT() (err error) {
	err = nil
	NFTContract, err = smartcontract.NewNft(NFTAddress, ETHClient)
	if err != nil {
		fmt.Printf("Error connecting to NFT Contract: %s\n", err.Error())
		return
	}
	return
}
