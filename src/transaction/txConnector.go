package transaction

import(
	"log"
	"contracts"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"fmt"
)

const key = `{"address":"f2759b4a699dae4fdc3383a0d7a92cfc246315cd","crypto":{"cipher":"aes-128-ctr","ciphertext":"a96fe235356c7ebe6520d2fa1dcc0fd67199cb490fb18c39ffabbb6880a6b3d6","cipherparams":{"iv":"47182104a4811f8da09c0bafc3743e2a"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"81c82f97edb0ee1036e63d1de57b7851271273971803e60a5cbb011e85baa251"},"mac":"09f107c9af8efcb932354d939beb7b2c0cebcfd70362d68905de554304a7cfff"},"id":"eb7ed04f-e996-4bda-893b-28dc6ac24626","version":3}`

func TxRegConn(rootID, rootPKf, rootPKo, sig, rootPointer []byte) (string, error)  {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract
	id, err := contracts.NewTxID(common.BytesToAddress(rootID), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a ID transaction contract: %v", err)
	}

	// Create an authorized transactor
	auth, err := bind.NewTransactor(strings.NewReader(key), "1234567890")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Identity contracts constructor
	_, err = id.TxReg(auth, common.BytesToAddress(rootID), rootPKf, rootPKo, sig, rootPointer)
	if err != nil {
		log.Fatalf("Failed to register the identity: %v", err)
	}

	// Deploy the new identity contract
	// new(big.Int), "Contracts in Go!!!", 0, "Go!")
	address, tx, tx_identifier, err := contracts.DeployTxID(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new identity contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx_identifier)

	return tx.String(), nil
}