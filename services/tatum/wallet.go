package tatum

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/decendgame/bot/cache"
	"github.com/decendgame/bot/model"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func CreateWallet() (wallet hdwallet.Wallet, err error) {
	reqUrl := BASE_ENDPOINT + "/ethereum/wallet"
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return
	}

	// query := req.URL.Query()
	// query.Add("mnemonic", "string")
	// req.URL.RawQuery = query.Encode()
	req.Header.Add("x-testnet-type", "ethereum-rinkeby")
	req.Header.Add("x-api-key", cache.TATUM_API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	// fmt.Println("Body: ", string(body))

	var tatumResp model.TatumCreateWalletResp
	err = json.Unmarshal(body, &tatumResp)
	if err != nil {
		return
	}
	// fmt.Printf("Request: %+v\n", req)
	// fmt.Println("Mnemonic: ", tatumResp.Mnemonic)
	tmp, err := hdwallet.NewFromMnemonic(tatumResp.Mnemonic)
	if err != nil {
		return
	}
	wallet = *tmp
	return
}
