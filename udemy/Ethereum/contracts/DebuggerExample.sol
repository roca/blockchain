// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

contract DebuggerExample {
	uint public myUint;

	function setUint(uint _myUint) public {
		myUint = _myUint;
	}
}