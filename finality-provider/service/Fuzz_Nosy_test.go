package service

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/finality-provider/finality-provider/proto"
	"github.com/babylonlabs-io/finality-provider/finality-provider/store"
	"github.com/cometbft/cometbft/crypto/merkle"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"google.golang.org/grpc"
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

func Fuzz_Nosy_ChainPoller_GetBlockInfoChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.GetBlockInfoChan()
	})
}

func Fuzz_Nosy_ChainPoller_IsRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.IsRunning()
	})
}

func Fuzz_Nosy_ChainPoller_NextHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.NextHeight()
	})
}

func Fuzz_Nosy_ChainPoller_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.Start(startHeight)
	})
}

func Fuzz_Nosy_ChainPoller_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.Stop()
	})
}

func Fuzz_Nosy_ChainPoller_blocksWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end uint64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		var limit uint32
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.blocksWithRetry(start, end, limit)
	})
}

func Fuzz_Nosy_ChainPoller_getLatestBlockWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.getLatestBlockWithRetry()
	})
}

func Fuzz_Nosy_ChainPoller_pollChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.pollChain()
	})
}

func Fuzz_Nosy_ChainPoller_setNextHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.setNextHeight(height)
	})
}

func Fuzz_Nosy_ChainPoller_waitForActivation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cp *ChainPoller
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		if cp == nil {
			return
		}

		cp.waitForActivation()
	})
}

func Fuzz_Nosy_CriticalError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ce *CriticalError
		fill_err = tp.Fill(&ce)
		if fill_err != nil {
			return
		}
		if ce == nil {
			return
		}

		ce.Error()
	})
}

func Fuzz_Nosy_FinalityProviderApp_CreateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
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
		var passPhrase string
		fill_err = tp.Fill(&passPhrase)
		if fill_err != nil {
			return
		}
		var eotsPk *types.BIP340PubKey
		fill_err = tp.Fill(&eotsPk)
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
		if app == nil || eotsPk == nil || description == nil || commission == nil {
			return
		}

		app.CreateFinalityProvider(keyName, chainID, passPhrase, eotsPk, description, commission)
	})
}

func Fuzz_Nosy_FinalityProviderApp_CreatePop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fpAddress types.AccAddress
		fill_err = tp.Fill(&fpAddress)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if app == nil || fpPk == nil {
			return
		}

		app.CreatePop(fpAddress, fpPk, passphrase)
	})
}

func Fuzz_Nosy_FinalityProviderApp_GetConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.GetConfig()
	})
}

func Fuzz_Nosy_FinalityProviderApp_GetFinalityProviderInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if app == nil || fpPk == nil {
			return
		}

		app.GetFinalityProviderInfo(fpPk)
	})
}

func Fuzz_Nosy_FinalityProviderApp_GetFinalityProviderInstance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.GetFinalityProviderInstance()
	})
}

func Fuzz_Nosy_FinalityProviderApp_GetFinalityProviderStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.GetFinalityProviderStore()
	})
}

func Fuzz_Nosy_FinalityProviderApp_GetPubRandProofStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.GetPubRandProofStore()
	})
}

func Fuzz_Nosy_FinalityProviderApp_IsFinalityProviderRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if app == nil || fpPk == nil {
			return
		}

		app.IsFinalityProviderRunning(fpPk)
	})
}

func Fuzz_Nosy_FinalityProviderApp_ListAllFinalityProvidersInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.ListAllFinalityProvidersInfo()
	})
}

func Fuzz_Nosy_FinalityProviderApp_Logger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.Logger()
	})
}

func Fuzz_Nosy_FinalityProviderApp_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.Start()
	})
}

func Fuzz_Nosy_FinalityProviderApp_StartFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if app == nil || fpPk == nil {
			return
		}

		app.StartFinalityProvider(fpPk, passphrase)
	})
}

func Fuzz_Nosy_FinalityProviderApp_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.Stop()
	})
}

func Fuzz_Nosy_FinalityProviderApp_SyncAllFinalityProvidersStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.SyncAllFinalityProvidersStatus()
	})
}

func Fuzz_Nosy_FinalityProviderApp_UnjailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fpPk *types.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if app == nil || fpPk == nil {
			return
		}

		app.UnjailFinalityProvider(fpPk)
	})
}

func Fuzz_Nosy_FinalityProviderApp_getFpPrivKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fpPk []byte
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.getFpPrivKey(fpPk)
	})
}

func Fuzz_Nosy_FinalityProviderApp_metricsUpdateLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.metricsUpdateLoop()
	})
}

func Fuzz_Nosy_FinalityProviderApp_monitorCriticalErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.monitorCriticalErr()
	})
}

func Fuzz_Nosy_FinalityProviderApp_putFpFromResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fp *types.FinalityProviderResponse
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var chainID string
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if app == nil || fp == nil {
			return
		}

		app.putFpFromResponse(fp, chainID)
	})
}

func Fuzz_Nosy_FinalityProviderApp_registrationLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.registrationLoop()
	})
}

func Fuzz_Nosy_FinalityProviderApp_removeFinalityProviderInstance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.removeFinalityProviderInstance()
	})
}

func Fuzz_Nosy_FinalityProviderApp_setFinalityProviderSlashed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var fpi *FinalityProviderInstance
		fill_err = tp.Fill(&fpi)
		if fill_err != nil {
			return
		}
		if app == nil || fpi == nil {
			return
		}

		app.setFinalityProviderSlashed(fpi)
	})
}

func Fuzz_Nosy_FinalityProviderApp_startFinalityProviderInstance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		var pk *types.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if app == nil || pk == nil {
			return
		}

		app.startFinalityProviderInstance(pk, passphrase)
	})
}

func Fuzz_Nosy_FinalityProviderApp_unjailFpLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var app *FinalityProviderApp
		fill_err = tp.Fill(&app)
		if fill_err != nil {
			return
		}
		if app == nil {
			return
		}

		app.unjailFpLoop()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_CommitPubRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.CommitPubRand(startHeight)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_DetermineStartHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.DetermineStartHeight()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetBtcPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetBtcPk()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetBtcPkBIP340__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetBtcPkBIP340()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetBtcPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetBtcPkHex()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetChainID()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetConfig()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetFinalityProviderSlashedOrJailedWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetFinalityProviderSlashedOrJailedWithRetry()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetLastCommittedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetLastCommittedHeight()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetLastVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetLastVotedHeight()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetStatus()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetStoreFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetStoreFinalityProvider()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_GetVotingPowerWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.GetVotingPowerWithRetry(height)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_HelperCommitPubRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var tipHeight uint64
		fill_err = tp.Fill(&tipHeight)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.HelperCommitPubRand(tipHeight)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_IsJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.IsJailed()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_IsRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.IsRunning()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_MustSetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var s proto.FinalityProviderStatus
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.MustSetStatus(s)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_MustUpdateStateAfterFinalitySigSubmission__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.MustUpdateStateAfterFinalitySigSubmission(height)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_SetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var s proto.FinalityProviderStatus
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.SetStatus(s)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_ShouldCommitRandomness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.ShouldCommitRandomness()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.Start()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.Stop()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_SubmitBatchFinalitySignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var blocks []*types.BlockInfo
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.SubmitBatchFinalitySignatures(blocks)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_SubmitFinalitySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var b *types.BlockInfo
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp == nil || b == nil {
			return
		}

		fp.SubmitFinalitySignature(b)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_TestCommitPubRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var targetBlockHeight uint64
		fill_err = tp.Fill(&targetBlockHeight)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.TestCommitPubRand(targetBlockHeight)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_TestCommitPubRandWithStartHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var targetBlockHeight uint64
		fill_err = tp.Fill(&targetBlockHeight)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.TestCommitPubRandWithStartHeight(startHeight, targetBlockHeight)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_TestSubmitFinalitySignatureAndExtractPrivKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var b *types.BlockInfo
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var useSafeEOTSFunc bool
		fill_err = tp.Fill(&useSafeEOTSFunc)
		if fill_err != nil {
			return
		}
		if fp == nil || b == nil {
			return
		}

		fp.TestSubmitFinalitySignatureAndExtractPrivKey(b, useSafeEOTSFunc)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_checkBlockFinalization__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.checkBlockFinalization(height)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_commitPubRandPairsWithTiming__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.commitPubRandPairsWithTiming(startHeight)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_finalitySigSubmissionLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.finalitySigSubmissionLoop()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_getBatchBlocksFromChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.getBatchBlocksFromChan()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_getFinalityActivationHeightWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.getFinalityActivationHeightWithRetry()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_getLatestBlockWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.getLatestBlockWithRetry()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_getPubRandList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var numPubRand uint32
		fill_err = tp.Fill(&numPubRand)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.getPubRandList(startHeight, numPubRand)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_highestVotedHeightWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.highestVotedHeightWithRetry()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_lastCommittedPublicRandWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.lastCommittedPublicRandWithRetry(count)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_latestFinalizedHeightWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.latestFinalizedHeightWithRetry()
	})
}

func Fuzz_Nosy_FinalityProviderInstance_processBlocksToVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var blocks []*types.BlockInfo
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.processBlocksToVote(blocks)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_randomnessCommitmentLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.randomnessCommitmentLoop()
	})
}

// skipping Fuzz_Nosy_FinalityProviderInstance_reportCriticalErr__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_FinalityProviderInstance_retrySubmitSigsUntilFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var targetBlocks []*types.BlockInfo
		fill_err = tp.Fill(&targetBlocks)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.retrySubmitSigsUntilFinalized(targetBlocks)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_signFinalitySig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var b *types.BlockInfo
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp == nil || b == nil {
			return
		}

		fp.signFinalitySig(b)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_signPubRandCommit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var numPubRand uint64
		fill_err = tp.Fill(&numPubRand)
		if fill_err != nil {
			return
		}
		var commitment []byte
		fill_err = tp.Fill(&commitment)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.signPubRandCommit(startHeight, numPubRand, commitment)
	})
}

func Fuzz_Nosy_FinalityProviderInstance_updateStateAfterFinalitySigSubmission__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProviderInstance
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.updateStateAfterFinalitySigSubmission(height)
	})
}

func Fuzz_Nosy_Server_RunUntilShutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.RunUntilShutdown(ctx)
	})
}

// skipping Fuzz_Nosy_Server_startGrpcListen__ because parameters include func, chan, or unsupported interface: []net.Listener

func Fuzz_Nosy_fpState_getStoreFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *store.StoredFinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var s *store.FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if fp == nil || s == nil {
			return
		}

		fps := newFpState(fp, s)
		fps.getStoreFinalityProvider()
	})
}

func Fuzz_Nosy_fpState_setLastVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *store.StoredFinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var s *store.FinalityProviderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fp == nil || s == nil {
			return
		}

		fps := newFpState(fp, s)
		fps.setLastVotedHeight(height)
	})
}

func Fuzz_Nosy_fpState_setStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *store.StoredFinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var s2 *store.FinalityProviderStore
		fill_err = tp.Fill(&s2)
		if fill_err != nil {
			return
		}
		var s3 proto.FinalityProviderStatus
		fill_err = tp.Fill(&s3)
		if fill_err != nil {
			return
		}
		if fp == nil || s2 == nil {
			return
		}

		fps := newFpState(fp, s2)
		fps.setStatus(s3)
	})
}

func Fuzz_Nosy_pubRandState_addPubRandProofList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *store.PubRandProofStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
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

		st := newPubRandState(s)
		st.addPubRandProofList(pk, chainID, height, numPubRand, proofList)
	})
}

func Fuzz_Nosy_pubRandState_getPubRandProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *store.PubRandProofStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
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

		st := newPubRandState(s)
		st.getPubRandProof(pk, chainID, height)
	})
}

func Fuzz_Nosy_pubRandState_getPubRandProofList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *store.PubRandProofStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pk []byte
		fill_err = tp.Fill(&pk)
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
		var numPubRand uint64
		fill_err = tp.Fill(&numPubRand)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		st := newPubRandState(s)
		st.getPubRandProofList(pk, chainID, height, numPubRand)
	})
}

func Fuzz_Nosy_rpcServer_AddFinalitySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.AddFinalitySignatureRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if fpa == nil || req == nil {
			return
		}

		r := newRPCServer(fpa)
		r.AddFinalitySignature(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_CreateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.CreateFinalityProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if fpa == nil || req == nil {
			return
		}

		r := newRPCServer(fpa)
		r.CreateFinalityProvider(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_EditFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.EditFinalityProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if fpa == nil || req == nil {
			return
		}

		r := newRPCServer(fpa)
		r.EditFinalityProvider(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_GetInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *proto.GetInfoRequest
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if fpa == nil || _x3 == nil {
			return
		}

		r := newRPCServer(fpa)
		r.GetInfo(_x2, _x3)
	})
}

func Fuzz_Nosy_rpcServer_QueryFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.QueryFinalityProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if fpa == nil || req == nil {
			return
		}

		r := newRPCServer(fpa)
		r.QueryFinalityProvider(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_QueryFinalityProviderList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *proto.QueryFinalityProviderListRequest
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if fpa == nil || _x3 == nil {
			return
		}

		r := newRPCServer(fpa)
		r.QueryFinalityProviderList(_x2, _x3)
	})
}

func Fuzz_Nosy_rpcServer_RegisterWithGrpcServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var grpcServer *grpc.Server
		fill_err = tp.Fill(&grpcServer)
		if fill_err != nil {
			return
		}
		if fpa == nil || grpcServer == nil {
			return
		}

		r := newRPCServer(fpa)
		r.RegisterWithGrpcServer(grpcServer)
	})
}

func Fuzz_Nosy_rpcServer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		if fpa == nil {
			return
		}

		r := newRPCServer(fpa)
		r.Start()
	})
}

func Fuzz_Nosy_rpcServer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		if fpa == nil {
			return
		}

		r := newRPCServer(fpa)
		r.Stop()
	})
}

func Fuzz_Nosy_rpcServer_UnjailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.UnjailFinalityProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if fpa == nil || req == nil {
			return
		}

		r := newRPCServer(fpa)
		r.UnjailFinalityProvider(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_UnsafeRemoveMerkleProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpa *FinalityProviderApp
		fill_err = tp.Fill(&fpa)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.RemoveMerkleProofRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if fpa == nil || req == nil {
			return
		}

		r := newRPCServer(fpa)
		r.UnsafeRemoveMerkleProof(_x2, req)
	})
}

func Fuzz_Nosy_getHashToSignForCommitPubRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, numPubRand uint64, commitment []byte) {
		getHashToSignForCommitPubRand(startHeight, numPubRand, commitment)
	})
}

func Fuzz_Nosy_getMsgToSignForVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockHeight uint64, blockHash []byte) {
		getMsgToSignForVote(blockHeight, blockHash)
	})
}
