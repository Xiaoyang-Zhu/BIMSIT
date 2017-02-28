package keystore

import (
	"fmt"
	"btcsuite/btcutil/base58"
	"bytes"
	"errors"
)

//Single identity information
type IDInfo struct {
	Identifier []byte //160 bytes
	Credentials []EXTKeys //The extended public or private keys: the first ExtKey is offline keys; the second is online
	ChildrenNum	uint32
//	Attributes []byte //Attributes
}

//Entire FAKE tree structure of one identity
//Credentials in each IDInfo are composed by extended private keys which can be used to derive the public keys
type HIDS struct {
	SIDData map[string] *IDInfo //string stands for the pathway of identity
	index	uint32	//The number of identities derived from one seed
//	description string
}

//Identity information storage file encrypted by users' defined password: it's a secured vault of entire identity info
type Keystore struct {
//	secret string
	seed []byte // Seed
	masterKeys *EXTKeys //Master keys: the extended private keys <privKey, chaincode> serialized string 78 Bytes
	masterChildKeys *EXTKeys //Master child keys: m/0' the extended private keys <privKey, chaincode> serialized string 78 Bytes
	idData HIDS // All identity info
}

//Online Identity information set (IDS): the online HIDS
//Credentials in each IDInfo are composed by extended public keys rather than the extended private keys
type OHIDS struct {
	OSIDData map[string] IDInfo
//	description string
}


func  NewKeystore(length uint16) *Keystore {

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
	idData := NewHIDS(0, "m/0'", 1, extpriv_masterchild)

	return &Keystore{seed, extpriv_master, extpriv_masterchild, *idData}
}


// String and Serialized functions: Converting keystore struct to string so that ID info could be stored in keystore file
// String returns the base58-encoded string form of the keystore.
func (ks *Keystore) String() string  {
	return base58.Encode(ks.Serialize())
}

// Serialize returns the serialized form of the keystore.
func (ks *Keystore) Serialize() []byte  {
	//streamdata = seed||masterkeys||masterchildkeys||array of HIDS
	mk := ks.masterKeys.Serialize()
	mck := ks.masterChildKeys.Serialize()
	idata := ks.idData.Serialize()
	seedlen := uint32ToByte(uint32(len(ks.seed)))
	mklen := uint32ToByte(uint32(len(mk)))
	mcklen := uint32ToByte(uint32(len(mck)))
	idatalen := uint32ToByte(uint32(len(idata)))
	fmt.Println("The bytes number of seed/masterkey/masterchildkey/idata:", seedlen, mklen, mcklen, idatalen)
	streamdata := append(seedlen, append(ks.seed, append(mklen, append(mk, append(mcklen, append(mck, append(idatalen, idata...)...)...)...)...)...)...)
//	streamdata := append(ks.seed, append(mk, append(mck, idata...)...)...)
	chksum := dblSha256(streamdata)[:4]
	return append(streamdata,chksum...)
}


//SIDData map[string] *IDInfo
//ChildrenNum	uint32
func (hids *HIDS) Serialize() []byte {
	var bsiddata []byte
	for key, value := range hids.SIDData {
//		fmt.Printf("%d : %s\n", key, value)
		bkey := []byte(key)
		bvalue := value.Serialize()
		bsiddata = append(bkey, bvalue...)
	}
	bchildren := uint32ToByte(uint32(hids.index))

	return append(bsiddata, bchildren...)


}
//Identifier []byte
//Credentials []EXTKeys
//ChildrenNum uint32
func (idinfo *IDInfo) Serialize() []byte {
	return append(idinfo.Identifier, append(idinfo.Credentials[0].Serialize(), append(idinfo.Credentials[1].Serialize(), uint32ToByte(uint32(idinfo.ChildrenNum))...)...)...)
}


// StringKeystore returns a Keystore struct given a base58-encoded string
func StringKeystore(data string) (*Keystore,error) {
	ks := base58.Decode(data)
	if err := ByteCheck(ks); err != nil {
		return &Keystore{}, err
	}
	if bytes.Compare(dblSha256(ks[:(len(ks)-4)])[:4], ks[(len(ks)-4):]) != 0 {
		return &Keystore{}, errors.New("Invalid checksum")
	}
	seedlen := ks[0:4]
	
	depth := byteToUint16(dbin[4:5])
	fingerprint := dbin[5:9]
	i := dbin[9:13]
	chaincode := dbin[13:45]
	key := dbin[45:78]
	return &Keystore{vbytes, depth, fingerprint, i, chaincode, key}, nil
}

// Generate Single ID information
func NewHIDS(index uint32, pathway string, childrenNum uint32, parentalEXTKeys *EXTKeys) *HIDS {
	//String operation: turn m/0' into m/0'/childrenNum-1
	childpathway := pathway + "/0"

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

	return &IDInfo{identifier, credentials, 0}
}



