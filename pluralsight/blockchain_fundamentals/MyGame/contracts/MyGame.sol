pragma solidity ^0.4.4;
contract IScoreStore {
    function GetScore(string name) returns (int);
}

contract MyGame {
    function ShowScore(string name) returns (int) {
        IScoreStore scoreStore = IScoreStore(0x16fe544b61365498dc5c7353e3be33e514c858f7);
        return scoreStore.GetScore(name);
    }
    
}