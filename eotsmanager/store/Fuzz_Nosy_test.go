package store

import (
	"testing"

	"github.com/babylonlabs-io/finality-provider/eotsmanager/proto"
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

func Fuzz_Nosy_EOTSStore_AddEOTSKeyName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EOTSStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var btcPk *btcec.PublicKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var keyName string
		fill_err = tp.Fill(&keyName)
		if fill_err != nil {
			return
		}
		if s == nil || btcPk == nil {
			return
		}

		s.AddEOTSKeyName(btcPk, keyName)
	})
}

func Fuzz_Nosy_EOTSStore_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EOTSStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Close()
	})
}

func Fuzz_Nosy_EOTSStore_GetAllEOTSKeyNames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EOTSStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetAllEOTSKeyNames()
	})
}

func Fuzz_Nosy_EOTSStore_GetEOTSKeyName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EOTSStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetEOTSKeyName(pk)
	})
}

func Fuzz_Nosy_EOTSStore_GetSignRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EOTSStore
		fill_err = tp.Fill(&s)
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
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetSignRecord(eotsPk, chainID, height)
	})
}

func Fuzz_Nosy_EOTSStore_SaveSignRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EOTSStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
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
		var publicKey []byte
		fill_err = tp.Fill(&publicKey)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SaveSignRecord(height, chainID, msg, publicKey, signature)
	})
}

func Fuzz_Nosy_EOTSStore_initBuckets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EOTSStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.initBuckets()
	})
}

func Fuzz_Nosy_SigningRecord_FromProto__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SigningRecord
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var sr *proto.SigningRecord
		fill_err = tp.Fill(&sr)
		if fill_err != nil {
			return
		}
		if s == nil || sr == nil {
			return
		}

		s.FromProto(sr)
	})
}

func Fuzz_Nosy_getSignRecordKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID []byte, pk []byte, height uint64) {
		getSignRecordKey(chainID, pk, height)
	})
}
