// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.8.6;
contract TestContract
{
    mapping(address => uint256) balances;
    function balanceOf(address tokenOwner) public view returns (uint) {
        require(tokenOwner != address(0), "owner not exist");
        return balances[tokenOwner];
    }
}