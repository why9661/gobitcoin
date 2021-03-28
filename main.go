package main

import (
	"fmt"
	"goblockchain/block"
)

func main() {
	bc := block.NewBlockChain()

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")

	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PreviousHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
