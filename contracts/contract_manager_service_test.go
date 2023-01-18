package contracts

import (
	"contract-service/signing"
	"contract-service/storage"
	"contract-service/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var claimAbi_Flattened = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_artieAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dispensary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_withdrawalAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"artie\",\"outputs\":[{\"internalType\":\"contractIArtie\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dispensary\",\"outputs\":[{\"internalType\":\"contractIDispensary\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"nonce\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintArtie\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"nonce\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintDispensary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setSignerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"givenWithdrawalAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawalAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]`"

var claimAbi = `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_artieAddress",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_dispensary",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_signer",
				"type": "address"
			},
			{
				"internalType": "address payable",
				"name": "_withdrawalAddress",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "artie",
		"outputs": [
			{
				"internalType": "contract IArtie",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "disable",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "dispensary",
		"outputs": [
			{
				"internalType": "contract IDispensary",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "enable",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "enabled",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes16",
				"name": "nonce",
				"type": "bytes16"
			},
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "mintArtie",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes16",
				"name": "nonce",
				"type": "bytes16"
			},
			{
				"internalType": "uint256",
				"name": "tokenId",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "mintDispensary",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_signer",
				"type": "address"
			}
		],
		"name": "setSignerAddress",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "givenWithdrawalAddress",
				"type": "address"
			}
		],
		"name": "setWithdrawalAddress",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "signingAddress",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes16",
				"name": "",
				"type": "bytes16"
			}
		],
		"name": "usedNonces",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawEth",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawalAddress",
		"outputs": [
			{
				"internalType": "address payable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

var testAbi = `[
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "artieCharAddress",
				"type": "address"
			},
			{
				"internalType": "address payable",
				"name": "withdrawAddress",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "signer",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "current",
				"type": "uint256"
			}
		],
		"name": "Season01Mint",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "MAX_TOKEN",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "PURCHASE_LIMIT",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "artie",
		"outputs": [
			{
				"internalType": "contract Artie",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "current",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes16",
				"name": "nonce",
				"type": "bytes16"
			},
			{
				"internalType": "uint256",
				"name": "numberOfTokens",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "transactionNumber",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "mint",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "price",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "saleStarted",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "signer",
				"type": "address"
			}
		],
		"name": "setSignerAddress",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address payable",
				"name": "givenWithdrawalAddress",
				"type": "address"
			}
		],
		"name": "setWithdrawalAddress",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "signingAddress",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "startSale",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "stopSale",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes16",
				"name": "",
				"type": "bytes16"
			}
		],
		"name": "usedNonces",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawEth",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "withdrawalAddress",
		"outputs": [
			{
				"internalType": "address payable",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

/*
//Contract used for testing:
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Tester {
    function TestInputsA(uint a) public pure returns (uint) {
        return a;
    }
    function TestInputsB(uint8 b) public pure returns (uint8) {
        return b;
    }
    function TestInputsC(uint16 c) public pure returns (uint16) {
        return c;
    }
    function TestInputsD(uint32 d) public pure returns (uint32) {
        return d;
    }
    function TestInputsE(uint64 e) public pure returns (uint64) {
        return e;
    }
    function TestInputsF(uint256 f) public pure returns (uint256) {
        return f;
    }
    function TestInputsG(int g) public pure returns (int) {
        return g;
    }
    function TestInputsH(int8 h) public pure returns (int8) {
        return h;
    }
    function TestInputsI(int16 i) public pure returns (int16) {
        return i;
    }
    function TestInputsJ(int32 j) public pure returns (int32) {
        return j;
    }
    function TestInputsK(int64 k) public pure returns (int64) {
        return k;
    }
    function TestInputsL(address l) public pure returns (address) {
        return l;
    }
    function TestInputsM(bool m) public pure returns (bool) {
        return m;
    }
    function TestInputsN(bytes24 n) public pure returns (bytes24) {
        return n;
    }
    function TestInputsO(bytes calldata o) public pure returns (bytes calldata) {
        return o;
    }
    function TestInputsP(bytes8 p) public pure returns (bytes8) {
        return p;
    }
    function TestInputsQ(bytes16 q) public pure returns (bytes16) {
        return q;
    }
    function TestInputsR(bytes4 r) public pure returns (bytes4) {
        return r;
    }
    function TestInputsS(bytes32 s) public pure returns (bytes32) {
        return s;
    }
} */

var fullTestAbi =`[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "a",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsA",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint8",
				"name": "b",
				"type": "uint8"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsB",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint16",
				"name": "c",
				"type": "uint16"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsC",
		"outputs": [
			{
				"internalType": "uint16",
				"name": "",
				"type": "uint16"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "d",
				"type": "uint32"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsD",
		"outputs": [
			{
				"internalType": "uint32",
				"name": "",
				"type": "uint32"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint64",
				"name": "e",
				"type": "uint64"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsE",
		"outputs": [
			{
				"internalType": "uint64",
				"name": "",
				"type": "uint64"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "f",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsF",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "int256",
				"name": "g",
				"type": "int256"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsG",
		"outputs": [
			{
				"internalType": "int256",
				"name": "",
				"type": "int256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "int8",
				"name": "h",
				"type": "int8"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsH",
		"outputs": [
			{
				"internalType": "int8",
				"name": "",
				"type": "int8"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "int16",
				"name": "i",
				"type": "int16"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsI",
		"outputs": [
			{
				"internalType": "int16",
				"name": "",
				"type": "int16"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "int32",
				"name": "j",
				"type": "int32"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsJ",
		"outputs": [
			{
				"internalType": "int32",
				"name": "",
				"type": "int32"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "int64",
				"name": "k",
				"type": "int64"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsK",
		"outputs": [
			{
				"internalType": "int64",
				"name": "",
				"type": "int64"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "l",
				"type": "address"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsL",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bool",
				"name": "m",
				"type": "bool"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsM",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes24",
				"name": "n",
				"type": "bytes24"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsN",
		"outputs": [
			{
				"internalType": "bytes24",
				"name": "",
				"type": "bytes24"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes",
				"name": "o",
				"type": "bytes"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsO",
		"outputs": [
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes8",
				"name": "p",
				"type": "bytes8"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsP",
		"outputs": [
			{
				"internalType": "bytes8",
				"name": "",
				"type": "bytes8"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes16",
				"name": "q",
				"type": "bytes16"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsQ",
		"outputs": [
			{
				"internalType": "bytes16",
				"name": "",
				"type": "bytes16"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes4",
				"name": "r",
				"type": "bytes4"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsR",
		"outputs": [
			{
				"internalType": "bytes4",
				"name": "",
				"type": "bytes4"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "TestInputsS",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	}
]`


func TestContractManagerService_UnpackArgs(t *testing.T) {
	s := signing.NewSigningService()
	key, _, keyErr := s.GenerateKey()
	assert.Nil(t, keyErr)

	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(testAbi))
	assert.Nil(t, err)
	nonce, nonceErr := utils.GetNonce()
	assert.Nil(t, nonceErr)
	nonceDecoded, decodeErr := hex.DecodeString(nonce[2:])
	assert.Nil(t, decodeErr)
	arguments, byteArgs, argumentsErr := cms.UnpackArgs( [][]byte{nonceDecoded, []byte("3"), []byte("1")}, abiDef.Methods["mint"], storage.Function{Arguments: []storage.Argument{
		{Name: "nonce", Type: "bytes16"},
		{Name: "msg.sender", Type: "address"},
		{Name: "numberOfTokens", Type: "uint256"},
		{Name: "transactionNumber", Type: "uint256"},
	}})
	assert.Nil(t, argumentsErr)
	_, signature, signingErr := s.SignTxn(key, byteArgs)
	assert.Nil(t, signingErr)
	decodedSignature, decodeSigErr := hex.DecodeString(signature[2:])
	assert.Nil(t, decodeSigErr)
	fmt.Println(byteArgs)

	arguments = append(arguments, decodedSignature)
	fmt.Println(arguments)
	fmt.Printf("Signature byte length: %d\n", len(decodedSignature))



	packed, packingErr := abiDef.Pack("mint", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
	fmt.Printf("Length of packed bytes: %d\n", len(packed))

	fmt.Println(string(packed))
	packedHex := hex.EncodeToString(packed)
	fmt.Println(packedHex)

	decodedSig, sigDecodeErr := hex.DecodeString(packedHex[:8])
	assert.Nil(t, sigDecodeErr)
	fmt.Println(decodedSig)
	method, methodErr := abiDef.MethodById(decodedSig)
	assert.Nil(t, methodErr)

	decodedData, decodedErr := hex.DecodeString(packedHex[8:])
	assert.Nil(t, decodedErr)

	unpacked, unpackErr := method.Inputs.Unpack(decodedData)
	assert.Nil(t, unpackErr)
	fmt.Println(unpacked)

	unpackedSignature := hex.EncodeToString((unpacked[len(unpacked) - 1]).([]byte))
	fmt.Println(unpackedSignature)
	fmt.Println(signature)
	assert.Equal(t, signature[2:], unpackedSignature)

}

func TestContractManagerService_UnpackArgsA(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsA"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsA", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsB(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsB"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsB", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsC(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsC"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsC", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsD(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsD"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsD", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsE(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsE"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsE", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsF(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsF"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsF", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsG(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsG"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	fmt.Println(arguments)
	packed, packingErr := abiDef.Pack("TestInputsG", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsH(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsH"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsH", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsI(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsI"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsI", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsJ(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsJ"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsJ", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsK(t *testing.T) {
	args := [][]byte{[]byte("10")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsK"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsK", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsL(t *testing.T) {
	args := [][]byte{common.HexToAddress("0xE2A7f3ADb39C5b11Acb35c02A80ea977D67E1ebc").Bytes()}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsL"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsL", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsM(t *testing.T) {
	args := [][]byte{[]byte("true")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsM"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsM", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsN(t *testing.T) {
	args := [][]byte{[]byte("abcdefghijklmnopqrstuvwx")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsN"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsN", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsO(t *testing.T) {
	args := [][]byte{[]byte("Hey whats up nerds")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsO"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsO", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsP(t *testing.T) {
	args := [][]byte{[]byte("12345678")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsP"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsP", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsQ(t *testing.T) {
	args := [][]byte{[]byte("abcdefghijklmnop")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsQ"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsQ", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsR(t *testing.T) {
	args := [][]byte{[]byte("1234")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsR"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsR", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
func TestContractManagerService_UnpackArgsS(t *testing.T) {
	args := [][]byte{[]byte("abcdefghijklmnopqrstuvwxyz123456")}
	cms := &ContractManagerService{}
	abiDef, err := abi.JSON(strings.NewReader(fullTestAbi))
	assert.Nil(t, err)

	arguments, byteArgs, argumentsErr := cms.UnpackArgs(args, abiDef.Methods["TestInputsS"], storage.Function{})
	assert.Nil(t, argumentsErr)
	fmt.Println(byteArgs)
	arguments = append(arguments, []byte("signature :)"))
	packed, packingErr := abiDef.Pack("TestInputsS", arguments...)
	assert.Nil(t, packingErr)
	fmt.Println(packed)
}
