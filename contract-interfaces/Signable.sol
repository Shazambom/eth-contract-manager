pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

library Signing {
    using ECDSA for bytes32;


    function verify (bytes32 hash, bytes calldata signature, address validator) internal pure returns (bool) {
        return hash.recover(signature) == validator;
    }

    function hashRequestWithSender(bytes memory args) internal view returns(bytes32) {
        return keccak256(abi.encodePacked(msg.sender, msg.value, args)).toEthSignedMessageHash();
    }

    function hashRequest(bytes memory args) internal view returns(bytes32) {
        return keccak256(abi.encodePacked(msg.value, args)).toEthSignedMessageHash();
    }

}

interface Signable {
    modifier validSender(bytes memory args, bytes calldata signature, address signingAddress) {
        require(Signing.verify(Signing.hashRequestWithSender(args), signature, signingAddress), "INVALID SIGNATURE");
        _;
    }

    modifier valid(bytes memory args, bytes calldata signature, address signingAddress) {
        require(Signing.verify(Signing.hashRequest(args), signature, signingAddress), "INVALID SIGNATURE");
        _;
    }
}
