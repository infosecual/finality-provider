package store

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/finality-provider/finality-provider/proto"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cometbft/cometbft/crypto/merkle"
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

func Fuzz_Nosy_FinalityProviderStore_CreateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var fpAddr types.AccAddress
		fill_err = tp.Fill(&fpAddr)
		if fill_err != nil {
			return
		}
		var btcPk *btcec.PublicKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var description *types.Description
		fill_err = tp.Fill(&description)
		if fill_err != nil {
			return
		}
		var commission *math.LegacyDec
		fill_err = tp.Fill(&commission)
		if fill_err != nil {
			return
		}
		var chainID string
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if s == nil || btcPk == nil || description == nil || commission == nil {
			return
		}

		s.CreateFinalityProvider(fpAddr, btcPk, description, commission, chainID)
	})
}

func Fuzz_Nosy_FinalityProviderStore_GetAllStoredFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetAllStoredFinalityProviders()
	})
}

func Fuzz_Nosy_FinalityProviderStore_GetFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var btcPk *btcec.PublicKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		if s == nil || btcPk == nil {
			return
		}

		s.GetFinalityProvider(btcPk)
	})
}

func Fuzz_Nosy_FinalityProviderStore_MustSetFpStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var btcPk *btcec.PublicKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var status proto.FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if s == nil || btcPk == nil {
			return
		}

		s.MustSetFpStatus(btcPk, status)
	})
}

func Fuzz_Nosy_FinalityProviderStore_SetFpDescription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var btcPk *btcec.PublicKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var desc *types.Description
		fill_err = tp.Fill(&desc)
		if fill_err != nil {
			return
		}
		var rate *math.LegacyDec
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if s == nil || btcPk == nil || desc == nil || rate == nil {
			return
		}

		s.SetFpDescription(btcPk, desc, rate)
	})
}

func Fuzz_Nosy_FinalityProviderStore_SetFpLastVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var btcPk *btcec.PublicKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var lastVotedHeight uint64
		fill_err = tp.Fill(&lastVotedHeight)
		if fill_err != nil {
			return
		}
		if s == nil || btcPk == nil {
			return
		}

		s.SetFpLastVotedHeight(btcPk, lastVotedHeight)
	})
}

func Fuzz_Nosy_FinalityProviderStore_SetFpStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var btcPk *btcec.PublicKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var status proto.FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if s == nil || btcPk == nil {
			return
		}

		s.SetFpStatus(btcPk, status)
	})
}

func Fuzz_Nosy_FinalityProviderStore_UpdateFpStatusFromVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var vp uint64
		fill_err = tp.Fill(&vp)
		if fill_err != nil {
			return
		}
		var fp *StoredFinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if s == nil || fp == nil {
			return
		}

		s.UpdateFpStatusFromVotingPower(vp, fp)
	})
}

func Fuzz_Nosy_FinalityProviderStore_createFinalityProviderInternal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var fp *proto.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if s == nil || fp == nil {
			return
		}

		s.createFinalityProviderInternal(fp)
	})
}

func Fuzz_Nosy_FinalityProviderStore_initBuckets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FinalityProviderStore
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

// skipping Fuzz_Nosy_FinalityProviderStore_setFinalityProviderState__ because parameters include func, chan, or unsupported interface: func(provider *github.com/babylonlabs-io/finality-provider/finality-provider/proto.FinalityProvider) error

func Fuzz_Nosy_PubRandProofStore_AddPubRandProofList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PubRandProofStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var numPubRand uint64
		fill_err = tp.Fill(&numPubRand)
		if fill_err != nil {
			return
		}
		var proofList []*merkle.Proof
		fill_err = tp.Fill(&proofList)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddPubRandProofList(chainID, pk, height, numPubRand, proofList)
	})
}

func Fuzz_Nosy_PubRandProofStore_GetPubRandProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PubRandProofStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
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

		s.GetPubRandProof(chainID, pk, height)
	})
}

func Fuzz_Nosy_PubRandProofStore_GetPubRandProofList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PubRandProofStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var numPubRand uint64
		fill_err = tp.Fill(&numPubRand)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetPubRandProofList(chainID, pk, height, numPubRand)
	})
}

func Fuzz_Nosy_PubRandProofStore_RemovePubRandProofList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PubRandProofStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var chainID []byte
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var targetHeight uint64
		fill_err = tp.Fill(&targetHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.RemovePubRandProofList(chainID, pk, targetHeight)
	})
}

func Fuzz_Nosy_PubRandProofStore_initBuckets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PubRandProofStore
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

func Fuzz_Nosy_StoredFinalityProvider_GetBIP340BTCPK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *proto.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		sfp, err := protoFpToStoredFinalityProvider(fp)
		if err != nil {
			return
		}
		sfp.GetBIP340BTCPK()
	})
}

func Fuzz_Nosy_StoredFinalityProvider_ToFinalityProviderInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *proto.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		sfp, err := protoFpToStoredFinalityProvider(fp)
		if err != nil {
			return
		}
		sfp.ToFinalityProviderInfo()
	})
}

func Fuzz_Nosy_buildKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID []byte, pk []byte, height uint64, num uint64) {
		buildKeys(chainID, pk, height, num)
	})
}

func Fuzz_Nosy_getKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID []byte, pk []byte, height uint64) {
		getKey(chainID, pk, height)
	})
}

func Fuzz_Nosy_getPrefixKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID []byte, pk []byte) {
		getPrefixKey(chainID, pk)
	})
}
