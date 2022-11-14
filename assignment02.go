package assignment02

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var t_id = 0

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {

	return rand.Intn(100)
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	block := new(Block)
	block.Nonce = GenerateNonce(blockData)
	block.BlockData = blockData
	block.PrevPointer = chainHead
	block.CurrentHash = CalculateHash(block.BlockData, block.Nonce)

	if chainHead != nil {
		block.PrevHash = chainHead.CurrentHash

	} else {
		block.PrevHash = ""
	}

	return block
}

func ListBlocks(chainHead *Block) {

	fmt.Println("Blockchain : ")

	if chainHead == nil {
		fmt.Println("No Blockchain Created.")
		return
	}
	currentNode := chainHead
	fmt.Println(strings.Repeat("=", 25))
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.PrevPointer != nil {
		currentNode = currentNode.PrevPointer
		fmt.Println(strings.Repeat("=", 25))
		fmt.Printf("%+v\n", *currentNode)
	}
}

func DisplayTransactions(blockData []Transaction) {
	for index, val := range blockData {
		fmt.Printf("%s Transaction :  %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))

		fmt.Printf(" Transaction id: %s \n Sender :  %s \n Receiver :  %s \n Amount:  %d \n \n ", val.TransactionID, val.Sender, val.Receiver, val.Amount)
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	t_id++
	block := new(Transaction)
	block.TransactionID = strconv.Itoa(t_id)
	block.Sender = sender
	block.Receiver = receiver
	block.Amount = amount
	return *block
}

func main() {

	// Create blockchain
	var blockchain Blockchain

	// Create BlockData
	var blockData []Transaction

	// Create transactions # 1
	transaction := NewTransaction("Satoshi", "Vitalik", 10)
	// Add transaction to Block
	blockData = append(blockData, transaction)

	// Create transactions # 2
	transaction = NewTransaction("Satoshi", "Cardano", 23)
	// Add transaction to Block
	blockData = append(blockData, transaction)

	// Add block to blockchain
	blockchain.ChainHead = NewBlock(blockData, nil)

	// Create transactions # 3
	transaction = NewTransaction("Alice", "Bob", 87)
	// Add transaction to Block
	blockData = append(blockData, transaction)

	// Create transactions # 4
	transaction = NewTransaction("Bob", "Alice", 10)
	// Add transaction to Block
	blockData = append(blockData, transaction)

	DisplayTransactions(blockData)

	fmt.Println(blockData)

	// Add second block to blockchain
	blockchain.ChainHead = NewBlock(blockData, blockchain.ChainHead)
	// Display blockchain
	ListBlocks(blockchain.ChainHead)

	// Verify cheating
	for blockchain.ChainHead != nil {
		if blockchain.ChainHead.CurrentHash != CalculateHash(blockchain.ChainHead.BlockData, blockchain.ChainHead.Nonce) {
			fmt.Println("Verification failed!")
			return
		}
		blockchain.ChainHead = blockchain.ChainHead.PrevPointer
	}
	fmt.Println("Verification passed!")

}
