package main

import "fmt"

func main() {
	fmt.Println("Generating the blockchain...")
	blockchain := generateBlockchain(1)

	blockchain.addBlock("A", "B", 7)
	blockchain.addBlock("B", "C", 9)
	blockchain.addBlock("C", "D", 3)
	blockchain.addBlock("D", "E", 5)
	blockchain.addBlock("E", "F", 2)

	blockchain.viewBlockchain()
	fmt.Printf("Blockchain valid: %v\n", blockchain.isValid())
}
