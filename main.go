package main

import (
	"fmt"

	"gochain/src/blockchain"
)

func main() {
	fmt.Println("Generating the blockchain...")

	bc := blockchain.GenerateBlockchain(1)

	bc.AddBlock("A", "B", 7)
	bc.AddBlock("B", "C", 9)
	bc.AddBlock("C", "D", 3)
	bc.AddBlock("D", "E", 5)
	bc.AddBlock("E", "F", 2)

	bc.ViewBlockchain()
}
