package requests

import "io"

type GetRequest struct {
	Header     RequestHeader
	RegionName PackedString
	Key        PackedString
}

func DoGetRequest(w io.Writer, key string, value string) {

}
