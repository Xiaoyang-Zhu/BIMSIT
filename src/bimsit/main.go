// Package hdwallet implements heirarchical deterministic Bitcoin wallets, as defined in BIP 32.
//
// BIP 32 - https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
//
// This package provides utilities for generating hierarchical deterministic Bitcoin wallets.
//
// Examples
//
//          // Generate a random 256 bit seed
//          seed, err := hdwallet.GenSeed(256)
//
//          // Create a master private key
//          masterprv := hdwallet.MasterKey(seed)
//
//          // Convert a private key to public key
//          masterpub := masterprv.Pub()
//
//          // Generate new child key based on private or public key
//          childprv, err := masterprv.Child(0)
//          childpub, err := masterpub.Child(0)
//
//          // Create bitcoin address from public key
//          address := childpub.Address()
//
//          // Convenience string -> string Child and Address functions
//          walletstring := childpub.String()
//          childstring, err := hdwallet.StringChild(walletstring,0)
//          childaddress, err := hdwallet.StringAddress(childstring)
//
// Extended Keys
//
// Hierarchical deterministic wallets are simply deserialized extended keys. Extended Keys can be imported and exported as base58-encoded strings. Here are two examples:
//          public key:   "xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8"
//          private key:  "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi"
//

package main

import (
	"hdwallet"
	"fmt"
)

var (
	seed []byte
)


func main() {

	// Generate a random 256 bit seed
	seed, err := hdwallet.GenSeed(256)
	if err != nil {
		fmt.Errorf("%s should have been nil",err.Error())
	}
	fmt.Println(seed)

	// Create a master private key
	masterprv := hdwallet.MasterKey(seed)

	// Convert a private key to public key
	masterpub := masterprv.Pub()

	// Generate new child key based on private or public key
	childprv, err := masterprv.Child(0)
	fmt.Println(childprv)
	childpub, err := masterpub.Child(0)

	// Create bitcoin address from public key
	address := childpub.Address()
	fmt.Println(address)

	// Convenience string -> string Child and ToAddress functions
	walletstring := childpub.String()
	childstring, err := hdwallet.StringChild(walletstring,0)
	childaddress, err := hdwallet.StringAddress(childstring)
	fmt.Println(childaddress)

}