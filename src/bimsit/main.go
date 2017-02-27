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
		ks := keystore.NewKeystore(256)

		//check the content of the generated keystore
		fmt.Println(*ks)

	}

	defer inputFile.Close()

}