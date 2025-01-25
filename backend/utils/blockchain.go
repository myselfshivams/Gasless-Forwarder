package utils

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// To Get the current nonce for the given Ethereum address
func GetNonce(client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, fmt.Errorf("failed to get nonce: %v", err)
	}
	return nonce, nil
}

// To Get Chain ID from blockchain
func GetChainID(client *ethclient.Client) (*big.Int, error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch chain ID: %v", err)
	}
	return chainID, nil
}

// Sending Transaction forwards the signed transaction to Ethereum
func SendTransaction(from, to, value string, nonce uint64, data, signed string) (string, error) {
	client, err := ethclient.Dial(os.Getenv("INFURA_RPC_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("invalid private key")
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %v", err)
	}

	toAddress := common.HexToAddress(to)
	amount := new(big.Int)
	amount.SetString(value, 10)

	gasLimit := uint64(21573)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, common.Hex2Bytes(data))
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}
func SendERC20Transaction(from, to, value string, signed string) (string, error) {
	client, err := ethclient.Dial(os.Getenv("INFURA_RPC_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("invalid private key")
	}

	tokenAddress := common.HexToAddress(os.Getenv("ERC20_CONTRACT_ADDRESS"))
	toAddress := common.HexToAddress(to)

	erc20ABI, _ := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"recipient","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))

	amount, ok := big.NewInt(0).SetString(value, 10)
	if !ok {
		return "", fmt.Errorf("invalid value for amount")
	}
	data, err := erc20ABI.Pack("transfer", toAddress, amount)
	if err != nil {
		return "", fmt.Errorf("failed to pack data: %v", err)
	}

	// Get correct nonce
	fromAddress := common.HexToAddress(from)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get correct chain ID
	chainID, err := GetChainID(client)
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), 21573, big.NewInt(20000000000), data)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Println("Transaction sent successfully:", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}

func SendERC721Transaction(from, to, tokenID string, signed string) (string, error) {
	client, err := ethclient.Dial(os.Getenv("INFURA_RPC_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return "", fmt.Errorf("invalid private key")
	}

	tokenAddress := common.HexToAddress(os.Getenv("ERC721_CONTRACT_ADDRESS"))
	toAddress := common.HexToAddress(to)

	erc721ABI, _ := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))

	tokenIDBigInt := new(big.Int)
	tokenIDBigInt.SetString(tokenID, 10)

	data, err := erc721ABI.Pack("safeTransferFrom", common.HexToAddress(from), toAddress, tokenIDBigInt)
	if err != nil {
		return "", fmt.Errorf("failed to pack data: %v", err)
	}

	// Get the nonce
	fromAddress := common.HexToAddress(from)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get correct chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Create and sign the transaction
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), 200000, big.NewInt(20000000000), data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Println("Transaction sent successfully:", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}
