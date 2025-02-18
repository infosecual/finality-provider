package client

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

func Fuzz_Nosy_EOTSManagerGRpcClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.Close()
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_CreateKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string, name string, passphrase string, hdPath string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.CreateKey(name, passphrase, hdPath)
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_CreateRandomnessPairList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string, uid []byte, chainID []byte, startHeight uint64, num uint32, passphrase string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.CreateRandomnessPairList(uid, chainID, startHeight, num, passphrase)
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_KeyRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string, uid []byte, passphrase string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.KeyRecord(uid, passphrase)
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.Ping()
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_SaveEOTSKeyName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var remoteAddr string
		fill_err = tp.Fill(&remoteAddr)
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
		if pk == nil {
			return
		}

		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.SaveEOTSKeyName(pk, keyName)
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_SignEOTS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string, uid []byte, chaiID []byte, msg []byte, height uint64, passphrase string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.SignEOTS(uid, chaiID, msg, height, passphrase)
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_SignSchnorrSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string, uid []byte, msg []byte, passphrase string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.SignSchnorrSig(uid, msg, passphrase)
	})
}

func Fuzz_Nosy_EOTSManagerGRpcClient_UnsafeSignEOTS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string, uid []byte, chaiID []byte, msg []byte, height uint64, passphrase string) {
		c, err := NewEOTSManagerGRpcClient(remoteAddr)
		if err != nil {
			return
		}
		c.UnsafeSignEOTS(uid, chaiID, msg, height, passphrase)
	})
}
