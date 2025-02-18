package keyring

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_ChainKeyringController_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx client.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var keyringBackend string
		fill_err = tp.Fill(&keyringBackend)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}

		kc, err := NewChainKeyringController(ctx, name, keyringBackend)
		if err != nil {
			return
		}
		kc.Address(passphrase)
	})
}

func Fuzz_Nosy_ChainKeyringController_CreateChainKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx client.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var keyringBackend string
		fill_err = tp.Fill(&keyringBackend)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var hdPath string
		fill_err = tp.Fill(&hdPath)
		if fill_err != nil {
			return
		}
		var mnemonic string
		fill_err = tp.Fill(&mnemonic)
		if fill_err != nil {
			return
		}

		kc, err := NewChainKeyringController(ctx, name, keyringBackend)
		if err != nil {
			return
		}
		kc.CreateChainKey(passphrase, hdPath, mnemonic)
	})
}

func Fuzz_Nosy_ChainKeyringController_CreatePop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx client.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var keyringBackend string
		fill_err = tp.Fill(&keyringBackend)
		if fill_err != nil {
			return
		}
		var fpAddr types.AccAddress
		fill_err = tp.Fill(&fpAddr)
		if fill_err != nil {
			return
		}
		var btcPrivKey *btcec.PrivateKey
		fill_err = tp.Fill(&btcPrivKey)
		if fill_err != nil {
			return
		}
		if btcPrivKey == nil {
			return
		}

		kc, err := NewChainKeyringController(ctx, name, keyringBackend)
		if err != nil {
			return
		}
		kc.CreatePop(fpAddr, btcPrivKey)
	})
}

func Fuzz_Nosy_ChainKeyringController_GetChainPrivKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx client.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var keyringBackend string
		fill_err = tp.Fill(&keyringBackend)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}

		kc, err := NewChainKeyringController(ctx, name, keyringBackend)
		if err != nil {
			return
		}
		kc.GetChainPrivKey(passphrase)
	})
}

func Fuzz_Nosy_ChainKeyringController_GetKeyring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx client.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var keyringBackend string
		fill_err = tp.Fill(&keyringBackend)
		if fill_err != nil {
			return
		}

		kc, err := NewChainKeyringController(ctx, name, keyringBackend)
		if err != nil {
			return
		}
		kc.GetKeyring()
	})
}
