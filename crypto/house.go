package crypto

import (
	"math/big"

	"github.com/decendgame/bot/config"
	"github.com/decendgame/bot/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/jeffprestes/goethereumhelper"
)

func HouseOwnershipTransfer(buyer, seller *model.Player, houseID int) (txID string, err error) {
	var opts *bind.TransactOpts
	opts, err = goethereumhelper.GetKeyedTransactorWithOptions(config.ETHClient, 1, 1, seller.Key)
	if err != nil {
		return
	}
	tx, err := config.NFTContract.SafeTransferFrom(opts, seller.EthereumAddress(), buyer.EthereumAddress(), big.NewInt(int64(houseID)), big.NewInt(1), nil)
	if err != nil {
		return
	}
	txID = tx.Hash().String()
	return
}
