// Package lib provides mempool file primitives and functionality to load a mempool.dat file.
//
// You'd want to be calling ReadMempoolFromPath( ) to read a mempool.dat file.
// See https://github.com/0xB10C/mempool-dat for some usage examples.
package lib

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"

	"github.com/btcsuite/btcd/wire"
)

// ReadMempoolFromPath reads a mempool file from a given path and returns a Mempool type
func ReadMempoolFromPath(path string) (mem Mempool, err error) {

	file, err := os.Open(path)
	if err != nil {
		return mem, errors.New("Could not read mempool file: " + err.Error())
	}
	defer file.Close()

	r := bufio.NewReader(file)
	header, err := readFileHeader(r)
	if err != nil {
		return mem, errors.New("Could not read header: " + err.Error())
	}
	mem.header = header

	var entries []MempoolEntry
	// for the number of entries specified in the header
	for i := int64(0); i < header.numTx; i++ {
		mementry, err := readMempoolEntry(r)
		if err != nil {
			return mem, errors.New("Could not read mempoolEntry at index " + string(i) + " : " + err.Error())
		}
		entries = append(entries, mementry)
	}
	mem.entries = entries

	var feeDelta []byte
	for {
		remainingBytes, err := r.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return mem, errors.New("Could not read feeDelta: " + err.Error())
		}
		feeDelta = append(feeDelta, remainingBytes)
	}

	return
}

func readFileHeader(r *bufio.Reader) (header FileHeader, err error) {
	fileVersion, err := readLEint64(r)
	if err != nil {
		return header, err
	}

	numberOfTx, err := readLEint64(r)
	if err != nil {
		return header, err
	}

	header = FileHeader{fileVersion, numberOfTx}
	return
}

func readMempoolEntry(r *bufio.Reader) (mementry MempoolEntry, err error) {
	msgTx := wire.NewMsgTx(1) // TODO: check if version 2 is correct
	err = msgTx.Deserialize(r)
	if err != nil {
		return mementry, err
	}

	timestamp, err := readLEint64(r)
	if err != nil {
		return mementry, err
	}

	feeDelta, err := readLEint64(r)
	if err != nil {
		return mementry, err
	}

	mementry = MempoolEntry{msgTx, timestamp, feeDelta}
	return
}

// reads the next 64bit in Little Endian and returns a int64
// used here to get the mempoolEntry's timestamp and feeDelta
func readLEint64(r *bufio.Reader) (res int64, err error) {
	next64Bit := make([]byte, 8, 8)
	_, err = io.ReadFull(r, next64Bit)
	if err != nil {
		return 0, err
	}

	res = int64(binary.LittleEndian.Uint64(next64Bit))
	return
}
