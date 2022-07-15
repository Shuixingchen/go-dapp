// SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.8.6;

// solc --abi contracts/erc721.sol
// abigen --abi=abi/erc721.abi --pkg=erc721 --out=artificial/erc721/erc721.go

interface ERC721 {

    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
    event TransferOld(address from, address to, uint256 tokenId);
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);
    // erc1155的event
    event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value);
    event TransferBatch(address indexed operator,address indexed from,address indexed to,uint256[] ids,uint256[] values);
    event URI(string value, uint256 indexed id);
    // OpenSea的sale
    event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint price, bytes32 indexed metadata);
    event OrderCancelled(bytes32 indexed hash);
    function atomicMatch_(address[14] memory addrs,uint[18] memory uints,uint8[8] memory feeMethodsSidesKindsHowToCalls,bytes memory calldataBuy,bytes memory calldataSell,bytes memory replacementPatternBuy,bytes memory replacementPatternSell,bytes memory staticExtradataBuy,bytes memory staticExtradataSell,uint8[2] memory vs,bytes32[5] memory rssMetadata)external payable; 
    
    function balanceOf(address _owner) external pure returns (uint256);
    function ownerOf(uint256 _tokenId) external pure returns (address);
    
    function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes memory data) external payable;
    function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable;
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) external;
    
    function approve(address _approved, uint256 _tokenId) external payable;
    function getApproved(uint256 _tokenId) external view returns (address);
    function setApprovalForAll(address _operator, bool _approved) external;    
    function isApprovedForAll(address _owner, address _operator) external view returns (bool);

    function name() external pure returns (string memory);
    function symbol() external pure returns (string memory);
    function tokenURI(uint256 _tokenId) external view returns (string memory);
    function uri(uint256 id) external view returns (string memory);

    // erc1155的nft
    function balanceOfBatch(address[] calldata accounts, uint256[] calldata ids) external view returns (uint256[] memory);
    function safeTransferFrom(address from,address to,uint256 id,uint256 amount,bytes calldata data) external;
    function safeBatchTransferFrom(address from,address to,uint256[] calldata ids,uint256[] calldata amounts,bytes calldata data) external;

    // 额外的一些方法
    function totalSupply() external view returns (uint);
    // 铸造NFT
    function awardItem(address player, string memory tokenURI) external returns (uint256);
    // erc165 判断是否实现interfaceId
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
}