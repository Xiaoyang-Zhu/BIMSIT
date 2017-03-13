package keystore

import (
	"fmt"
	"log"
)

func GenHID_test(ks *Keystore) {

	//Pass a parental identity path and output the new keystore which comprises the identity
	newks, err:= ks.AddNewIDKeystore("m/0'/0")
	if err != nil {
		log.Fatalln("Errors in AddNewIDKeystore")
		return
	}
	fmt.Printf("The new serialized keystore struct data: \n%d\n", newks.Serialize())
	newks.ListAllIDPath()

	//Test v2
	new2ks, err:= newks.AddNewIDKeystore("m/0'/0")
	if err != nil {
		log.Fatalln("Errors in AddNewIDKeystore")
		return
	}
	fmt.Println(new2ks.Serialize())
	new2ks.ListAllIDPath()

	//Test v3
	new3ks, err:= new2ks.AddNewIDKeystore("m/0'/0")
	if err != nil {
		log.Fatalln("Errors in AddNewIDKeystore")
		return
	}
	fmt.Println(new3ks.Serialize())
	new3ks.ListAllIDPath()

	//Test v4
	new4ks, err:= new3ks.AddNewIDKeystore("m/0'/0/1")
	if err != nil {
		log.Fatalln("Errors in AddNewIDKeystore")
		return
	}
	fmt.Println(new4ks.Serialize())
	new4ks.ListAllIDPath()

	return
}
