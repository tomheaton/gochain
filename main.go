package main

import "fmt"

func main() {
	fmt.Println("Generating the blockchain...")
	blockchain := generateBlockchain(1)

	blockchain.addBlock("A", "B", 7)
	blockchain.addBlock("B", "C", 9)

	if blockchain.isValid() {
		fmt.Println("The blockchain is valid.")
	} else {
		fmt.Println("The blockchain is not valid.")
	}

	blockchain.viewBlockchain()
}
