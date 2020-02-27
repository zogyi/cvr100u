package dev

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Opeartion string
type FieldType string

const (
	OpeartionAuth        Opeartion = `CVR_Authenticate`
	OpeartionReadContent Opeartion = `CVR_Read_Content`
	OperationCloseComm   Opeartion = `CVR_CloseComm`
	ReadName             FieldType = `GetPeopleName`
	ReadNation           FieldType = `GetPeopleNation`
	ReadBirthday         FieldType = `GetPeopleBirthday`
	ReadAddress          FieldType = `GetPeopleAddress`
	ReadIDCode           FieldType = `GetPeopleIDCode`
	ReadDepartment       FieldType = `GetDepartment`
	ReadStartDate        FieldType = `GetStartDate`
	ReadEndDate          FieldType = `GetEndDate`
	ReadSex              FieldType = `GetPeopleSex`
	ReadJPG              FieldType = `GetBMPData`
)

//Connector this is the class for invoke the low level method
type Connector struct {
	isInitialed     bool
	ReadyToReadData bool
	DllPath         string
	CardType        uint8
	handle          *syscall.Handle
}

func (conn *Connector) Initial() (successed bool) {
	handle, err := syscall.LoadLibrary(conn.DllPath)
	if err != nil {
		panic("can't find the dll file")
	}
	conn.handle = &handle
	cvrInitComm, err := syscall.GetProcAddress(handle, "CVR_InitComm")
	if err != nil {
		panic("can't get the init method from the dll")
	}
	var nargs uintptr = 1
	ret, _, errNo := syscall.Syscall(uintptr(cvrInitComm), nargs, uintptr(1001), 0, 0)
	if errNo != 0 {
		panic("something wrong has happened when innoke the init method   ")
	}
	if ret != 1 {
		return false
	}
	conn.isInitialed = true
	return true
}

//Authentication authentication
func (conn *Connector) Authentication() (successed bool) {
	var nargs uintptr = 0
	return conn.opertation(OpeartionAuth, nargs, 0)
}

//ReadContent read content
func (conn *Connector) ReadContent() (successed bool) {
	var nargs uintptr = 1
	return conn.opertation(OpeartionReadContent, nargs, uintptr(4))
}

func (conn *Connector) opertation(opType Opeartion, nargs uintptr, arg1 uintptr) (successed bool) {
	if conn.handle == nil || conn.isInitialed == false {
		panic("the connector hasn't been initialized")
	}
	authentication, err := syscall.GetProcAddress(*conn.handle, string(opType))
	if err != nil {
		fmt.Printf("Error------ %s\n", err)
		return
	}
	ret, _, errNo := syscall.Syscall(uintptr(authentication), nargs, arg1, 0, 0)
	if errNo != 0 {
		println("has some thing wrong again")
	}
	if int(ret) != 1 {
		return false
	}
	println(int(ret))
	time.Sleep(300)
	return true
}

//ReadFields read fields
func (conn *Connector) ReadFields(fieldType FieldType) (result string, result1 []byte, err error) {
	if conn.handle == nil || conn.isInitialed == false {
		panic("the connector hasn't been initialized")
	}
	readFieldMethod, err := syscall.GetProcAddress(*conn.handle, string(fieldType))
	if err != nil {
		fmt.Printf("Error------ %s\n", err)
		return result, nil, err
	}
	var nargs uintptr = 1
	var bytesObj = make([]byte, 70)
	if fieldType == ReadJPG {
		bytesObj = make([]byte, 40960)
	}
	var length = len(bytesObj)
	ret, _, errNo := syscall.Syscall(uintptr(readFieldMethod), nargs, uintptr(unsafe.Pointer(&bytesObj[0])), uintptr(unsafe.Pointer(&length)), 0)

	if errNo != 0 || ret != 1 {
		err = fmt.Errorf("can't execute read method of %s", fieldType)
		return result, nil, err
	}
	if fieldType == ReadJPG {
		return result, bytesObj, nil
	}
	n := bytes.IndexByte(bytesObj, 0)
	bytesObj = bytesObj[:n]

	reader := transform.NewReader(bytes.NewReader(bytesObj), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		err = errors.New("can't decode the result when read the file")
		return result, nil, err
	}
	time.Sleep(300)
	return string(d), nil, nil
}
