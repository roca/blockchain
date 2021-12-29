// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;
// In solidity ^0.5.13 uint8 as a wraparound effect when incremented passes 255;

contract WorkingWithVariables {
	uint256 public myInt;

	function setMyUint(uint256 _myInt) public {
		myInt = _myInt;
	}

	bool public myBool;

	function setMyBool(bool _myBool) public {
		myBool = _myBool;
	}

	uint8 public myUint8;

	function incrementUint() public {
		myUint8++;
	}

	function decrementUint() public {
		myUint8--;
	}

}
