package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net"
)

// message format:
// header byte
// get/put/whatever.

const (
	GEODE_NULL   byte = 41
	GEODE_STRING byte = 42
	GEODE_HEADER byte = 110
)
const (
	GET_REQUEST      uint32 = 0
	RESPONSE         uint32 = 1
	PARTIAL_RESPONSE uint32 = 2
	PUT_REQUEST      uint32 = 7
	PUTALL           uint32 = 56
	GETALL           uint32 = 57
	EXECUTE_FUNCTION uint32 = 59
)

func main() {
	conn, err := net.Dial("tcp", "poland.gemstone.com:40404")
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

// Request to put to a region.
func putRequest(writer io.Writer, region string, key string, value string) {
	messageBody := newPutRequest(region, key, value)
	// 17 byte message header
	// int: 32 bits, signed -- all ints
	// first 17 bytes:
	// [ message type | message len | number of parts | txnID || 1 byte flags ]
	// txnId -1 for now
	// flags 0 for now
	bodyLen := messageBody.Len()

	binary.Write(writer, binary.BigEndian, int32(GEODE_PUT))
	binary.Write(writer, binary.BigEndian, int32(bodyLen)+17)
	binary.Write(writer, binary.BigEndian, int32(7))
	binary.Write(writer, binary.BigEndian, int32(-1))
	writer.Write([]byte{0})
	bodyBytes, err := messageBody.WriteTo(writer)
	if err != nil {
		panic(err)
	} else {
		if bodyBytes != int64(bodyLen) {
			panic(fmt.Sprintf("wrong number of bytes? %d %d", bodyBytes, bodyLen))
		}
	}
}

// 7 parts:
// 1. region
// 2. key
// 3. value
func newPutRequest(region string, key string, value string) *bytes.Buffer {
	// Make the message body because we need len to send the message.
	messageBody := new(bytes.Buffer)
	writeStringPart(messageBody, region)

	messageBody.WriteByte(GEODE_NULL)
	binary.Write(messageBody, binary.BigEndian, 0)
	writeStringPart(messageBody, key)
	binary.Write(messageBody, binary.BigEndian, 0)
	writeStringPart(messageBody, value)
	messageBody.Write([]byte{GEODE_NULL})

	return messageBody
}

// for each part, we need length plus one boolean flag (0 for now)
// Two bytes of length, then null, followed by modified UTF-8.
// len as short
//func writeString(Writer conn, String message) {
func writeStringPart(writer io.Writer, message string) {
	//if len(message) > 65536 {
	//panic("Message too long!")
	//}
	len := uint16(len(message))
	//
	//binary.Write(writer, binary.BigEndian, len)
	//writer.Write([]byte{0})
	writer.Write([]byte(message))
}
