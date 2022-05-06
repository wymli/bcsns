package logic

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wymli/bcsns/common/errx"
)

func CopyBuildTransactionOpts(client *ethclient.Client, baseAuth *bind.TransactOpts) (*bind.TransactOpts, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_BC_GAS, "failed to suggest gas price, err:%v", err)
	}

	auth := *baseAuth

	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	nonce, err := client.PendingNonceAt(context.Background(), baseAuth.From)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_BC_GAS, "failed to get next nonce, err:%v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))

	return &auth, nil
}
