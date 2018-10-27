package main

import (
	contract "github.com/coderbunker/heike-network/contracts/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	// Get all the env vars
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

	// Instantiate the contract
	contract, err := contract.NewIERC20(common.HexToAddress(dai), blockchain)
	if err != nil {
		log.Fatalf("failed to instantiate a contract: %v", err)
	}

	// Get credentials for the account
	auth, err := bind.NewTransactor(strings.NewReader(key), secret)
	if err != nil {
		log.Fatalf("failed to create authorized transactor: %v", err)
	}

	// Call approve function from smart contract
	contract.Approve(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, common.HexToAddress(retainer), big.NewInt(1000000000000000))
}
