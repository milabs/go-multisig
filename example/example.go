package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/hdkeychain"
	goms "github.com/milabs/go-multisig/address"
)

func main() {
	masterPublicKeys := []string{
		// seed: develop begin cushion hurt crisp embody more image employ library help game
		"Vpub5fCbHSqkDPNxCWTxAwJfaHkEPswKSCXbYeEBPchDgecb51ZDrtwASLtWwgUZgwatXJNMPCfCyk5KACj66VRtcH73wViXo8hKwXjC3GsviHW",
		// seed: grief coffee round palace town clerk veteran fever social dawn appear afraid
		"Vpub5fTkDxY46KBtBwWmZPNrJXycYZf2a7TEbHa5Q5zLzqtMvDT2qRM5nAX7J6Y2Ya4iXcJCbrz9u3kxd9moyqatXNHkY5VdJWWB4CxGHS6UMTJ",
		// seed: tobacco develop can sing pudding account forest pond trophy rookie joke few
		"Vpub5fi4GEEsUmMjFnub4jQ3RkVP9mDXnvarch6uprZnW8VFYVbV1VG2QSppTmtQyJMYiaP6NFgdgzyvq3Domj62dQuK94w9ddmkbPxuQsTUsXM",
	}

	masterPKs := []*hdkeychain.ExtendedKey{}
	for i := range (masterPublicKeys) {
		if key, err := hdkeychain.NewKeyFromString(masterPublicKeys[i]); err != nil {
			panic(err)
		} else {
			masterPKs = append(masterPKs, key)
		}
	}

	for i := 0; i < 20; i++ {
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

	for i := 0; i < 10; i++ {
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
