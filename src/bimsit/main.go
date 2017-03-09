package main

//go:generate abigen -sol ../contracts/txID.sol -pkg contracts -out ../contracts/txID.go


import (
	"keystore"
	"transaction"
	"fmt"
	"os"
	"io/ioutil"
	"log"
)


func main() {

	//New section of HID: if para is nil, build the identity from a scratch; if para is real identity, generate child
	//OpenKeystore file: nil, build a new one; not nil, display all identities, choose one, and pass the one to CDK

	fmt.Println("Reading Keystore File!")

	inputFile, inputError := os.Open("keystore.data")
	if inputError == nil {
		//Read the keystore file
		data, err := ioutil.ReadFile("./keystore.data")
		if err != nil {
			fmt.Println("Errors in reading file")
			return
		}
		fmt.Printf("The keystore content string is: \n%s\n", string(data))

		//Convert the file string content into keystore struct
		//func StringKeystore(data string) (*Keystore,error)
		ks, err := keystore.StringKeystore(string(data))
		if err != nil {
			fmt.Println("Errors in StringKeystore")
			return
		}
		fmt.Printf("The serialized keystore struct data: \n%d\n", ks.Serialize())

		//List all identities tree-like structure and choose one as the parental identity
		ks.ListAllIDPath()

		//Test arbitrary identity generation
		keystore.GenHID_test(ks)


	} else {
		fmt.Printf("Cannot load the Keystore file!\n" + "Building a new one!\n")

		// Fetch the seed length from client then generate the seed: assuming length is 256 bits
		ks := keystore.NewKeystore(256) //256 bits: 32 bytes

		// Check the content of the generated keystore
		fmt.Printf("The new raw keystore\n%d\n", *ks)
		fmt.Printf("The serialized keystore\n%d\n", ks.Serialize())
		// Serialize the keystore struct
		fmt.Printf("The string format serialized Keystore:\n%s\n", ks)

		// Write into a keystore file
		err := ioutil.WriteFile("./keystore.data", []byte(ks.String()), 0644)
		if err != nil {
			fmt.Println("Errors in writing file")
			return
		}

		// Prepare data for identity registration process for smart contract
		rootID, rootPKf, rootPKo, err := ks.GetRootIDInfo()
		if err != nil {
			log.Fatalf("Failed to get the root identity information: %v", err)
			return
		}

		//Fake signature and pointer for prototyping
		rootPointer := rootID
			sig := rootID

		// Pass the value and deploy the contract
		tx_str, err := transaction.TxRegConn(rootID, rootPKf, rootPKo, sig, rootPointer)
		if err != nil {
			fmt.Println("Errors in identity registration process")
			return
		}

		fmt.Println(tx_str)

	}



	defer inputFile.Close()

}