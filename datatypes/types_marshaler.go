package datatypes

import (
	"io"
	"strconv"

	"github.com/devansh42/pastadb/utils"
)

var (
	_true         = []byte{'t'}
	_false        = []byte{'f'}
	_nil          = []byte{utils.TypeNil}
	_bulkString   = []byte{utils.TypeBulkString}
	_simpleString = []byte{utils.TypeSimpleString}
	_simpleErr    = []byte{utils.TypeSimpleErr}
	_integer      = []byte{utils.TypeInteger}
	_boolean      = []byte{utils.TypeBoolean}
	_array        = []byte{utils.TypeArray}
	_map          = []byte{utils.TypeMap}
	_set          = []byte{utils.TypeSet}
)

func (s SimpleString) Marshal(w io.Writer) {
	w.Write(_simpleString)
	w.Write(s)
	addTerm(w)
}

func (s BulkString) Marshal(w io.Writer) {
	w.Write(_bulkString)
	l := strconv.FormatInt(int64(len(s)), 10)
	w.Write([]byte(l))
	addTerm(w)
	w.Write(s)
	addTerm(w)
}

func (err SimpleErr) Marshal(w io.Writer) {
	w.Write(_simpleErr)
	w.Write(err)
	addTerm(w)
}

func (in Integer) Marshal(w io.Writer) {
	w.Write(_integer)
	i := strconv.FormatInt(int64(in), 10)
	w.Write([]byte(i))
	addTerm(w)
}

func (b Boolean) Marshal(w io.Writer) {
	w.Write(_boolean)
	if b {
		w.Write(_true)
	} else {
		w.Write(_false)
	}
	addTerm(w)
}

func (Nil) Marshal(w io.Writer) {
	w.Write(_nil)
	addTerm(w)
}

func (ar Array) Marshal(w io.Writer) {
	l := utils.List(ar)
	marshalSetorArrayorMap(l, w, _array, int64(l.Len()))
}

func (set Set) Marshal(w io.Writer) {
	l := utils.List(set)
	marshalSetorArrayorMap(l, w, _set, int64(l.Len()))

}

func (mp Map) Marshal(w io.Writer) {
	l := utils.List(mp)
	// We are halving map length
	// since we have key value pair
	marshalSetorArrayorMap(l, w, _map, int64(l.Len()/2))
}

func marshalSetorArrayorMap(l utils.List, w io.Writer, typ []byte, length int64) {
	w.Write(typ)
	lens := strconv.FormatInt(length, 10)
	w.Write([]byte(lens))
	addTerm(w)

	// Iterating it's elements
	for n := l.Head(); n != nil; n = n.Next() {
		marshler := n.Val().(RespMarshaler)
		// Writting elements to the buffer
		marshler.Marshal(w)
	}
}

func addTerm(w io.Writer) {
	w.Write(utils.Teminator)
}
