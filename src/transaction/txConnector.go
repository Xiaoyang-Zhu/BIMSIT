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

const key = `{"address":"5b824da66f06be2af5e99b070308cef7294adf9b","crypto":{"cipher":"aes-128-ctr","ciphertext":"98f3bd201e601d0eb71defd32639d715de8384b3647f1324f368e61e583d7abd","cipherparams":{"iv":"53388169c2a04012926c6c6cc9e960b7"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"4cb374e6c77edfed3b8b2cbeeaba03c05a4603986f300f3f14378e8555f3cab4"},"mac":"f1f22578415c55a7335670cb53341f8e40bec5a217f392f365ef41ae15e54808"},"id":"114aa052-3b05-4ba8-a41a-d60a59ade246","version":3}`

func TxRegConn(rootID, rootPKf, rootPKo, sig, rootPointer []byte) (string, error)  {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract
	_, err = contracts.NewTxReg(common.BytesToAddress(rootID), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a ID transaction contract: %v", err)
	}

	// Create an authorized transactor
	auth, err := bind.NewTransactor(strings.NewReader(key), "1234567890")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Deploy the new identity contract
	// new(big.Int), "Contracts in Go!!!", 0, "Go!")
	address, tx, tx_identifier, err := contracts.DeployTxReg(auth, conn, common.BytesToAddress(rootID), rootPKf, rootPKo, sig, rootPointer)
	if err != nil {
		log.Fatalf("Failed to deploy new identity contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx_identifier)

	return tx.String(), nil
}