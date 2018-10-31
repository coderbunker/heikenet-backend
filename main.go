package main

import (
	erc20_contract "github.com/coderbunker/heike-network/contracts/erc20"
	retainer_contract "github.com/coderbunker/heike-network/contracts/retainer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	// get all the env vars
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		log.Fatal("$DATABASE_URL must be set")
	}

	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("$KEY must be set")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("$SECRET must be set")
	}

	node := os.Getenv("NODE")
	if node == "" {
		log.Fatal("$NODE must be set")
	}

	dai := os.Getenv("DAI")
	if dai == "" {
		log.Fatal("$DAI must be set")
	}

	retainer := os.Getenv("RETAINER")
	if retainer == "" {
		log.Fatal("$RETAINER must be set")
	}

	// connect to an ethereum node hosted by infura
	blockchain, err := ethclient.Dial(node)
	if err != nil {
		log.Fatalf("unable to connect to network:%v\n", err)
	}

	// get credentials for the account
	auth, err := bind.NewTransactor(strings.NewReader(key), secret)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}

	// instantiate the ERC20 contract
	erc20_instance, err := erc20_contract.NewERC20(common.HexToAddress(dai), blockchain)
	if err != nil {
		log.Fatalf("failed to instantiate a contract: %v", err)
	}

	// call approve function from smart contract
	// erc20_instance.Approve(&bind.TransactOpts{
	// 	From:   auth.From,
	// 	Signer: auth.Signer,
	// 	Value:  nil,
	// }, common.HexToAddress(retainer), big.NewInt(4444))

	// instantiate the Retainer contract
	retainer_instance, err := retainer_contract.NewHeikeRetainer(common.HexToAddress(retainer), blockchain)
	if err != nil {
		log.Fatalf("failed to instantiate a contract: %v", err)
	}

	// "DAI" as [32]bytes
	// symbol := "DAI"
	// var token_symbol [32]byte
	// copy(token_symbol[:], symbol)

	symbol := "0x686573"
	symbol_hex, _ := hexutil.Decode(symbol)
	var token_symbol [32]byte
	copy(token_symbol[:], symbol_hex)

	// add token to retainer
	// a, err := retainer_instance.AddNewToken(&bind.TransactOpts{
	// 	From:     auth.From,
	// 	Signer:   auth.Signer,
	// 	GasLimit: 50000,
	// }, token_symbol, common.HexToAddress(dai))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("add result:", a)
	// log.Println("auth", auth)

	// tokens_res, err := retainer_instance.Tokens(&bind.CallOpts{}, token_symbol)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("tokens result:", tokens_res)

	// fund retainer
	// fund_res, err := retainer_instance.FundRetainer(&bind.TransactOpts{
	// 	From:     auth.From,
	// 	Signer:   auth.Signer,
	// 	GasLimit: 300000,
	// }, token_symbol, big.NewInt(444))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("fund result:", fund_res)

	balance, err := erc20_instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(retainer))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("retainer balance (wei): %s\n", balance)

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
}
