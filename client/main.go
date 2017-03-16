package main

import (
	"encoding/hex"
	"io"
	"math"
	"net"

	"github.com/lunixbochs/struc"
)

// message format:
// header byte
// get/put/whatever.

// object types
const (
	NULL          byte = 41
	UTF_STRING    byte = 42
	BYTE_ARRAY    byte = 43
	SHORT_ARRAY   byte = 44
	INTEGER_ARRAY byte = 45
	LONG_ARRAY    byte = 46
	FLOAT_ARRAY   byte = 47
	DOUBLE_ARRAY  byte = 48

	STRING_ARRAY byte = 64

	ARRAY     byte = 52
	BOOLEAN   byte = 53
	CHARACTER byte = 54
	BYTE      byte = 55
	SHORT     byte = 56
	INTEGER   byte = 57
	LONG      byte = 58
	FLOAT     byte = 59
	DOUBLE    byte = 60

	SET byte = 66
	MAP byte = 67

	DATA_SERIALIZATION byte = 37
	PDX_SERIALIZATION  byte = 93
	USER_SERIALIZATION byte = 40
)

const GEODE_HEADER byte = 110

// request types
const (
	GET_REQUEST      = 0
	RESPONSE         = 1
	PARTIAL_RESPONSE = 2
	PUT_REQUEST      = 7
	PUTALL           = 56
	GETALL           = 57
	EXECUTE_FUNCTION = 59
)

////type Header struct {
//Size        int32
//RequestType uint8
//Version     uint8
//RequestID   int32
//flag        uint8
//}

// Java modified UTF-8: unsigned short len. hope we don't run into the "modified" part.

func main() {
	conn, err := net.Dial("tcp", "localhost:40404")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte{GEODE_HEADER})
	putRequest(conn, "exampleRegion", "testKey", "testValue")

	response := make([]byte, 1000)
	bytes, err := conn.Read(response)
	if err != nil {
		panic(err)
	}
	if bytes != 0 {
		hex.EncodeToString(response)
	}
}

// getMessage().addXPart() is what we seem to do.

// format for the entire message of a put:
// 1. Region name
// 2. Op : GEODE_NULL
// 3. flags : int == 0
// 4. key: String
// 5. bool false : one byte
// 6. value: String
// 7. EventID: GEODE_null

// modified UTF-8 because java yay!

// request: 4 + 2 + 4 + 1 = 11 bytes
// size         int32
// RequestType  byte
// Version      byte
// RequestID    int32
// flag         byte

type RequestHeader struct {
	Size        int32
	RequestType int16
	Version     int16
	RequestId   int32
	Flag        uint8
}

type PackedString struct {
	Size  uint16 
	Value []byte
}

type PutRequest struct {
	Header      RequestHeader
	RegionName  PackedString
	Key         PackedString
	ValueHeader int32
	Value       PackedString
}

func (putRequest PutRequest) writeTo (w io.Writer) {
	w.Write(putRequest.Header
}

func packString(s string) PackedString {
	b := []byte(s)
	if len(b) > math.MaxUint16 {
		panic("string too long.")
	}
	return PackedString{uint16(len(b)), b}
}

// Request to put to a region.
func putRequest(writer io.Writer, region string, key string, value string) {
	//messageBody := newPutRequest(region, key, value)
	// 17 byte message header
	// int: 32 bits, signed -- all ints
	// first 17 bytes:
	// [ message type | message len | number of parts | txnID || 1 byte flags ]
	// txnId -1 for now
	// flags 0 for now
	//bodyLen := messageBody.Len()

	request := PutRequest{
		Header:     RequestHeader{Size: 1000 /* FIXME */, RequestType: PUT_REQUEST, Version: 110 /*fixme*/, RequestId: 1, Flag: 0},
		RegionName: packString(region),
		Key:        packString(key),
		Value:      packString(value),
	}

	err := struc.Pack(writer, request)

	if err != nil {
		panic(err)
	}
}
