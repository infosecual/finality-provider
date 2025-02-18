package client

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/finality-provider/finality-provider/proto"
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

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_AddFinalitySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpPk string
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var appHash []byte
		fill_err = tp.Fill(&appHash)
		if fill_err != nil {
			return
		}
		var checkDoubleSign bool
		fill_err = tp.Fill(&checkDoubleSign)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.AddFinalitySignature(ctx, fpPk, height, appHash, checkDoubleSign)
	})
}

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_CreateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var keyName string
		fill_err = tp.Fill(&keyName)
		if fill_err != nil {
			return
		}
		var chainID string
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var eotsPkHex string
		fill_err = tp.Fill(&eotsPkHex)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var description types.Description
		fill_err = tp.Fill(&description)
		if fill_err != nil {
			return
		}
		var commission *math.LegacyDec
		fill_err = tp.Fill(&commission)
		if fill_err != nil {
			return
		}
		if c == nil || commission == nil {
			return
		}

		c.CreateFinalityProvider(ctx, keyName, chainID, eotsPkHex, passphrase, description, commission)
	})
}

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_EditFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var desc *proto.Description
		fill_err = tp.Fill(&desc)
		if fill_err != nil {
			return
		}
		var rate string
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if c == nil || fpPk == nil || desc == nil {
			return
		}

		c.EditFinalityProvider(ctx, fpPk, desc, rate)
	})
}

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_GetInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetInfo(ctx)
	})
}

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_QueryFinalityProviderInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if c == nil || fpPk == nil {
			return
		}

		c.QueryFinalityProviderInfo(ctx, fpPk)
	})
}

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_QueryFinalityProviderList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.QueryFinalityProviderList(ctx)
	})
}

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_UnjailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpPk string
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UnjailFinalityProvider(ctx, fpPk)
	})
}

func Fuzz_Nosy_FinalityProviderServiceGRpcClient_UnsafeRemoveMerkleProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *FinalityProviderServiceGRpcClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var chainID string
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var targetHeight uint64
		fill_err = tp.Fill(&targetHeight)
		if fill_err != nil {
			return
		}
		if c == nil || fpPk == nil {
			return
		}

		c.UnsafeRemoveMerkleProof(ctx, fpPk, chainID, targetHeight)
	})
}

func Fuzz_Nosy_NewFinalityProviderServiceGRpcClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteAddr string) {
		NewFinalityProviderServiceGRpcClient(remoteAddr)
	})
}
