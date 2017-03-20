package requests

import (
	"encoding/binary"
	"io"
)

type RequestHeader struct {
	Size        int32
	RequestType int16
	Version     int16
	RequestId   int32
	Flag        uint8
}

func (requestHeader RequestHeader) writeTo(w io.Writer) {
	binary.Write(w, binary.BigEndian, requestHeader.Size)
	binary.Write(w, binary.BigEndian, requestHeader.RequestType)
	binary.Write(w, binary.BigEndian, requestHeader.Version)
	binary.Write(w, binary.BigEndian, requestHeader.RequestId)
	binary.Write(w, binary.BigEndian, requestHeader.Flag)
}
