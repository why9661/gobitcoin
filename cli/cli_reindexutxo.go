package cli

import (
	"bitcoin/base"
	"fmt"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := base.NewBlockchain(nodeID)
	UTXOSet := base.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
