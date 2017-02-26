package main

import (
	"keystore"
	"fmt"
	"os"
)

var (
	seed []byte
)


func main() {

	//New section of HID: if para is nil, build the identity from a scratch; if para is real identity, generate child
	//OpenKeystore file: nil, build a new one; not nil, display all identities, choose one, and pass the one to CDK

	inputFile, inputError := os.Open("keystore.data")
	if inputError == nil {

	} else {
		fmt.Printf("Cannot load the Keystore file!\n" + "Building a new one!\n")

		//Fetch the seed length from client then generate the seed: assuming length is 256 bits
		seed, err := keystore.GenSeed(256)
		if err != nil {
			fmt.Errorf("%s should have been nil",err.Error())
		}
		fmt.Printf("The seed is:\n%d\n", seed)

		//Calculate the master key and obtain the master extended private keys from the seed
		extpriv_master := keystore.MasterKey(seed)
		fmt.Printf("The extended private key is:\n%s\n", extpriv_master)
		fmt.Println(extpriv_master.Serialize())

		//Derive the m/0' keys: the hardened model based on extended private keys number: 0x80000000
		extpriv_masterchild, err := extpriv_master.Child(0)
		if err != nil {
			fmt.Errorf("%s should have been nil",err.Error())
		}
		fmt.Println(extpriv_masterchild)



		// Convert a private key to public key
		extpub := extpriv_master.Pub()
		fmt.Println(extpub)
		fmt.Println(extpub.Serialize())


		// Create bitcoin address from public key
		address := extpub.Address()
		fmt.Println(address)

		// Convenience string -> string Child and ToAddress functions
		walletstring := extpub.String()
		childstring, err := keystore.StringChild(walletstring,0)
		childaddress, err := keystore.StringAddress(childstring)
		fmt.Println(childaddress)

	}

	keystore.NewKeystore(nil)


	defer inputFile.Close()




}