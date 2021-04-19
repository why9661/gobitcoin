package cli

import (
	"bitcoin/base"
	"bitcoin/server"
	"bitcoin/wallet"
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := base.NewBlockchain(nodeID)
	UTXOSet := base.UTXOSet{bc}
	defer bc.DB.Close()

	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := base.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := base.NewCoinbaseTX(from, "")
		txs := []*base.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		server.SendTx(server.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
