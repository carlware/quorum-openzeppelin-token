pragma solidity ^0.5.0;

import "openzeppelin-solidity/contracts/token/ERC777/ERC777.sol";


/**
 * @title Simple777Token
 * @dev Very simple ERC777 Token example, where all tokens are pre-assigned to the creator.
 * Note they can later distribute these tokens as they wish using `transfer` and other
 * `ERC20` or `ERC777` functions.
 * Based on https://github.com/OpenZeppelin/openzeppelin-solidity/blob/master/contracts/examples/SimpleToken.sol
 */
contract Simple777Token is ERC777 {

/**
 * @dev Constructor that gives msg.sender all of existing tokens.
 */
				constructor () public ERC777("Simple777Token", "S7", new address[](0)) {
								_mint(msg.sender, msg.sender, 10000 * 10 ** 18, "", "");
				}
}
