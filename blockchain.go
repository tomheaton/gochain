package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// NumberOfWorkers Number of concurrent workers for mining
	NumberOfWorkers = 4
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
	data, err := json.Marshal(block.Data)
	if err != nil {
		log.Fatalf("Error marshaling block data: %v", err)
	}
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

func (block *Block) mineConcurrent(difficulty int) {
	target := strings.Repeat("0", difficulty)
	resultChannel := make(chan Block)
	var wg sync.WaitGroup

	worker := func(startNonce int) {
		defer wg.Done()
		localBlock := *block
		localBlock.ProofOfWork = startNonce
		for !strings.HasPrefix(localBlock.Hash, target) {
			localBlock.ProofOfWork++
			localBlock.Hash = localBlock.calculateHash()
		}
		resultChannel <- localBlock
	}

	wg.Add(NumberOfWorkers)
	for i := 0; i < NumberOfWorkers; i++ {
		// Each worker starts with a different nonce
		go worker(i * 1_000_000)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	*block = <-resultChannel
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
	//newBlock.mineConcurrent(blockchain.Difficulty)
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
		fmt.Printf("\t\tTimestamp: %v\n", block.Timestamp.Format(time.RFC3339))
		fmt.Printf("\t\tProof of Work: %v\n", block.ProofOfWork)
	}
}
