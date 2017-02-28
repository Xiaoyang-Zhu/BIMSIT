package main

import (
	"keystore"
	"fmt"
	"os"
	"io/ioutil"
)

var (
	seed []byte
)


func main() {

	//New section of HID: if para is nil, build the identity from a scratch; if para is real identity, generate child
	//OpenKeystore file: nil, build a new one; not nil, display all identities, choose one, and pass the one to CDK

	inputFile, inputError := os.Open("keystore.data")
	if inputError == nil {
		//Read the keystore file
		data, err := ioutil.ReadFile("./keystore.data")
		if err != nil {
			fmt.Println("Errors in reading file")
			return
		}
		fmt.Println(string(data))

		//Convert the file string content into keystore struct
		//func StringKeystore(data string) (*Keystore,error)
		ks, err := keystore.StringKeystore(string(data))
		if err != nil {
			fmt.Println("Errors in StringKeystore")
			return
		}
		fmt.Println(ks.Serialize())

		//List all identities tree-like structure and choose one as the parental identity



	} else {
		fmt.Printf("Cannot load the Keystore file!\n" + "Building a new one!\n")

		//Fetch the seed length from client then generate the seed: assuming length is 256 bits
		ks := keystore.NewKeystore(256) //256 bits: 32 bytes

		//check the content of the generated keystore
		fmt.Println(*ks)
		fmt.Println(ks.Serialize())
		//serialize the keystore struct
		fmt.Printf("The New Online ID extended private key is:\n%s\n", ks)

		//Write into a keystore file
		err := ioutil.WriteFile("./keystore.data", []byte(ks.String()), 0644)
		if err != nil {
			fmt.Println("Errors in writing file")
			return
		}

	}

	defer inputFile.Close()

}