// Copyright 2017 Rolevax. All rights reserved.

// Package sp4g implements functions to read and write
// simple "size + payload" messages.
//
// The size header takes 4 bytes thus the payload can be up to 4 GiB.
// The 4-byte-header is encoded in big endian.
package sp4g

import (
	"encoding/binary"
	"errors"
	"io"
	"strconv"
)

// Read is just read.
func Read(reader io.Reader) ([]byte, error) {
	size, err := readSize(reader)
	if err != nil {
		return nil, err
	}
	return readPayload(reader, size)
}

// ReadN reads up to n bytes
func ReadN(reader io.Reader, n uint32) ([]byte, error) {
	size, err := readSize(reader)
	if err != nil {
		return nil, err
	}

	if size > n {
		return nil, errors.New("sp4g: ReadN size blown")
	}

	return readPayload(reader, size)
}

func readSize(reader io.Reader) (uint32, error) {
	var size uint32
	err := binary.Read(reader, binary.BigEndian, &size)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func readPayload(reader io.Reader, size uint32) ([]byte, error) {
	p := make([]byte, size)
	_, err := io.ReadFull(reader, p)
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
		return errors.New("sp4g: too big data" + strconv.Itoa(l))
	}

	size := uint32(len(data))
	err := binary.Write(conn, binary.BigEndian, size)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)
	return err
}
