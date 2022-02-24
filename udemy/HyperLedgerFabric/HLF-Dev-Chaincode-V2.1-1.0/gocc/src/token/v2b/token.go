package main

/**
 * tokenv2
 * Shows the
 *    A) Use of Logger
 **/
import (
	"fmt"

	// April 2020, Updated for Fabric 2.0
	"github.com/hyperledger/fabric-chaincode-go/shim"
	
	peer "github.com/hyperledger/fabric-protos-go/peer"

	// Used for formatting the timestamp
	"time"

	// acloudFan custom Logger
	"acflogger"
)

var logger = acflogger.NewLogger()

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Invoke executed ")

	// V3   Print the transaction ID
	//fmt.Srintf("GetTxID() => %s\n", stub.GetTxID())

	//log a info GetTxID()
	logger.Info(fmt.Sprintf("GetTxID() => %s\n", stub.GetTxID()))

	// V3   Print the channel ID
	//fmt.Println("GetChannelID() =>", stub.GetChannelID())

	//log a info GetChannelID()
	logger.Info(fmt.Sprintf("GetChannelID() => %s\n", stub.GetChannelID()))

	// V3   Print the transaction Timestamp
	TxTimestamp, _ := stub.GetTxTimestamp()
	timeStr := time.Unix(TxTimestamp.GetSeconds(),0)
	//fmt.Printf("GetTxTimestamp() => %s\n", timeStr)

	//log a info GetTxTimestamp()
	logger.Info(fmt.Sprintf("GetTxTimestamp() => %s\n", timeStr))

	// Extract the information from proposal 
	PrintSignedProposalInfo(stub)

	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Println("Started Chaincode. token/v3")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
