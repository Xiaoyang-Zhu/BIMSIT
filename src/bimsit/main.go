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


		//Test function
		//Pass a parental identity path and output the new keystore which comprises the identity
		newks, err:= ks.AddNewIDKeystore("m/0'/0")
		if err != nil {
			fmt.Println("Errors in AddNewIDKeystore")
			return
		}
		fmt.Printf("The new serialized keystore struct data: \n%d\n", newks.Serialize())
		newks.ListAllIDPath()

		////Test v2
		//new2ks, err:= newks.AddNewIDKeystore("m/0'/0")
		//if err != nil {
		//	fmt.Println("Errors in AddNewIDKeystore")
		//	return
		//}
		//fmt.Println(new2ks.Serialize())
		//new2ks.ListAllIDPath()
		//
		////Test v2
		//new3ks, err:= new2ks.AddNewIDKeystore("m/0'/0")
		//if err != nil {
		//	fmt.Println("Errors in AddNewIDKeystore")
		//	return
		//}
		//fmt.Println(new3ks.Serialize())
		//new3ks.ListAllIDPath()


	} else {
		fmt.Printf("Cannot load the Keystore file!\n" + "Building a new one!\n")

		//Fetch the seed length from client then generate the seed: assuming length is 256 bits
		ks := keystore.NewKeystore(256) //256 bits: 32 bytes

		//check the content of the generated keystore
		fmt.Printf("The new raw keystore\n%d\n", *ks)
		fmt.Printf("The serialized keystore\n%d\n", ks.Serialize())
		//serialize the keystore struct
		fmt.Printf("The string format serialized Keystore:\n%s\n", ks)

		//Write into a keystore file
		err := ioutil.WriteFile("./keystore.data", []byte(ks.String()), 0644)
		if err != nil {
			fmt.Println("Errors in writing file")
			return
		}

	}

	defer inputFile.Close()

}