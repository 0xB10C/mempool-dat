# mempool-dat

This Go package parses bitcoin mempool `mempool.dat` files.
These are wirtten by Bitcoin Core v0.14+ on shutdown and since v0.15.0 with the RCP `savemempool`.

The package offers access to the `mempool.dat`
- header: version and number of transactions
- mempool entries: raw transaction (here parsed as [btcsuite/btcd/wire MsgTx](https://godoc.org/github.com/btcsuite/btcd/wire#MsgTx) ), first seen timestamp and the feeDelta  
- and the not-parsed mapDeltas as byte slices 

You may see this package as Work-In-Progress. There are no tests yet. 

## Example usage

For a runnable version of this snippet see [main.go](./main.go).


```go
import (
  "fmt"
  mempoolDat "github.com/0xb10c/mempool-dat/lib"
)

[...]

const path string = "mempool.dat"

[...]

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

```

The full documentation can be fund on https://godoc.org/github.com/0xb10c/mempool-dat.

