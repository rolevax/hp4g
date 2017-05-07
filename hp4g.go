// Copyright 2017 Rolevax. All rights reserved.

// Package hp4g implements functions to read and write
// simple "size + payload" messages.
//
// The size header takes 4 bytes thus the payload can be up to 4 GiB.
// The 4-byte-header is encoded in big endian.
package hp4g

import (
	"encoding/binary"
	"errors"
	"io"
	"strconv"
)

// Read is just read.
func Read(reader io.Reader) ([]byte, error) {
	var size uint32
	err := binary.Read(reader, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}

	p := make([]byte, size)
	_, err = io.ReadFull(reader, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Write is just write. Does nothing when data is empty.
func Write(conn io.Writer, data []byte) error {
	l := len(data)
	if l == 0 {
		return nil
	} else if l >= (1 << 32) {
		return errors.New("hp4g: too big data" + strconv.Itoa(l))
	}

	size := uint32(len(data))
	err := binary.Write(conn, binary.BigEndian, size)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)
	return err
}
