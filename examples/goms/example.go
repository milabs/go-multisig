package main

import (
	"fmt"
	"flag"
	"github.com/btcsuite/btcutil/hdkeychain"
	goms "github.com/milabs/go-multisig/address"
)

func main() {
	var n int = -1
	var n_receiving int = 20
	var n_change int = 10

	flag.IntVar(&n, "n", 2, "N of M keys")
	flag.IntVar(&n_receiving, "n_receiving", 20, "Number of receiving addresses to generate")
	flag.IntVar(&n_change, "n_change", 10, "Number of change addresses to generate")
	flag.Parse()

	var masterPublicKeys []string
	for i := range (flag.Args()) {
		masterPublicKeys = append(masterPublicKeys, flag.Args()[i])
	}

	masterPKs := []*hdkeychain.ExtendedKey{}
	for i := range (masterPublicKeys) {
		if key, err := hdkeychain.NewKeyFromString(masterPublicKeys[i]); err != nil {
			panic(err)
		} else {
			masterPKs = append(masterPKs, key)
		}
	}

	for i := 0; i < n_receiving; i++ {
		derivationPath := goms.DerivationPath{ goms.ReceivingAddress, i }
		addr, _, err := goms.GenerateTest3Net(
			goms.Params{
				N: 2,
				MasterPKs: masterPKs,
				DerivationPath: derivationPath,
			},
		)

		if err != nil {
			panic(err)
		}

		fmt.Printf("receiving | %s\n", addr.EncodeAddress())
	}

	for i := 0; i < n_change; i++ {
		derivationPath := goms.DerivationPath{ goms.ChangeAddress, i }
		addr, _, err := goms.GenerateTest3Net(
			goms.Params{
				N: 2,
				MasterPKs: masterPKs,
				DerivationPath: derivationPath,
			},
		)

		if err != nil {
			panic(err)
		}

		fmt.Printf("change | %s\n", addr.EncodeAddress())
	}
}
