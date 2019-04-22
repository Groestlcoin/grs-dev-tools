package main

import (
	"fmt"
	"os"

	"github.com/Groestlcoin/grsutil/base58" // use blockbook branch
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Converts address/xpub/WIF between Bitcoin and Groestlcoin formats\n\nusage: %s <base58-string>\n", os.Args[0])
		return
	}
	in := base58.Decode(os.Args[1])
	if len(in) < 4 {
		fmt.Printf("Input is too short\n")
		return
	}
	if in[0] == 36 {
		in[0] = 0 // Convert to Bitcoin P2PKH address
	}
	fmt.Printf("BTC: %s\n", base58.CheckEncode(in[:len(in)-4], []byte{}, base58.Sha256D))
	if in[0] == 0 {
		in[0] = 36 // Convert to GRS P2PKH address
	}
	fmt.Printf("GRS: %s\n", base58.CheckEncode(in[:len(in)-4], []byte{}, base58.Groestl512D))
}
