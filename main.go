package main

import (
	"fmt"
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

func main() {
	fmt.Println("Generating the GoChain")
}
