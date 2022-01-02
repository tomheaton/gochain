package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
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
	pow          int
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

func main() {
	fmt.Println("Generating the GoChain")
	// TODO: generate the blockchain.
	generateBlockchain(1)
}
