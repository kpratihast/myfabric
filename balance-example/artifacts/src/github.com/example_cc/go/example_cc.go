package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type PolicyChaincode struct {
}

type PolicyDetails struct {
	Aadhar           string `json:aadharNumber`
	Name             string `json:name`
	Age              string `json:age`
	PolicyStatus     string `json:status`
	InsuranceCompany string `json:insCompany`
	InsuranceType    string `json:insType`
}

// ===================================================================================
// Main
// ===================================================================================

func main() {
	err := shim.Start(new(PolicyChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *PolicyChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *PolicyChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	switch function {
	case "createPolicy":

		return t.createPolicy(stub, args)
	case "getPolicy":

		return t.getPolicy(stub, args)
	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")
	}
}

func (t *PolicyChaincode) createPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	result, err := stub.GetState(args[0])
	myPolicies := []PolicyDetails{}
	if err == nil {
		fmt.Println(result)
		if result != nil {
			_ = json.Unmarshal([]byte(result), &myPolicies)

			fmt.Println("result")
			fmt.Println(result)
			fmt.Println("result from getstate is nil")
		} else {
			fmt.Println("error")
		}
	} else {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Args ", args)
	fmt.Println("Inside createPolicy function")
	if len(args) < 6 {
		fmt.Println("Incorrect number of arguments. Expecting 8")
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	fmt.Println("policy details", myPolicies)

	aadharNumber := args[0]
	name := args[1]
	age := args[2]
	status := args[3]
	insCompany := args[4]
	insType := args[5]

	PolicyObject := &PolicyDetails{aadharNumber, name, age, status, insCompany, insType}
	fmt.Println("policy object", PolicyObject)
	myPolicies = append(myPolicies, *PolicyObject)
	PolicyDetailsJSONAsObject, err := json.Marshal(myPolicies)
	if err != nil {
		return shim.Error(err.Error())
	}

	//storing PolicyDetails Array Object against AaadharNumber as a key in ledger
	err = stub.PutState(aadharNumber, []byte(PolicyDetailsJSONAsObject))
	if err != nil {
		fmt.Println("Could not save Policy Data on blockchain ledger.", err)
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *PolicyChaincode) getPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var aadharNumber = args[0]
	PolicyDetails, err := stub.GetState(aadharNumber)
	if err != nil {
		return shim.Error(err.Error())
	}
	if PolicyDetails == nil {
		return shim.Error("Could not find Policy.")
	}

	return shim.Success(PolicyDetails)
}

