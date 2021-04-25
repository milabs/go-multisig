package address

import (
	"testing"
	"github.com/btcsuite/btcutil/hdkeychain"
)

func TestGenerateTest3Net_receiving(t *testing.T) {
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
			t.Fatal(err)
		} else {
			masterPKs = append(masterPKs, key)
		}
	}

	receivingOK := []string{
		"tb1q6p85lf6za5tg63yzsgwsderzxqs44lyq09vr4cjggwws755nu2jqk8nwn5",
		"tb1qa4eg0ztkejmlmla25x67fp37v0a739yteplhk69q6m8qvthv6y2sh3au6a",
		"tb1q2nymy2727lz28v7je2nc69amk9r8trm2r3d9n4sx3g9rpfskvfyqt9hfs4",
		"tb1qts0slstmlqdy9ldv4s33y3gyzapaq2lxl3e4xd60acxefwldtamsjqzszm",
		"tb1qkw5r60k2aef0x2we9dwan9jt6ayk7f8qfl30wrvsjf6hj0h0x3tqntjjcy",
		"tb1q402r9y755c08x4w3uft39xp4uzdav6uxu9g5qx45y8jskf5lpjjsfxu33h",
		"tb1qwueh72zveqpdvgc4fl5khnfeqntsvh7msy0dwam359rvycnx2p6qs7wmak",
		"tb1q8kh7q6vxgq5mqfakfgfcu2vuvhww3kvr8luyktw5ua0g47e3kzhs9cxne2",
		"tb1qqef9fxxl4qd9770czntpwx5w6ydr4vrjt4vygsuzj3ewyfen9luq8z4rj8",
		"tb1ql7ny9ek05099e7wza954vr847j6wuw94sydg0mvlzk8uytxx4sfsld64hj",
		"tb1qdhcplyjw80axqzc8dvnf4nkvsmjc2py3zv9qvhnfyk9fl2hhm69sk92skg",
		"tb1qj6826j88tc23e0dcz8myl7ults2lfk8s4rxv3u6crw6rv73lk24q9qfuyq",
		"tb1qfml2e009ec6r3vy6e0v9vh74pw2agxvxz5rlud2gfayt4v0zl03qxeaj5m",
		"tb1q6fp72vyr0x8xl53r30qcqzgh2tlv59qnyphgg87c35y4khthq6gs3x8r94",
		"tb1q4cpegsktntvp7pll97aktrgnhf9x0dcevrn305fpz9cvdyneqf8sh8unf6",
		"tb1qxlhmce7x9kne4um75vp85ch5f4z5ncfvffwaefwu9k6pfmw4ltqs5suhau",
		"tb1q6w3a07aumcjyhrnnmk3n2970k9ef2exkxm0lj76l3xh80wvtpluq88z6n0",
		"tb1qy797z6pjjdzm7kv587dqhppazn3s8dqgy8a52jpzqn0pf33shahsxqqvh9",
		"tb1qsdmdlgz5xku8nefqapmruzyc0mgwfc9tsh2ja83usepl8l2hgl3qmucfan",
		"tb1qz7jny90ly9wquzx0f7y0m8uj8g3q6dfn67u53wwer5js5uc8jsmsqenklh",
	}

	for i := 0; i < len(receivingOK); i++ {
		derivationPath := DerivationPath{ ReceivingAddress, i }
		addr, _, err := GenerateTest3Net(
			Params{
				N: 2,
				MasterPKs: masterPKs,
				DerivationPath: derivationPath,
			},
		)

		if err != nil {
			t.Fatal(err)
		}

		if receivingOK[i] != addr.EncodeAddress() {
			t.Fatalf("Got %s (expected %s)\n", addr.EncodeAddress(), receivingOK[i])
		}
	}
}

func TestGenerateTest3Net_change(t *testing.T) {
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
			t.Fatal(err)
		} else {
			masterPKs = append(masterPKs, key)
		}
	}

	changeOK := []string{
		"tb1qnkdxgs6j28m9g28f0mn3r0ew4ld9k8kh3w7ax2gfmkp29w6pmtwqmcfazv",
		"tb1qspuz2mnn0msx3drr34thw6y5cweel0yfvzksljq4mdsx0dvt603qn2329s",
		"tb1qj98qg4v7406u7vjl4fme6dwygaf3a95fcmj23tqlkq892lzexjfsatxpeh",
		"tb1qwznae5f5pkq2la8xr3hvkyf32qyug6awa6tynnz67mjw6vl762tszcs327",
		"tb1qw8cl66e3nvv2ujzhwfjefdp4pewcc8vz7a6dnfa645wdrtyk0m2qndztqy",
		"tb1qwmuqs5ju9nxxkfc39n6aqhucfavhh7fv85dwvftsydxlq2qppm3skuqxpc",
		"tb1q5tvvsuv2fhjyrh5cwz08x4fgmhls2g45wrauva9d0a9zeea8xpwqh8h5k7",
		"tb1qwue4lg5q86fj9v6rj5w5cpd3ww0awpsdgkcfmmj8j6cpxhasxpuq5jsrqg",
		"tb1qfgv6jamd747lpmy4l7yf259swqg02s5v4d0dlz4k583hschww9qsyr9z9j",
		"tb1qm8fm2j6zkkjd0870wf6pa0raujp8hl80ykjscmj6f5pjrh4t6l0sgysh2r",
	}

	for i := 0; i < len(changeOK); i++ {
		derivationPath := DerivationPath{ ChangeAddress, i }
		addr, _, err := GenerateTest3Net(
			Params{
				N: 2,
				MasterPKs: masterPKs,
				DerivationPath: derivationPath,
			},
		)

		if err != nil {
			t.Fatal(err)
		}

		if changeOK[i] != addr.EncodeAddress() {
			t.Fatalf("Got %s (expected %s)\n", addr.EncodeAddress(), changeOK[i])
		}
	}
}
