package crypto

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/decendgame/bot/config"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func InitalTopup(recipient *ecdsa.PrivateKey, amount int) (txID string, err error) {
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.ETHClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(int64(amount)) // in wei (1 eth)
	gasLimit := uint64(21000)          // in units
	gasPrice, err := config.ETHClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	publicKey = recipient.Public()
	publicKeyECDSA, ok = publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", err
	}

	toAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// chainID, err := client.NetworkID(context.Background())
	chainID := big.NewInt(31949730)
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = config.ETHClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	fmt.Printf("InitalTopup Transaction %s has been sent\n", signedTx.Hash().String())
	return signedTx.Hash().String(), nil
}
