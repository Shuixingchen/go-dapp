// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.8.6;

// solc --abi contracts/erc20.sol
// abigen --abi=abi/erc20.abi --pkg=erc20 --out=artificial/erc20/erc20.go

abstract contract  ERC20 {
    string public constant name = "";
    string public constant symbol = "";
    uint8 public constant decimals = 0;

    function totalSupply() external view virtual returns (uint);
    function balanceOf(address tokenOwner) external view virtual returns (uint balance);
    function allowance(address tokenOwner, address spender) external view virtual returns (uint remaining);
    function transfer(address to, uint tokens) external virtual returns (bool success);
    function approve(address spender, uint tokens) external virtual returns (bool success);
    function transferFrom(address from, address to, uint tokens) external virtual returns (bool success);
    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}