package eotsmanager

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
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

func Fuzz_Nosy_LocalEOTSManager_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.Close()
	})
}

func Fuzz_Nosy_LocalEOTSManager_CreateKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
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
		if lm == nil {
			return
		}

		lm.CreateKey(name, passphrase, hdPath)
	})
}

func Fuzz_Nosy_LocalEOTSManager_CreateKeyWithMnemonic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
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
		if lm == nil {
			return
		}

		lm.CreateKeyWithMnemonic(name, passphrase, hdPath, mnemonic)
	})
}

func Fuzz_Nosy_LocalEOTSManager_CreateRandomnessPairList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var num uint32
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.CreateRandomnessPairList(fpPk, chainID, startHeight, num, passphrase)
	})
}

func Fuzz_Nosy_LocalEOTSManager_KeyRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.KeyRecord(fpPk, passphrase)
	})
}

func Fuzz_Nosy_LocalEOTSManager_ListEOTSKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.ListEOTSKeys()
	})
}

func Fuzz_Nosy_LocalEOTSManager_LoadBIP340PubKeyFromKeyName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var keyName string
		fill_err = tp.Fill(&keyName)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.LoadBIP340PubKeyFromKeyName(keyName)
	})
}

func Fuzz_Nosy_LocalEOTSManager_SaveEOTSKeyName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var pk *btcec.PublicKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var keyName string
		fill_err = tp.Fill(&keyName)
		if fill_err != nil {
			return
		}
		if lm == nil || pk == nil {
			return
		}

		lm.SaveEOTSKeyName(pk, keyName)
	})
}

func Fuzz_Nosy_LocalEOTSManager_SignEOTS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var eotsPk []byte
		fill_err = tp.Fill(&eotsPk)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.SignEOTS(eotsPk, chainID, msg, height, passphrase)
	})
}

func Fuzz_Nosy_LocalEOTSManager_SignSchnorrSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.SignSchnorrSig(fpPk, msg, passphrase)
	})
}

func Fuzz_Nosy_LocalEOTSManager_SignSchnorrSigFromKeyname__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var keyName string
		fill_err = tp.Fill(&keyName)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.SignSchnorrSigFromKeyname(keyName, passphrase, msg)
	})
}

func Fuzz_Nosy_LocalEOTSManager_UnsafeSignEOTS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.UnsafeSignEOTS(fpPk, chainID, msg, height, passphrase)
	})
}

func Fuzz_Nosy_LocalEOTSManager_eotsPrivKeyFromKeyName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var keyName string
		fill_err = tp.Fill(&keyName)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.eotsPrivKeyFromKeyName(keyName)
	})
}

func Fuzz_Nosy_LocalEOTSManager_getEOTSPrivKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.getEOTSPrivKey(fpPk, passphrase)
	})
}

func Fuzz_Nosy_LocalEOTSManager_getRandomnessPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.getRandomnessPair(fpPk, chainID, height, passphrase)
	})
}

func Fuzz_Nosy_LocalEOTSManager_keyExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if lm == nil {
			return
		}

		lm.keyExists(name)
	})
}

func Fuzz_Nosy_LocalEOTSManager_signSchnorrSigFromPrivKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lm *LocalEOTSManager
		fill_err = tp.Fill(&lm)
		if fill_err != nil {
			return
		}
		var privKey *btcec.PrivateKey
		fill_err = tp.Fill(&privKey)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if lm == nil || privKey == nil {
			return
		}

		lm.signSchnorrSigFromPrivKey(privKey, fpPk, msg)
	})
}

// skipping Fuzz_Nosy_EOTSManager_Close__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager

// skipping Fuzz_Nosy_EOTSManager_CreateKey__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager

// skipping Fuzz_Nosy_EOTSManager_CreateRandomnessPairList__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager

// skipping Fuzz_Nosy_EOTSManager_KeyRecord__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager

// skipping Fuzz_Nosy_EOTSManager_SaveEOTSKeyName__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager

// skipping Fuzz_Nosy_EOTSManager_SignEOTS__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager

// skipping Fuzz_Nosy_EOTSManager_SignSchnorrSig__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager

// skipping Fuzz_Nosy_EOTSManager_UnsafeSignEOTS__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/finality-provider/eotsmanager.EOTSManager
