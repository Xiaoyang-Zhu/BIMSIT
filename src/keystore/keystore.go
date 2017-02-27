package keystore

import (
	"fmt"
	"strconv"
	"btcsuite/btcutil/base58"
)

//Single identity information
type IDInfo struct {
	Identifier []byte //160 bytes
	Credentials []EXTKeys //The extended public or private keys: the first ExtKey is offline keys; the second is online
//	Attributes []byte //Attributes
}

//Entire FAKE tree structure of one identity
//Credentials in each IDInfo are composed by extended private keys which can be used to derive the public keys
type HIDS struct {
	SIDData map[string] *IDInfo //string stands for the pathway of identity
	ChildrenNum	int
//	description string
}

//Identity information storage file encrypted by users' defined password: it's a secured vault of entire identity info
type Keystore struct {
//	secret string
	seed []byte // Seed
	masterKeys *EXTKeys //Master keys: the extended private keys <privKey, chaincode>
	masterChildKeys *EXTKeys //Master child keys: m/0' the extended private keys <privKey, chaincode>
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

func  NewKeystore(length int) *Keystore {

	seed, err := GenSeed(length)
	if err != nil {
		fmt.Errorf("%s should have been nil",err.Error())
	}

	fmt.Printf("The seed is:\n%d\n", seed)

	//Calculate the master key and obtain the master extended private keys from the seed
	extpriv_master := MasterKey(seed)
	fmt.Printf("The extended private key is:\n%s\n", extpriv_master)

	//Derive the m/0' keys: the hardened model based on extended private keys number: 0x80000000
	extpriv_masterchild, err := extpriv_master.Child(0)
	if err != nil {
		fmt.Errorf("%s should have been nil",err.Error())
	}
	fmt.Printf("The master's  child m/0' private key is:\n%s\n", extpriv_masterchild)

	//Build the root identity using offline pathway m/0'/0
	rootID := NewHIDS(0, "m/0'", 1, extpriv_masterchild)
	idData := []HIDS{*rootID}

	return &Keystore{seed, extpriv_master, extpriv_masterchild, idData}
}


// String and Serialized functions: Converting keystore struct to string so that ID info could be stored in keystore file
// String returns the base58-encoded string form of the keystore.
func (ks *Keystore) String() string  {
	return base58.Encode(ks.Serialize())
}

// Serialize returns the serialized form of the keystore.
func (ks *Keystore) Serialize() []byte  {
	depth := uint16ToByte(uint16(w.Depth % 256))
	//bindata = vbytes||depth||fingerprint||i||chaincode||key
	mk := *ks.masterKeys
	mck := *ks.masterChildKeys
	masterkeysdata := append(mk.Vbytes,append(depth,append(mk.Fingerprint,append(mk.I,append(mk.Chaincode,mk.Key...)...)...)...)...)
	masterchildkeysdata := append(mck.Vbytes,append(depth,append(mck.Fingerprint,append(mck.I,append(mck.Chaincode,mck.Key...)...)...)...)...)

	streamdata := append(ks.seed,append(masterkeysdata, append(masterchildkeysdata, append())) )
	bindata := append(w.Vbytes,append(depth,append(w.Fingerprint,append(w.I,append(w.Chaincode,w.Key...)...)...)...)...)
	chksum := dblSha256(bindata)[:4]
	return append(bindata,chksum...)
}


// Generate Single ID information
func NewHIDS(index uint32, pathway string, childrenNum int, parentalEXTKeys *EXTKeys) *HIDS {
	//String operation: turn m/0' into m/0'/childrenNum-1
	childpathway := pathway + "/" + strconv.Itoa(childrenNum - 1)

	SIDData := make(map[string] *IDInfo)
	SIDData[childpathway] = NewIDInfo(index, parentalEXTKeys)


	//Need to modify the number of children
	return &HIDS{SIDData, childrenNum - 1}
}

//Using index number and the parental extended keys to encapsulate a single identity
func NewIDInfo(index uint32,parentalEXTKeys *EXTKeys) *IDInfo {
	// Offline extended private keys
	childEXTPrivF, err := parentalEXTKeys.Child(index)
	if err != nil {
		fmt.Errorf("%s should have been nil",err.Error())
	}
	fmt.Printf("The New Offline ID extended private key is:\n%s\n", childEXTPrivF)
	//Online extended private keys
	childEXTPrivOn, err := parentalEXTKeys.Child(index + uint32(0x80000000))
	if err != nil {
		fmt.Errorf("%s should have been nil",err.Error())
	}
	fmt.Printf("The New Online ID extended private key is:\n%s\n", childEXTPrivOn)

	// Convert the offline private key to public key and derive the identitifer
	extpub := childEXTPrivF.Pub()
	identifier := hash160(privToPub(extpub.Key))
	fmt.Printf("The New ID identifier is:\n%d\n", identifier)

	//Assembly the online and offline extended private keys
	credentials := []EXTKeys{*childEXTPrivF, *childEXTPrivOn}

	return &IDInfo{identifier, credentials}
}



