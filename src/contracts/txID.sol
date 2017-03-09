pragma solidity ^0.4.0;


contract txID {

    address public owner;

    struct IDreceipt {
        address identifier;
        bytes pkf;
        bytes pko;
        bytes sigma_skf;
        bytes pointer;
    }

    mapping(address => IDreceipt) identity;

    //Initialization Function: establish the identity contracts using identity information delivered by HID generation
    function txReg(address rootID, bytes rootPKf, bytes rootPKo, bytes sig, bytes rootPointer) {
        //transafer the money then conduct the following operations
        owner = msg.sender;

        //Initialize the value
        identity[rootID].pkf = rootPKf;
        identity[rootID].pko = rootPKo;
        identity[rootID].sigma_skf = sig;
        identity[rootID].pointer = rootPointer;

    }

    //HOWTO define the miners verification?
    function verifier() {

    }

    function txUPD() {

    }

    function txRVK() {

    }

    function txLKP() {

    }

}
