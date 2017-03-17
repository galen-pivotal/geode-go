package requests

import (
	"encoding/binary"
	"io"
	"math"
	"net"

	"github.com/galen-pivotal/geode-go/constants"
)

func ConnectGeode(hostname string) (io.ReadWriteCloser, error) {
	conn, err := net.Dial("tcp", hostname)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write([]byte{constants.GEODE_HEADER})
	if err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}

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

func (requestHeader RequestHeader) writeTo(w io.Writer) {
	binary.Write(w, binary.BigEndian, requestHeader.Size)
	binary.Write(w, binary.BigEndian, requestHeader.RequestType)
	binary.Write(w, binary.BigEndian, requestHeader.Version)
	binary.Write(w, binary.BigEndian, requestHeader.RequestId)
	binary.Write(w, binary.BigEndian, requestHeader.Flag)
}

func (packedString PackedString) writeTo(w io.Writer) {
	binary.Write(w, binary.BigEndian, packedString.Size)
	w.Write(packedString.Value)
}

func (putRequest PutRequest) writeTo(w io.Writer) {
	putRequest.Header.writeTo(w)
	putRequest.RegionName.writeTo(w)
	putRequest.Key.writeTo(w)
	binary.Write(w, binary.BigEndian, putRequest.ValueHeader)
	putRequest.Value.writeTo(w)
}

func packString(s string) PackedString {
	b := []byte(s)
	if len(b) > math.MaxUint16 {
		panic("string too long.")
	}
	return PackedString{uint16(len(b)), b}
}

// Request to put to a region.
func DoPutRequest(writer io.Writer, region string, key string, value string) {
	//messageBody := newPutRequest(region, key, value)
	// 17 byte message header
	// int: 32 bits, signed -- all ints
	// first 17 bytes:
	// [ message type | message len | number of parts | txnID || 1 byte flags ]
	// txnId -1 for now
	// flags 0 for now
	//bodyLen := messageBody.Len()

	request := PutRequest{
		Header:     RequestHeader{Size: 1000 /* FIXME */, RequestType: constants.PUT_REQUEST, Version: 110 /*fixme*/, RequestId: 1, Flag: 0},
		RegionName: packString(region),
		Key:        packString(key),
		Value:      packString(value),
	}

	request.writeTo(writer)

	//if err != nil {
	//panic(err)
	//}
}
