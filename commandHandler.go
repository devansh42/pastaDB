package main

import (
	"bytes"

	"github.com/devansh42/pastadb/datatypes"
)

var (
	_pingCommand = []byte("ping")
	_PONGMsg     = datatypes.SimpleString("PONG")
)

func handleCommand(args []datatypes.BulkString, cli *clientConn) {

	if len(args) > 0 {
		switch {
		case bytes.EqualFold(args[0], _pingCommand):
			if len(args) > 1 {
				// Ping command with a message
				handlePingCommand(args[1:], cli)
				return
			}
			// Ping command w/o a message
			handlePingCommand(nil, cli)

		}

	}

}

func handlePingCommand(args []datatypes.BulkString, cli *clientConn) {
	if len(args) > 0 {
		// Ping with a message
		// We need to send that message in the reply
		// as a BulkString as per Redis specification
		msg := args[0]
		msg.Marshal(cli.conn)
	} else {
		// Ping w/o a message
		// We need to send PONG in the reply
		// as a simple string as per Redis specification

		_PONGMsg.Marshal(cli.conn)
	}
}
