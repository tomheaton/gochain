package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	//"crypto/sha256"
	//"encoding/json"
	//"strconv"
	//"strings"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int // Proof of Work
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (block Block) calculateHash() string {
	data, _ := json.Marshal(block.data)
	blockData := block.previousHash + string(data) + block.timestamp.String() + strconv.Itoa(block.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func generateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	return Blockchain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficulty,
	}
}

func (block *Block) mine(difficulty int) {
	for !strings.HasPrefix(block.hash, strings.Repeat("0", difficulty)) {
		block.pow++
		block.hash = block.calculateHash()
	}
}

func (blockchain *Blockchain) addBlock(from string, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := blockchain.chain[len(blockchain.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(blockchain.difficulty)
	blockchain.chain = append(blockchain.chain, newBlock)
}

func (blockchain Blockchain) isValid() bool {
	for i := range blockchain.chain[1:] {
		previousBlock := blockchain.chain[i]
		currentBlock := blockchain.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func (blockchain Blockchain) viewBlockchain() {
	fmt.Println("Blockchain Information:")
	fmt.Printf("\tLength: %x\n", len(blockchain.chain))
	// TODO: pretty print this.
	fmt.Printf("\tChain: %v\n", blockchain.chain)
}

func main() {
	fmt.Println("Generating the GoChain")
	blockchain := generateBlockchain(1)

	// Interact with blockchain.
	blockchain.addBlock("A", "B", 7)
	blockchain.addBlock("B", "C", 9)

	if blockchain.isValid() {
		fmt.Println("The GoChain is valid.")
	} else {
		fmt.Println("The GoChain is not valid.")
	}

	blockchain.viewBlockchain()
}
