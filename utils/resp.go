package utils

// Holds construct for RESP Datatypes
const (
	TypeSimpleString byte = '+'
	TypeBulkString   byte = '$'
	TypeSimpleErr    byte = '-'
	TypeInteger      byte = ':'
	TypeArray        byte = '*'
	TypeNil          byte = '_'
	TypeBoolean      byte = '#'
	TypeMap          byte = '%'
	TypeSet          byte = '~'
)

var Teminator = []byte{'\r', '\n'}
