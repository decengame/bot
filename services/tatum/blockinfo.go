package tatum

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/decendgame/bot/cache"
)

func GetLastestBlock() (hash string, err error) {
	reqUrl := BASE_ENDPOINT + "/ethereum/block/current"
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return
	}
	req.Header.Add("x-testnet-type", "ethereum-rinkeby")
	req.Header.Add("x-api-key", cache.TATUM_API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var blockNumber int
	err = json.Unmarshal(body, &blockNumber)
	if err != nil {
		return
	}
	hash = strconv.Itoa(blockNumber)
	hash = hash[len(hash)-1:]
	return
}
