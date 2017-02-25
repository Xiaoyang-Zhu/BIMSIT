package main

import (
	"keystore"
	"fmt"
)

var (
	seed []byte
)


func main() {

	//New section of HID: if para is nil, build the identity from a scratch; if para is real identity, generate child
	keystore.NewKeystore(nil)




	// Generate a random 256 bit seed
	seed, err := keystore.GenSeed(256)
	if err != nil {
		fmt.Errorf("%s should have been nil",err.Error())
	}
	fmt.Println(seed)

	// Create a master private key
	masterprv := keystore.MasterKey(seed)
	fmt.Println(masterprv)
	fmt.Println(masterprv.Serialize())
	// Convert a private key to public key
	masterpub := masterprv.Pub()
	fmt.Println(masterpub)
	fmt.Println(masterpub.Serialize())

	// Generate new child key based on private or public key
	childprv, err := masterprv.Child(0)
	fmt.Println(childprv)
	childpub, err := masterpub.Child(0)
	fmt.Println(childpub)


	// Create bitcoin address from public key
	address := childpub.Address()
	fmt.Println(address)

	// Convenience string -> string Child and ToAddress functions
	walletstring := childpub.String()
	childstring, err := keystore.StringChild(walletstring,0)
	childaddress, err := keystore.StringAddress(childstring)
	fmt.Println(childaddress)

}