package main

import (
	"fmt"

	mempoolDat "github.com/0xb10c/mempool-dat/lib"
)

const path string = "mempool.dat"

func main() {
	// please handle errors, omitted for brevity 
	mempool, _ := mempoolDat.ReadMempoolFromPath(path)

	// prints the version and number of tx
	fmt.Println(mempool.GetFileHeader())

	// get all mempool entries with GetMempoolEntries (here only the first three are used)
	for _, e := range mempool.GetMempoolEntries()[:3]{
		fmt.Println(e.Info())
	}

	// get the firstSeen timestamp of a transaction
	fmt.Println(mempool.GetMempoolEntries()[4].GetFirstSeen())

	/* Furthermore you can use the full functionality of MsgTx (https://godoc.org/github.com/btcsuite/btcd/wire#MsgTx):
		  - transaction has a witness (spends a SegWit output)? with <entry>.transaction.HasWitness() 
		  - access input and outputs
		  - ...
		 But there is no way to get feerates for transactions, because they are (rightfully) not stored in the `mempool.dat`.
		 You'd have to the current UTXO set to get input amounts (which you need to calculate the fees)
	*/
}

func readFile(path string) (mempool mempoolDat.Mempool){
	mempool, err := mempoolDat.ReadMempoolFromPath(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return 
}
