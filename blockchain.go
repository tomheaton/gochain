package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	proofOfWork  int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func generateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now().UTC(),
	}

	return Blockchain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficulty,
	}
}

func (block *Block) calculateHash() string {
	data, _ := json.Marshal(block.data)
	blockData := block.previousHash + string(data) + block.timestamp.String() + strconv.Itoa(block.proofOfWork)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (block *Block) mine(difficulty int) {
	for !strings.HasPrefix(block.hash, strings.Repeat("0", difficulty)) {
		block.proofOfWork++
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

func (blockchain *Blockchain) isValid() bool {
	for i := range blockchain.chain[1:] {
		previousBlock := blockchain.chain[i]
		currentBlock := blockchain.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func (blockchain *Blockchain) viewBlockchain() {
	fmt.Println("Blockchain Information:")
	fmt.Printf("\tLength: %x\n", len(blockchain.chain))
	for i, block := range blockchain.chain {
		fmt.Printf("\tBlock %d\n", i)
		fmt.Printf("\t\tData: %v\n", block.data)
		fmt.Printf("\t\tHash: %v\n", block.hash)
		fmt.Printf("\t\tPrevious Hash: %v\n", block.previousHash)
		fmt.Printf("\t\tTimestamp: %v\n", block.timestamp)
		fmt.Printf("\t\tProof of Work: %v\n", block.proofOfWork)
	}
}
