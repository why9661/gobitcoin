package cli

import (
	"bitcoin/base"
	"bitcoin/wallet"
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := base.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := base.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
