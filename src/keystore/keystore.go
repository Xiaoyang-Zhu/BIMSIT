package keystore

import (
	"fmt"
	"bytes"
	"errors"
	"encoding/binary"

	"github.com/btcsuite/btcutil/base58"
	"log"
)

// Basic Struct Definition

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

// Three Keystore Struct Global Constructor: NewKeystore, NewHIDS and NewIDInfo

func  NewKeystore(length uint16) *Keystore {

	seed, err := GenSeed(length)
	if err != nil {
		log.Fatalf("%s should have been nil",err.Error())
	}

	fmt.Printf("The seed is:\n%d\n", seed)

	//Calculate the master key and obtain the master extended private keys from the seed
	extpriv_master := MasterKey(seed)
	fmt.Printf("The extended private key is:\n%s\n", extpriv_master)
	fmt.Println(extpriv_master.Serialize())

	//Derive the m/0' keys: the hardened model based on extended private keys number: 0x80000000
	extpriv_masterchild, err := extpriv_master.Child(0)
	if err != nil {
		log.Fatalf("%s should have been nil",err.Error())
	}
	fmt.Printf("The master's  child m/0' private key is:\n%s\n", extpriv_masterchild)
	fmt.Println(extpriv_masterchild.Serialize())

	//Build the root identity using offline pathway m/0'/0
	idData := NewHIDS(0, "m/0'", 0, extpriv_masterchild)

	return &Keystore{seed, extpriv_master, extpriv_masterchild, *idData}
}

// Generate entire hierarchical ID information
func NewHIDS(index uint32, pathway string, childrenNum uint32, parentalEXTKeys *EXTKeys) *HIDS {
	//String operation: turn m/0' into m/0'/childrenNum
	childpathway := fmt.Sprintf("%s/%d", pathway, childrenNum)

	SIDData := make(map[string] *IDInfo)
	SIDData[childpathway] = NewIDInfo(index, parentalEXTKeys)

	//Need to modify the number of children
	return &HIDS{SIDData, index}
}

//Using index number and the parental extended keys to encapsulate a single identity
func NewIDInfo(index uint32,parentalEXTKeys *EXTKeys) *IDInfo {
	// Offline extended private keys
	childEXTPrivF, err := parentalEXTKeys.Child(index)
	if err != nil {
		log.Fatalf("%s should have been nil",err.Error())
	}
	fmt.Printf("The New Offline ID extended private key is:\n%s\n", childEXTPrivF)
	fmt.Println(childEXTPrivF.Serialize())

	//Online extended private keys
	childEXTPrivOn, err := parentalEXTKeys.Child(index + uint32(0x80000000))
	if err != nil {
		log.Fatalf("%s should have been nil",err.Error())
	}
	fmt.Printf("The New Online ID extended private key is:\n%s\n", childEXTPrivOn)
	fmt.Println(childEXTPrivOn.Serialize())

	// Convert the offline private key to public key and derive the identitifer
	extpub := childEXTPrivF.Pub()
	identifier := hash160(privToPub(extpub.Key))
	fmt.Printf("The New ID identifier is:\n%d\n", identifier)
	fmt.Printf("The New ID identifier string is:\n%x\n", identifier)

	//Assembly the online and offline extended private keys
	credentials := []EXTKeys{*childEXTPrivF, *childEXTPrivOn}

	return &IDInfo{identifier, credentials, 0}
}

// Four Keystore Struct Serializing Functions: Converting keystore struct into string so that ID info could be stored in keystore file

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
	//streamdata := append(ks.seed, append(mk, append(mck, idata...)...)...)
	chksum := dblSha256(streamdata)[:4]
	return append(streamdata,chksum...)
}

func (hids *HIDS) Serialize() []byte {
	var bsiddata []byte
	var idnum uint32
	for key, value := range hids.SIDData {
		fmt.Printf("%s : %d\n", key, value)
		bkey := []byte(key)
		bkeylen := uint32ToByte(uint32(len(bkey)))
		bvalue := value.Serialize()
		fmt.Printf("The IDInfo binary value: \n%d\n", bvalue)
		bvaluelen := uint32ToByte(uint32(len(bvalue)))
		idnum = idnum + 1
		bsiddata = append(bsiddata, append(bkeylen, append(bkey, append(bvaluelen, bvalue...)...)...)...)
	}
	index := uint32ToByte(uint32(hids.index))

	fmt.Printf("The iterator value: \n%d\n", bsiddata)

	return append(uint32ToByte(idnum), append(bsiddata, index...)...)


}

func (idinfo *IDInfo) Serialize() []byte {
	var IDstr []byte
	for _, value := range idinfo.Credentials {
		cred := value.Serialize()
		credlen := uint32ToByte(uint32(len(cred)))
		IDstr = append(IDstr, append(credlen, cred...)...)
	}

	return append(idinfo.Identifier, append(IDstr, uint32ToByte(uint32(idinfo.ChildrenNum))...)...)

	//credoffline := idinfo.Credentials[0].Serialize()
	//credofflinelen := uint32ToByte(uint32(len(credoffline)))
	//credonline := idinfo.Credentials[1].Serialize()
	//credonlinelen := uint32ToByte(uint32(len(credonline)))
	//return append(idinfo.Identifier, append(credofflinelen, append(credoffline, append(credonlinelen, append(credonline, uint32ToByte(uint32(idinfo.ChildrenNum))...)...)...)...)...)
}

// Four Keystore File Deserializing Functions: StringKeystore, DeserializeEXTKeys, DeserializeHIDS, DeserializeIDInfo

// StringKeystore returns a Keystore struct given a base58-encoded string
func StringKeystore(data string) (*Keystore,error) {
	//BASE58 Decoding and check checksum value
	ks := base58.Decode(data)
	if bytes.Compare(dblSha256(ks[:(len(ks)-4)])[:4], ks[(len(ks)-4):]) != 0 {
		return &Keystore{}, errors.New("Invalid checksum")
	}

	//Obtain seed
	seedlen := binary.BigEndian.Uint32(ks[0:4])
	seed := ks[4:4 + seedlen]

	//Obtain masterkey and masterchildkey
	mklen := binary.BigEndian.Uint32(ks[4 + seedlen:8 + seedlen])
	mkstr := ks[8 + seedlen:8 + seedlen + mklen]
	mcklen := binary.BigEndian.Uint32(ks[8 + seedlen + mklen:12 + seedlen + mklen])
	mckstr := ks[12 + seedlen + mklen:12 + seedlen + mklen + mcklen]

	//Obtain identity information string
	iddatalen := binary.BigEndian.Uint32(ks[12 + seedlen + mklen + mcklen:16 + seedlen + mklen + mcklen])
	iddatastr := ks[16 + seedlen + mklen + mcklen:16 + seedlen + mklen + mcklen + iddatalen]

	//fmt.Printf("keystore is: \n%d\n seed is: \n%d\n masterkey is: \n%d\n masterkey child is:\n%d\n iddata is: \n%d\n", ks, seed, mkstr, mckstr, iddatastr)


	mk, err := DeserializeEXTKeys(mkstr)
	if err != nil {
		log.Fatalln("Errors in analyzing master keys")
		return &Keystore{}, err
	}
	mck, err := DeserializeEXTKeys(mckstr)
	if err != nil {
		log.Fatalln("Errors in analyzing master child keys")
		return &Keystore{}, err
	}

	iddata, err := DeserializeHIDS(iddatastr)
	if err != nil {
		log.Fatalln("Errors in analyzing HIDS struct")
		return &Keystore{}, err
	}
	return &Keystore{seed, mk, mck, *iddata}, nil
}

//DeserializeEXTKeys returns EXTKeys struct pointer -- a extended private keys given a base58-encoded extended key
func DeserializeEXTKeys(extkeystr []byte) (*EXTKeys,error) {
	vbytes := extkeystr[0:4]
	depth := byteToUint16(extkeystr[4:5])
	fingerprint := extkeystr[5:9]
	i := extkeystr[9:13]
	chaincode := extkeystr[13:45]
	key := extkeystr[45:78]
	return &EXTKeys{vbytes, depth, fingerprint, i, chaincode, key}, nil
}

//DeserializeHIDS returns HIDS struct pointer -- the entire identity struct given a []byte string
func DeserializeHIDS(hids []byte) (*HIDS, error) {

	var base uint32
	idnum := binary.BigEndian.Uint32(hids[0:4])
	SIDData := make(map[string] *IDInfo)

	for idnum > 0 {
		pathstrlen := binary.BigEndian.Uint32(hids[4 + base:8 + base])
		pathstrb := hids[8 + base:8 + pathstrlen + base]
		idinfostrlen := binary.BigEndian.Uint32(hids[8 + pathstrlen + base:12 + pathstrlen + base])
		idinfostr := hids[12 + pathstrlen + base: 12 + pathstrlen + idinfostrlen + base]
		idinfo, err := DeserializeIDInfo(idinfostr)
		if err != nil {
			log.Fatalln("Errors in analyzing identity structure")
			return &HIDS{}, err
		}

		//Convert []byte type pathway into pathway string
		pathstr := string(pathstrb[:pathstrlen])
		SIDData[pathstr] = idinfo
		base = base + 12 + pathstrlen + idinfostrlen

		idnum--
	}

	index := binary.BigEndian.Uint32(hids[base:4 + base])

	return &HIDS{SIDData, index}, nil
}

//DeserializeIDInfo returns IDInfo struct pointer -- a single identity struct given a []byte string
func DeserializeIDInfo(idinfostr []byte) (*IDInfo,error) {
	//Get identifier
	identifier := idinfostr[0:20]

	//Get offline and online credentials
	credofflinelen := binary.BigEndian.Uint32(idinfostr[20:24])
	credofflinestr := idinfostr[24:24 + credofflinelen]
	credoffline, err := DeserializeEXTKeys(credofflinestr)
	if err != nil {
		log.Fatalln("Errors in analyzing master child keys")
		return &IDInfo{}, err
	}
	credonlinelen := binary.BigEndian.Uint32(idinfostr[24 + credofflinelen:28 + credofflinelen])
	credonlinestr := idinfostr[28 + credofflinelen:28 + credofflinelen + credonlinelen]
	credonline, err := DeserializeEXTKeys(credonlinestr)
	if err != nil {
		log.Fatalln("Errors in analyzing master child keys")
		return &IDInfo{}, err
	}
	//Get children number
	childrenNum := binary.BigEndian.Uint32(idinfostr[28 + credofflinelen + credonlinelen:32 + credofflinelen + credonlinelen])
	credentials :=[]EXTKeys{*credoffline, *credonline}

	return &IDInfo{identifier, credentials, childrenNum}, nil

}


// Keystore Struct Operation Functions

// String returns the base58-encoded string form of the keystore.
func (ks *Keystore) ListAllIDPath() {
	fmt.Println("All identities' path:")
	for key, _ := range ks.idData.SIDData {
		fmt.Printf("%s\n", key)
	}
	return
}

// AddNewIDKeystore return a new Keystore struct given one parent path like "m/0'/0"
func (ks *Keystore) AddNewIDKeystore(parentalPath string) (*Keystore, error) {

	//Get parental identity information
	pid := ks.idData.SIDData[parentalPath]
	fmt.Printf("The ID info is :\n%d\n", pid.Serialize())

	//String operation: turn m/0' into m/0'/childrenNum
	childpathway := fmt.Sprintf("%s/%d", parentalPath, pid.ChildrenNum)

	//Pass the parental identity info, get a new IDInfo struct and put the string and struct into map
	fmt.Println(pid.Credentials)
	ks.idData.SIDData[childpathway] = NewIDInfo(ks.idData.index, &pid.Credentials[0])

	//Increase the number of identities in HIDS
	ks.idData.index ++

	//Recreate the parental identity and put the new value into the map
	ks.idData.SIDData[parentalPath] = pid.ModifyChildNum()

	fmt.Printf("The modified new keystore deserialized result is:\n%d\n", ks.idData)

	return &Keystore{ks.seed, ks.masterKeys, ks.masterChildKeys, ks.idData}, nil
}

// Modify the children number of parent
func (sid *IDInfo) ModifyChildNum() *IDInfo {
	return &IDInfo{sid.Identifier, sid.Credentials, sid.ChildrenNum + 1}
}

// GetRootIDInfo returns root identity information to construct the identity smart contract.
func (ks *Keystore) GetRootIDInfo() (rootID, rootPKf, rootPKo []byte, err error) {
	rootID = ks.idData.SIDData["m/0'/0"].Identifier
	for counter, value := range ks.idData.SIDData["m/0'/0"].Credentials {
		if counter == 0 {
			rootPKf = value.Pub().Key
		} else if counter == 1 {
			rootPKo = value.Pub().Key
		} else {
			log.Fatalln("Has more than two credential keys")
			return
		}
	}

	return rootID, rootPKf, rootPKo, nil
}



