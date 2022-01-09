// SPDX-License-Identifier: MIT
pragma solidity ^0.8.11;

contract helloworld {
    function say() public pure returns (string memory) {
        return "hello etherworld";
    }

    function sayName(string memory name) public pure returns (string memory) {
        return name;
    }
}
