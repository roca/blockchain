pragma solidity ^0.4.4;

contract ScoreStore {
    mapping (string => int) PersonScore;
    function AddPersonScore(string name, int startingScore) {
        if (PersonScore[name]>0) {
            throw;
        } else {
           PersonScore[name] = startingScore;
        }
    }

    function GetScore(string name) returns (int) {
       return PersonScore[name]; 
    }
}