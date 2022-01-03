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

func (blockchain Blockchain) addBlock(from string, to string, amount float64) {
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

func main() {
	fmt.Println("Generating the GoChain")
	// TODO: generate the blockchain.
	generateBlockchain(1)
}
