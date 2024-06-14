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
	Data         map[string]interface{} `json:"data"`
	Hash         string                 `json:"hash"`
	PreviousHash string                 `json:"previous_hash"`
	Timestamp    time.Time              `json:"timestamp"`
	ProofOfWork  int                    `json:"proof_of_work"`
}

type Blockchain struct {
	GenesisBlock Block   `json:"genesis_block"`
	Chain        []Block `json:"chain"`
	Difficulty   int     `json:"difficulty"`
}

func generateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		Hash:      "0",
		Timestamp: time.Now().UTC(),
	}

	return Blockchain{
		GenesisBlock: genesisBlock,
		Chain:        []Block{genesisBlock},
		Difficulty:   difficulty,
	}
}

func (block *Block) calculateHash() string {
	data, _ := json.Marshal(block.Data)
	blockData := block.PreviousHash + string(data) + block.Timestamp.String() + strconv.Itoa(block.ProofOfWork)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (block *Block) mine(difficulty int) {
	for !strings.HasPrefix(block.Hash, strings.Repeat("0", difficulty)) {
		block.ProofOfWork++
		block.Hash = block.calculateHash()
	}
}

func (blockchain *Blockchain) addBlock(from string, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := blockchain.Chain[len(blockchain.Chain)-1]
	newBlock := Block{
		Data:         blockData,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now(),
	}
	newBlock.mine(blockchain.Difficulty)
	blockchain.Chain = append(blockchain.Chain, newBlock)
}

func (blockchain *Blockchain) isValid() bool {
	for i := range blockchain.Chain[1:] {
		previousBlock := blockchain.Chain[i]
		currentBlock := blockchain.Chain[i+1]
		if currentBlock.Hash != currentBlock.calculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func (blockchain *Blockchain) viewBlockchain() {
	fmt.Println("Blockchain Information:")
	fmt.Printf("\tLength: %x\n", len(blockchain.Chain))
	for i, block := range blockchain.Chain {
		fmt.Printf("\tBlock %d\n", i)
		fmt.Printf("\t\tData: %v\n", block.Data)
		fmt.Printf("\t\tHash: %v\n", block.Hash)
		fmt.Printf("\t\tPrevious Hash: %v\n", block.PreviousHash)
		fmt.Printf("\t\tTimestamp: %v\n", block.Timestamp)
		fmt.Printf("\t\tProof of Work: %v\n", block.ProofOfWork)
	}
}
