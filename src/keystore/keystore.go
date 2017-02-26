package keystore

import (
	"fmt"
)

//Single identity information
type IDInfo struct {
	Identifier []byte //160 bytes
	Credentials EXTKeys //The extended public or private keys
//	Attributes []byte //Attributes
}

//Entire FAKE tree structure of one identity
//Credentials in each IDInfo are composed by extended private keys which can be used to derive the public keys
type HIDS struct {
	SIDData map[string] IDInfo
//	description string
}

//Identity information storage file encrypted by users' defined password: it's a secured vault of entire identity info
type Keystore struct {
//	secret string
	seed []byte // Seed
	masterKeys []byte //Master keys: the extended private keys <privKey, chaincode>
	masterChildKeys []byte //Master child keys: m/0' the extended private keys <privKey, chaincode>
	idData []HIDS // All identity info
}

//Online Identity information set (IDS): the online HIDS
//Credentials in each IDInfo are composed by extended public keys rather than the extended private keys
type OHIDS struct {
	OSIDData map[string] IDInfo
//	description string
}

// 1 class: seed; master(m); m/0';
// 2 class: root identity offline (m/0'/0); root identity online (m/0'/0')
func (i *IDInfo) GenerateSingleIDInfo () (){

}

func  NewKeystore (parentalID *IDInfo) *Keystore {
	switch {
	case parentalID == nil:

	case parentalID != nil:

	}
	return &Keystore{}
}

func (k *Keystore) GenerateSeed (length int) () {
	//scan user input to set the length of seed
	seed, err := GenSeed(256)
	if err != nil {
		fmt.Errorf("%s should have been nil",err.Error())
	}
	fmt.Println(seed)
}

