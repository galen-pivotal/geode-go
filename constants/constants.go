package constants

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
