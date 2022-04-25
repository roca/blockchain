package main

import "fmt"

type TestChaincode struct {
}

func (t *TestChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *TestChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"query\"")
}

func main() {
	fmt.Printf("hello, world\n")
}
