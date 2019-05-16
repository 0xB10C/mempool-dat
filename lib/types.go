package lib

import (
	"fmt"

	"github.com/btcsuite/btcd/wire"
)

/*  FileHeader  */

// FileHeader represents the mempool file header
type FileHeader struct {
	version int64
	numTx   int64
}

// returns the version and the number of transactions in the file
func (header FileHeader) String() string {
	return fmt.Sprintf("Version %d, contains %d transactions", header.version, header.numTx)
}

// GetVersion returns the mempool file version
func (header FileHeader) GetVersion() int64 {
	return header.version
}

// GetTxCount returns the number of transactions in the corresponding file
func (header FileHeader) GetTxCount() int64 {
	return header.numTx
}

/*  MempoolEntry  */

// MempoolEntry represents a mempool entry
type MempoolEntry struct {
	transaction   *wire.MsgTx
	firstSeenTime int64
	feeDelta      int64
}

// returns the transaction hash of the entry
func (entry MempoolEntry) String() string {
	return entry.transaction.TxHash().String()
}

// GetFirstSeen returns the firstSeen time of the entry as timestamp
func (entry MempoolEntry) GetFirstSeen() int64 {
	return entry.firstSeenTime
}

// GetFeeDelta returns feeDelta of the entry
func (entry MempoolEntry) GetFeeDelta() int64 {
	return entry.feeDelta
}

// Info returns a string with information for a given MempoolEntry
func (entry MempoolEntry) Info() string {
	firstSeen := entry.firstSeenTime
	numInputs := len(entry.transaction.TxIn)
	numOutputs := len(entry.transaction.TxOut)
	isSegWit := entry.transaction.HasWitness()
	hash := entry.transaction.TxHash()

	return fmt.Sprintf("txid: %v, in: %d, out: %d, firstSeen: %d, isSegWit %t", hash, numInputs, numOutputs, firstSeen, isSegWit)
}

/*  Mempool  */

// Mempool represents a parsed mempool.dat file
type Mempool struct {
	header    FileHeader
	entries   []MempoolEntry
	mapDeltas []byte // not parsed
}

// GetMempoolEntries returns a slice with mempool entries
func (mempool Mempool) GetMempoolEntries() []MempoolEntry {
	return mempool.entries
}

// GetFileHeader returns a the mempool file header
func (mempool Mempool) GetFileHeader() FileHeader {
	return mempool.header
}

// GetMapDeltas returns a byte slice of not parsed mapDelta entries
func (mempool Mempool) GetMapDeltas() []byte {
	return mempool.mapDeltas
}
