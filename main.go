package main

import (
	"fmt"
)

func main() {

	CumChain := NewBlockChain()

	CumChain.AddBlock("hey this is Second block to the cum chain ")
	CumChain.AddBlock("Give Me Money") //insert spongebob meme Here

	for _, block := range CumChain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
