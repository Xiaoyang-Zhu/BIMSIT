package keystore

import (
    "bytes"
    "crypto/hmac"
    "crypto/sha512"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "math/big"

    "github.com/btcsuite/btcutil/base58"
    )

var (
    //MainNet
    Public []byte
    Private []byte
    //TestNet
    TestPublic []byte
    TestPrivate []byte
    )

func init() {
    Public,_ = hex.DecodeString("0488B21E")
    Private,_ = hex.DecodeString("0488ADE4")
    TestPublic,_ = hex.DecodeString("043587CF")
    TestPrivate,_ = hex.DecodeString("04358394")
}

// EXTKeys defines the components of a hierarchical deterministic wallet: 78 Bytes
type EXTKeys struct {
    Vbytes []byte //4 bytes: version bytes (mainnet: 0x0488B21E public, 0x0488ADE4 private; testnet: 0x043587CF public, 0x04358394 private)
    Depth uint16 //1 byte: depth: 0x00 for master nodes, 0x01 for level-1 derived keys, ....
    Fingerprint []byte //4 bytes: the fingerprint of the parent's key (0x00000000 if master key)
    I []byte //4 bytes: child number
    Chaincode []byte //32 bytes: the chain code
    Key []byte //33 bytes: the public key or private key data
}

// Child returns the ith child of wallet w. Values of i >= 2^31
// signify private key derivation. Attempting private key derivation
// with a public key will throw an error.
func (w *EXTKeys) Child(i uint32) (*EXTKeys,error) {
    var fingerprint, I , newkey []byte
    switch {
    case bytes.Compare(w.Vbytes, Private) == 0, bytes.Compare(w.Vbytes, TestPrivate) == 0:
        pub := privToPub(w.Key)
        mac := hmac.New(sha512.New, w.Chaincode)
        if i >= uint32(0x80000000) {
            mac.Write(append(w.Key,uint32ToByte(i)...))
        } else {
            mac.Write(append(pub,uint32ToByte(i)...))
        }
        I = mac.Sum(nil)
         iL := new(big.Int).SetBytes(I[:32])
        if iL.Cmp(curve.N) >= 0 || iL.Sign() == 0 {
            return &EXTKeys{}, errors.New("Invalid Child")
        }
        newkey = addPrivKeys(I[:32], w.Key)
        fingerprint = hash160(privToPub(w.Key))[:4]

    case bytes.Compare(w.Vbytes, Public) == 0, bytes.Compare(w.Vbytes, TestPublic) == 0:
        mac := hmac.New(sha512.New, w.Chaincode)
        if i >= uint32(0x80000000) {
            return &EXTKeys{}, errors.New("Can't do Private derivation on Public key!")
        }
        mac.Write(append(w.Key,uint32ToByte(i)...))
        I = mac.Sum(nil)
        iL := new(big.Int).SetBytes(I[:32])
        if iL.Cmp(curve.N) >= 0 || iL.Sign() == 0 {
            return &EXTKeys{}, errors.New("Invalid Child")
        }
        newkey = addPubKeys(privToPub(I[:32]), w.Key)
        fingerprint = hash160(w.Key)[:4]
    }
    return &EXTKeys{w.Vbytes, w.Depth + 1, fingerprint, uint32ToByte(i), I[32:], newkey}, nil
}

// Serialize returns the serialized form of the wallet.
func (w *EXTKeys) Serialize() []byte  {
    depth := uint16ToByte(uint16(w.Depth % 256))
    //bindata = vbytes||depth||fingerprint||i||chaincode||key
    bindata := append(w.Vbytes,append(depth,append(w.Fingerprint,append(w.I,append(w.Chaincode,w.Key...)...)...)...)...)
    chksum := dblSha256(bindata)[:4]
    return append(bindata,chksum...)
}

// String returns the base58-encoded string form of the wallet.
func (w *EXTKeys) String() string  {
    return base58.Encode(w.Serialize())
}

// StringWallet returns a wallet given a base58-encoded extended key
func StringWallet(data string) (*EXTKeys,error) {
    dbin := base58.Decode(data)
    if err := ByteCheck(dbin); err != nil {
        return &EXTKeys{}, err
    }
    if bytes.Compare(dblSha256(dbin[:(len(dbin)-4)])[:4], dbin[(len(dbin)-4):]) != 0 {
        return &EXTKeys{}, errors.New("Invalid checksum")
    }
    vbytes := dbin[0:4]
    depth := byteToUint16(dbin[4:5])
    fingerprint := dbin[5:9]
    i := dbin[9:13]
    chaincode := dbin[13:45]
    key := dbin[45:78]
    return &EXTKeys{vbytes, depth, fingerprint, i, chaincode, key}, nil
}

// Pub returns a new wallet which is the public key version of w.
// If w is a public key, Pub returns a copy of w
func (w *EXTKeys) Pub() *EXTKeys {
    if bytes.Compare(w.Vbytes,Public) == 0 {
        return &EXTKeys{w.Vbytes, w.Depth, w.Fingerprint, w.I, w.Chaincode, w.Key}
    } else {
        return &EXTKeys{Public, w.Depth, w.Fingerprint, w.I, w.Chaincode, privToPub(w.Key)}
    }
}

// StringChild returns the ith base58-encoded extended key of a base58-encoded extended key.
func StringChild(data string ,i uint32) (string, error) {
    w, err := StringWallet(data)
    if err != nil {
        return "", err
    } else {
        w, err = w.Child(i)
        if err != nil {
            return "", err
        } else {
            return w.String(), nil
        }
    }
}

//StringToAddress returns the Bitcoin address of a base58-encoded extended key.
func StringAddress(data string) (string, error) {
    w, err := StringWallet(data)
    if err != nil {
        return "", err
    } else {
        return w.Address(), nil
    }
}

// Address returns bitcoin address represented by wallet w.
func (w *EXTKeys) Address() string {
    x, y := expand(w.Key)
    four,_ := hex.DecodeString("04")
    padded_key := append(four,append(x.Bytes(),y.Bytes()...)...)
    var prefix []byte
    if bytes.Compare(w.Vbytes,TestPublic) == 0 || bytes.Compare(w.Vbytes,TestPrivate) == 0 {
        prefix,_ = hex.DecodeString("6F")
    } else {
        prefix,_ = hex.DecodeString("00")
    }
    addr_1 := append(prefix,hash160(padded_key)...)
    chksum := dblSha256(addr_1)
    return base58.Encode(append(addr_1,chksum[:4]...))
}

// GenSeed returns a random seed with a length measured in bytes.
// The length must be at least 128.
func GenSeed(length uint16) ([]byte, error) {
    b := make([]byte, length/8)
    if length < 128 {
        return b, errors.New("length must be at least 128 bits")
    }
    _, err := rand.Read(b)
    return b, err
}

// MasterKey returns a new wallet given a random seed.
func MasterKey(seed []byte) *EXTKeys {
    key := []byte("Bitcoin seed")
    mac := hmac.New(sha512.New, key)
    mac.Write(seed)
    I := mac.Sum(nil)
    secret := I[:len(I)/2]
    chain_code := I[len(I)/2:]
    depth := 0
    i := make([]byte, 4) //2^32 children
    fingerprint := make([]byte, 4)
    zero := make([]byte,1)
    return &EXTKeys{Private,uint16(depth),fingerprint,i,chain_code,append(zero,secret...)}
}

// StringCheck is a validation check of a base58-encoded extended key.
func StringCheck(key string) error {
    return ByteCheck(base58.Decode(key))
}

func ByteCheck(dbin []byte) error{
    // check proper length
    if len(dbin) != 82 {
        return errors.New("invalid string")
    }
    // check for correct Public or Private vbytes
    if bytes.Compare(dbin[:4],Public) != 0 && bytes.Compare(dbin[:4],Private) != 0 && bytes.Compare(dbin[:4],TestPublic) != 0 && bytes.Compare(dbin[:4],TestPrivate) != 0 {
        return errors.New("invalid string")
    }
    // if Public, check x coord is on curve
    x, y := expand(dbin[45:78])
    if bytes.Compare(dbin[:4],Public) == 0 || bytes.Compare(dbin[:4],TestPublic) == 0 {
        if !onCurve(x,y) {
            return errors.New("invalid string")
        }
    }
    return nil
}
