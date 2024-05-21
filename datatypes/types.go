package datatypes

import (
	"io"

	"github.com/devansh42/pastadb/utils"
)

// Datatypes declaration
type (
	SimpleString []byte
	BulkString   []byte
	SimpleErr    []byte
	Integer      int64
	Boolean      bool
	Array        utils.List
	Set          utils.List
	Map          utils.List
	Nil          struct{}
)

type RespMarshaler interface {
	Marshal(io.Writer)
}
