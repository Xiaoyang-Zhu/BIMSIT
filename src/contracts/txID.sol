pragma solidity ^0.4.0;


contract txID {

    address public owner;

    struct Voter {
        uint weight;
        bool voted;
        uint8 vote;
        address delegate;
    }

    mapping(uint => auction) Auctions;

    //Initialization Function: establish the identity contracts using identity information delivered by HID generation
    function txReg(){
        owner = msg.sender;



    }

    function txUPD(){

    }

    function txRVK(){

    }

    function txLKP(){

    }

}
