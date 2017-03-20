package main

import (
	"fmt"

	"github.com/galen-pivotal/geode-go/client/requests"
)

func main() {
	conn, err := requests.ConnectGeode("localhost:40404")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	requests.DoPutRequest(conn, "exampleRegion", "testKey", "testValue")

	response := make([]byte, 1000)
	bytesRead := 0
	for bytesRead < 9 {
		bytes, err := conn.Read(response[bytesRead:])
		if err != nil {
			panic(err)
		}
		bytesRead = bytesRead + bytes
	}

	response = response[:bytesRead]
	fmt.Printf("Response length %d into string of length %d, string %q\n", bytesRead, len(response), response)
}

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
