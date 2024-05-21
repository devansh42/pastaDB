package main

import (
	"bytes"
	"errors"
	"log"
	"net"
	"strconv"

	"github.com/devansh42/pastadb/datatypes"
	"github.com/devansh42/pastadb/utils"
)

const (

	// We are ready to read next command from client
	readyToRead byte = 1

	// we are reading argument length and yet to find terminator
	readingArgumentLen byte = 2

	// We are reading argument and yet to find terminator
	readingArgument byte = 3

	// We have read the whole command and ready to execute it
	readyToExecute byte = 4
)
const (
	_CR = '\r'
	_LF = '\n'
)

type clientConn struct {
	conn net.Conn

	// queryBuf represents buffer that directly reads
	// from tcp connection
	// e.g. {*,2,\r,\n,$,5,\r,\n,h,e,l,l,o,\r,\n,$,5,\r,\n,w,o,r,l,d,\r,\n}
	queryBuf []byte

	// arguments represents decoded command arguments
	// e.g. {set,k,v}
	arguments []datatypes.BulkString

	// connReader helps us read from this query buffer
	connReader *bytes.Reader

	//readerState records state of the reader
	readerState byte

	//curArgumentReadUpto records number of bytes
	// read from query buffer for current argument
	curArgumentReadUpto int

	// curArgumentIndex records index of argument being read
	// from query buffer
	curArgumentIndex int
}

const (
	_QueryBufferSize = 4 * 1024 // 4 KB
)

func handleNewConnection(conn net.Conn) {
	var cc = clientConn{
		conn:        conn,
		queryBuf:    make([]byte, _QueryBufferSize),
		readerState: readyToRead,
	}
	cc.connReader = bytes.NewReader(cc.queryBuf)
	startReading(&cc)
}

func startReading(cli *clientConn) {
	for {
		switch cli.readerState {
		case readyToRead:
			handleReadyToRead(cli)
		}
	}
}

func handleReadyToRead(cli *clientConn) {
	readBytes, err := cli.conn.Read(cli.queryBuf)
	if err != nil {
		log.Panic("error while reading from connection: ", err)
	}
	if readBytes == 0 {
		// Idle Connection
	}
	prefixByte, err := cli.connReader.ReadByte()
	if err != nil {
		panic(err)
	}
	if prefixByte == utils.TypeArray {
		// *2\r\n$5\r\nhello\r\n$5\r\nworld\r\n
		var lenArr []byte
		for {
			digit, err := cli.connReader.ReadByte()
			if err != nil {
				panic(err)
			}
			if digit == _CR {
				break
			}

			lenArr = append(lenArr, digit)
		}
		argumentLen, err := strconv.Atoi(string(lenArr))
		if err != nil {
			panic(err)
		}
		cli.arguments = make([]datatypes.BulkString, argumentLen)
		data, err := cli.connReader.ReadByte()
		panicIfRequired(err)
		if data != _LF {
			panicIfRequired(errors.New("line feed not found"))
		}

	}
}

func panicIfRequired(err error) {
	if err != nil {
		panic(err)
	}
}
