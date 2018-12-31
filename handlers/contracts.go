package handlers

import (
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"

	erc20_contract "github.com/coderbunker/heikenet-backend/contracts/erc20"
	retainer_contract "github.com/coderbunker/heikenet-backend/contracts/retainer"
	mid "github.com/coderbunker/heikenet-backend/middleware"
)

func Approve(c echo.Context) error {
	// TODO: get params from json

	// get config from context
	config, err := mid.GetConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	// connect to an ethereum node hosted by infura
	blockchain, err := ethclient.Dial(config.Node)
	if err != nil {
		log.Fatal("unable to connect to network:%v\n", err)
	}

	// get credentials for the account
	auth, err := bind.NewTransactor(strings.NewReader(config.Key), config.Secret)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}

	// instantiate the ERC20 contract
	erc20_instance, err := erc20_contract.NewERC20(common.HexToAddress(config.Dai), blockchain)
	if err != nil {
		log.Fatalf("failed to instantiate a contract: %v", err)
	}

	// call approve function from smart contract
	erc20_instance.Approve(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, common.HexToAddress(config.Retainer), big.NewInt(4444))

	return c.JSON(http.StatusOK, "approve")
}

func Fund(c echo.Context) error {
	// TODO: get params from json

	// get config from context
	config, err := mid.GetConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	// connect to an ethereum node hosted by infura
	blockchain, err := ethclient.Dial(config.Node)
	if err != nil {
		log.Fatal("unable to connect to network:%v\n", err)
	}

	// get credentials for the account
	auth, err := bind.NewTransactor(strings.NewReader(config.Key), config.Secret)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}

	// instantiate the Retainer contract
	retainer_instance, err := retainer_contract.NewHeikeRetainer(common.HexToAddress(config.Retainer), blockchain)
	if err != nil {
		log.Fatalf("failed to instantiate a contract: %v", err)
	}

	// "DAI" as [32]bytes
	symbol_hex, _ := hexutil.Decode(config.Symbol)
	var token_symbol [32]byte
	copy(token_symbol[:], symbol_hex)

	// fund retainer
	fund_res, err := retainer_instance.FundRetainer(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 300000,
	}, token_symbol, big.NewInt(444))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("fund result:", fund_res)

	return c.JSON(http.StatusOK, "fund")
}

func Withdraw(c echo.Context) error {
	// TODO: get params from json

	// get config from context
	config, err := mid.GetConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	// connect to an ethereum node hosted by infura
	blockchain, err := ethclient.Dial(config.Node)
	if err != nil {
		log.Fatal("unable to connect to network:%v\n", err)
	}

	// get credentials for the account
	auth, err := bind.NewTransactor(strings.NewReader(config.Key), config.Secret)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}

	// instantiate the Retainer contract
	retainer_instance, err := retainer_contract.NewHeikeRetainer(common.HexToAddress(config.Retainer), blockchain)
	if err != nil {
		log.Fatalf("failed to instantiate a contract: %v", err)
	}

	// "DAI" as [32]bytes
	symbol_hex, _ := hexutil.Decode(config.Symbol)
	var token_symbol [32]byte
	copy(token_symbol[:], symbol_hex)

	// fund retainer
	fund_res, err := retainer_instance.FundRetainer(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 300000,
	}, token_symbol, big.NewInt(444))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("fund result:", fund_res)

	// fl_address := common.HexToAddress("0x5b872ee8aa4ed09b6b92f631f42d3aaf7622b53e")
	fl_address := common.HexToAddress("0x122e6c2B891C7EB0F31a8191016D19A9663665B9")

	with_res, err := retainer_instance.WithdrawRetainer(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, token_symbol, fl_address, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("withdr result:", with_res)

	return c.JSON(http.StatusOK, "withdraw")
}
