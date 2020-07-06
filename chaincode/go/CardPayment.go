package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CardPaymentRecordChaincode struct {
}

type cardpaymentrecord struct {
	ObjectType      string `json:"objectType"`
	AP_ID           string `json:"AP_ID"`
	CLASS_ID_UP     string `json:"CLASS_ID_UP"`
	CLASS_ID_DN     string `json:"CLASS_ID_DN"`
	BEG_DT          string `json:"BEG_DT"`
	END_DT          string `json:"END_DT"`
	WEEKDAY         string `json:"WEEKDAY"`
	RESERVE         string `json:"RESERVE"`
	AP_ID_UP        string `json:"AP_ID_UP"`
	AP_ID_UP_PERIOD string `json:"AP_ID_UP_PERIOD"`
	ARRIVAL_TIME    string `json:"ARRIVAL_TIME"`
	AP_ID_DN        string `json:"AP_ID_DN"`
	AP_ID_DN_NEXT   string `json:"AP_ID_DN_NEXT"`
	LEAVE_TIME      string `json:"LEAVE_TIME"`
	SESSION_ID      string `json:"SESSION_ID"`
	SESSION_DT      string `json:"SESSION_DT"`
	RAW_DATA        string `json:"RAW_DATA"`
	FLAG            string `json:"FLAG"`
	cprId           string `json:"cprId"`
}

func (t *CardPaymentRecordChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *CardPaymentRecordChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "initCPR" {
		return t.initCPR(stub, args)
	} else if fn == "readCPR" {
		return t.readCPR(stub, args)
	} else if fn == "deleteCPR" {
		return t.deleteCPR(stub, args)
	} else if fn == "getCPRByRange" {
		return t.getCPRByRange(stub, args)
	} else if fn == "queryCPRByBank" {
		return t.queryCPRByBank(stub, args)
	} else if fn == "getHistoryForCPR" {
		return t.getHistoryForCPR(stub, args)
	} else if fn == "queryCPRByDate" {
		return t.queryCPRByDate(stub, args)
	}
	return shim.Error("没有對應的方法！")

}
func (t *CardPaymentRecordChaincode) initCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 18 {
		return shim.Error("Incorrect number of arguments. Expecting 18")
	}

	objectType := "cardpaymentrecoed"
	AP_ID := args[0]
	CLASS_ID_UP := args[1]
	CLASS_ID_DN := args[2]
	BEG_DT := args[3]
	END_DT := args[4]
	WEEKDAY := args[5]
	RESERVE := args[6]
	AP_ID_UP := args[7]
	AP_ID_UP_PERIOD := args[8]
	ARRIVAL_TIME := args[9]
	AP_ID_DN := args[10]
	AP_ID_DN_NEXT := args[11]
	LEAVE_TIME := args[12]
	SESSION_ID := args[13]
	SESSION_DT := args[14]
	RAW_DATA := args[15]
	FLAG := args[16]
	cprId := args[17]

	// size, err := strconv.Atoi(args[3]) //轉int
	// if err != nil {
	// 	return shim.Error("size must be a numeric string")
	// }
	// ipfsHash := args[4]
	//,Amount,StoreID,StoreCity,FeeIndicator,BillType,PaymentType,SettlementDate,CardType
	cardpaymentrecord := &cardpaymentrecord{objectType, AP_ID, CLASS_ID_UP, CLASS_ID_DN, BEG_DT, END_DT, WEEKDAY, RESERVE, AP_ID_UP, AP_ID_UP_PERIOD, ARRIVAL_TIME, AP_ID_DN, AP_ID_DN_NEXT, LEAVE_TIME, SESSION_ID, SESSION_DT, RAW_DATA, FLAG, cprId}

	CPRJsonAsBytes, err := json.Marshal(cardpaymentrecord)

	err = stub.PutState(cprId, CPRJsonAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *CardPaymentRecordChaincode) readCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	cprId := args[0]

	CPRAsBytes, err := stub.GetState(cprId)

	if err != nil {
		return shim.Error(err.Error())
	} else if CPRAsBytes == nil {
		return shim.Error("Card payment Record not exist")
	}

	return shim.Success(CPRAsBytes)

}

func (t *CardPaymentRecordChaincode) deleteCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	cprId := args[0]

	CPRAsBytes, err := stub.GetState(cprId)

	if err != nil {
		return shim.Error(err.Error())
	} else if CPRAsBytes != nil {

		err = stub.DelState(cprId)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	return shim.Success(nil)

}

// func (t *CardPaymentRecordChaincode) updateCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

// 	cprId := args[0]
// 	// newSize, err := strconv.Atoi(args[1]) //轉int
// 	// if err != nil {
// 	// 	return shim.Error("size must be a numeric string")
// 	// }
// 	// newIPFSHash := args[2]

// 	//判斷紀錄是否存在
// 	CPRAsBytes, err := stub.GetState(cprId)

// 	if err != nil {
// 		return shim.Error(err.Error())
// 	} else if CPRAsBytes == nil {
// 		return shim.Error("Card payment Record not exist")
// 	}

// 	CPRInfo := cardpaymentrecord{}
// 	err = json.Unmarshal(CPRAsBytes, &CPRInfo) //json 轉 struct

// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	//CPRInfo.Size = newSize
// 	//CPRInfo.IPFSHash = newIPFSHash

// 	CPRJsonAsBytes, _ := json.Marshal(CPRInfo)

// 	err = stub.PutState(cprId, CPRJsonAsBytes)

// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return shim.Success(nil)

// }

func (t *CardPaymentRecordChaincode) getCPRByRange(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	startKey := "1"
	endKey := "999"

	resultIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("- getCPRByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

//富查詢
func (t *CardPaymentRecordChaincode) queryCPRByBank(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	bank := args[0]
	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"cardpaymentrecoed\",\"acquiringBank\":\"%s\"}}", bank)

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

//富查詢
func (t *CardPaymentRecordChaincode) queryCPRByDate(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	date := args[0]

	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"cardpaymentrecoed\",\"acquiringDate\":\"%s\"}}", date)

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)

}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *CardPaymentRecordChaincode) getHistoryForCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	cprId := args[0]

	fmt.Printf("- start getHistoryForMarble: %s\n", cprId)

	resultIterator, err := stub.GetHistoryForKey(cprId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	var buffer bytes.Buffer

	buffer.WriteString("[")
	buffer.WriteString(cprId)

	isWrite := false
	for resultIterator.HasNext() {

		queryResponse, err := resultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if isWrite == true {
			buffer.WriteString(",")
		}

		buffer.WriteString("{ \"TxId\":")
		buffer.WriteString(queryResponse.TxId)

		buffer.WriteString(",\"Timestamp\": ")
		buffer.WriteString(time.Unix(queryResponse.Timestamp.Seconds, int64(queryResponse.Timestamp.Nanos)).String())

		buffer.WriteString(",\"Value\": ")
		buffer.WriteString(string(queryResponse.Value))

		buffer.WriteString(",\"IsDelete\": ")
		buffer.WriteString(strconv.FormatBool(queryResponse.IsDelete))
		buffer.WriteString("}")

		isWrite = true
	}

	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func main() {
	err := shim.Start(new(CardPaymentRecordChaincode))
	if err != nil {
		fmt.Println("chaincode start error!")
	}
}
