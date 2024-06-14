package main

import "fmt"

func main() {
	fmt.Println("Generating the blockchain...")
	blockchain := generateBlockchain(1)

	blockchain.addBlock("A", "B", 7)
	blockchain.addBlock("B", "C", 9)

	blockchain.viewBlockchain()
	fmt.Printf("Blockchain valid: %v\n", blockchain.isValid())
}
