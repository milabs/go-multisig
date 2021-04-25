package address

import (
	"fmt"
	"sort"
	"strings"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"crypto/sha256"
)

const (
	ReceivingAddress = 0 // BIP32 m/0/<n>
	ChangeAddress = 1 // BIP32 m/1/<n>
)

type DerivationPath []int

func (d DerivationPath) String() string {
	var items []string
	for _, n := range (d) {
		items = append(items, fmt.Sprintf("%d", n))
	}
	return fmt.Sprintf("m/%s", strings.Join(items, "/"))
}

type Params struct {
	N int // N of MasterPKs
	MasterPKs []*hdkeychain.ExtendedKey
	DerivationPath []int
	Network *chaincfg.Params
}

func (p Params) Validate() error {
	if p.Network == nil {
		return fmt.Errorf("Invalid network parameter")
	}

	if len(p.DerivationPath) == 0 {
		return fmt.Errorf("Invalid derivation path length")
	}

	if len(p.MasterPKs) < 2 || p.N > len(p.MasterPKs) {
		return fmt.Errorf("Invalid N of []MasterPKs parameter set")
	}

	return nil
}

func Generate(p Params) (btcutil.Address, []byte, error) {
	if err := p.Validate(); err != nil {
		return nil, []byte{}, err
	}

	// derive public keys from MasterPKs using derivation path
	
	publicKeys := make([]*btcec.PublicKey, len(p.MasterPKs))
	for i := range (p.MasterPKs) {
		var key *hdkeychain.ExtendedKey = p.MasterPKs[i]
		for j := 0; j < len(p.DerivationPath); j++ {
			if child, err := key.Child(uint32(p.DerivationPath[j])); err != nil {
				return nil, []byte{}, err
			} else {
				key = child
			}
		}
		if pub, err := key.ECPubKey(); err != nil {
			return nil, []byte{}, err
		} else {
			publicKeys[i] = pub
		}
	}

	// sort keys lexicographically

	sort.Slice(publicKeys, func(i, j int) bool {
		return hex.EncodeToString(publicKeys[i].SerializeCompressed()) < hex.EncodeToString(publicKeys[j].SerializeCompressed())
	})

	// gather together compressed public keys
	
	compressedPublicKeys := make([]*btcutil.AddressPubKey, len(publicKeys))
	for i := range(publicKeys) {
		if pub, err := btcutil.NewAddressPubKey(publicKeys[i].SerializeCompressed(), p.Network); err != nil {
			return nil, []byte{}, err
		} else {
			compressedPublicKeys[i] = pub
		}
	}

	// generate N-M multisig redeem (witness) script

	script, err := txscript.MultiSigScript(compressedPublicKeys, p.N)
	if err != nil {
		return nil, []byte{}, err
	}

	// generate witness script address
	
	sum := sha256.Sum256(script)
	addr, err := btcutil.NewAddressWitnessScriptHash(sum[:], p.Network)
	if err != nil {
		return nil, []byte{}, err
	}

	return addr, script, nil
}

func GenerateMainNet(p Params) (btcutil.Address, []byte, error) {
	p.Network = &chaincfg.MainNetParams
	return Generate(p)
}

func GenerateTest3Net(p Params) (btcutil.Address, []byte, error) {
	p.Network = &chaincfg.TestNet3Params
	return Generate(p)
}
