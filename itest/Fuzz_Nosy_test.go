package e2etest

import (
	"context"
	"math/rand"
	"testing"

	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/babylonlabs-io/finality-provider/eotsmanager/config"
	"github.com/babylonlabs-io/finality-provider/finality-provider/service"
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

func Fuzz_Nosy_EOTSServerHandler_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var eotsHomeDir string
		fill_err = tp.Fill(&eotsHomeDir)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil || cfg == nil {
			return
		}

		eh := NewEOTSServerHandler(t1, cfg, eotsHomeDir)
		eh.Start(ctx)
	})
}

func Fuzz_Nosy_EOTSServerHandler_startServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var eotsHomeDir string
		fill_err = tp.Fill(&eotsHomeDir)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil || cfg == nil {
			return
		}

		eh := NewEOTSServerHandler(t1, cfg, eotsHomeDir)
		eh.startServer(ctx)
	})
}

func Fuzz_Nosy_TestManager_AddFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c2 context.Context
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, c2)
		tm.AddFinalityProvider(t3, c4)
	})
}

func Fuzz_Nosy_TestManager_CheckBlockFinalization__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var num int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.CheckBlockFinalization(t3, height, num)
	})
}

func Fuzz_Nosy_TestManager_FinalizeUntilEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.FinalizeUntilEpoch(t3, epoch)
	})
}

func Fuzz_Nosy_TestManager_GetFpPrivKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.GetFpPrivKey(t3, fpPk)
	})
}

func Fuzz_Nosy_TestManager_InsertBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var fpPks []*btcec.PublicKey
		fill_err = tp.Fill(&fpPks)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount int64
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.InsertBTCDelegation(t3, fpPks, stakingTime, stakingAmount)
	})
}

func Fuzz_Nosy_TestManager_InsertCovenantSigForDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var btcDel *types.BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil || btcDel == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.InsertCovenantSigForDelegation(t3, btcDel)
	})
}

func Fuzz_Nosy_TestManager_InsertWBTCHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil || r == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.InsertWBTCHeaders(t3, r)
	})
}

func Fuzz_Nosy_TestManager_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.Stop(t3)
	})
}

func Fuzz_Nosy_TestManager_StopAndRestartFpAfterNBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var fpIns *service.FinalityProviderInstance
		fill_err = tp.Fill(&fpIns)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil || fpIns == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.StopAndRestartFpAfterNBlocks(t3, n, fpIns)
	})
}

func Fuzz_Nosy_TestManager_WaitForFpPubRandTimestamped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var fpIns *service.FinalityProviderInstance
		fill_err = tp.Fill(&fpIns)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil || fpIns == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.WaitForFpPubRandTimestamped(t3, fpIns)
	})
}

func Fuzz_Nosy_TestManager_WaitForFpVoteCast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var fpIns *service.FinalityProviderInstance
		fill_err = tp.Fill(&fpIns)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil || fpIns == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.WaitForFpVoteCast(t3, fpIns)
	})
}

func Fuzz_Nosy_TestManager_WaitForNActiveDels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.WaitForNActiveDels(t3, n)
	})
}

func Fuzz_Nosy_TestManager_WaitForNFinalizedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.WaitForNFinalizedBlocks(t3, n)
	})
}

func Fuzz_Nosy_TestManager_WaitForNPendingDels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.WaitForNPendingDels(t3, n)
	})
}

func Fuzz_Nosy_TestManager_WaitForServicesStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		tm := StartManager(t1, ctx)
		tm.WaitForServicesStart(t3)
	})
}

func Fuzz_Nosy_StartManagerWithFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		StartManagerWithFinalityProvider(t1, n, ctx)
	})
}

func Fuzz_Nosy_generateCovenantCommittee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var numCovenants int
		fill_err = tp.Fill(&numCovenants)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if t2 == nil {
			return
		}

		generateCovenantCommittee(numCovenants, t2)
	})
}

func Fuzz_Nosy_tempDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var pattern string
		fill_err = tp.Fill(&pattern)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		tempDir(t1, pattern)
	})
}
