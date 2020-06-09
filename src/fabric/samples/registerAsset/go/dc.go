package main

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/sfl0r3nz05/HFB-Testing-Environment/src/fabric/utils/go/m"
)

type DeviceContract interface {
	Init(shim.ChaincodeStubInterface) sc.Response
	Invoke(shim.ChaincodeStubInterface) sc.Response
	AddAsset(shim.ChaincodeStubInterface, []string) sc.Response
	GetAsset(shim.ChaincodeStubInterface, []string) sc.Response
	Synchro() sc.Response
}

// Define the Device Contract structure
type ChainCode struct {
	DeviceContract
}

func NewDeviceContract() DeviceContract {
	return new(ChainCode)
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (cc *ChainCode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(m.OK)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (cc *ChainCode) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "AddAsset" {
		return cc.AddAsset(APIstub, args)
	} else if function == "GetAsset" {
		return cc.GetAsset(APIstub, args)
	} else if function == "Synchro" {
		return cc.Synchro()
	}

	return shim.Error("Invalid Smart Contract function name.")
}

//This is the main smart contract of system
func (cc *ChainCode) AddAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	rfid := args[0]
	hash_ := args[1]
	r := m.Resource{Timestamp: time.Now().Unix(), hashToString: hash_}

	// put k-v to DB
	err := APIstub.PutState(rfid, r.ToBytes())
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(m.OK)
}

func (cc *ChainCode) GetAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	b, _ := APIstub.GetState(args[0])
	return shim.Success(b)
}

func (cc *ChainCode) Synchro() sc.Response {
	return shim.Success(m.OK)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	// Create a new Smart Contract
	err := shim.Start(NewDeviceContract())
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
