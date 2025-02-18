package mocks

import (
	"testing"

	math "cosmossdk.io/math"
	"github.com/babylonlabs-io/finality-provider/types"
	btcec "github.com/btcsuite/btcd/btcec/v2"
	schnorr "github.com/btcsuite/btcd/btcec/v2/schnorr"
	gomock "github.com/golang/mock/gomock"
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

func Fuzz_Nosy_MockClientController_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.Close()
	})
}

func Fuzz_Nosy_MockClientController_CommitPubRandList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
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
		var sig *schnorr.Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil || sig == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.CommitPubRandList(fpPk, startHeight, numPubRand, commitment, sig)
	})
}

func Fuzz_Nosy_MockClientController_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockClientController_EditFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var commission *math.LegacyDec
		fill_err = tp.Fill(&commission)
		if fill_err != nil {
			return
		}
		var description []byte
		fill_err = tp.Fill(&description)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil || commission == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.EditFinalityProvider(fpPk, commission, description)
	})
}

func Fuzz_Nosy_MockClientController_QueryActivatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryActivatedHeight()
	})
}

func Fuzz_Nosy_MockClientController_QueryBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryBestBlock()
	})
}

func Fuzz_Nosy_MockClientController_QueryBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryBlock(height)
	})
}

func Fuzz_Nosy_MockClientController_QueryBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint64
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var limit uint32
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryBlocks(startHeight, endHeight, limit)
	})
}

func Fuzz_Nosy_MockClientController_QueryFinalityActivationBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryFinalityActivationBlockHeight()
	})
}

func Fuzz_Nosy_MockClientController_QueryFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryFinalityProvider(fpPk)
	})
}

func Fuzz_Nosy_MockClientController_QueryFinalityProviderHighestVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryFinalityProviderHighestVotedHeight(fpPk)
	})
}

func Fuzz_Nosy_MockClientController_QueryFinalityProviderSlashedOrJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryFinalityProviderSlashedOrJailed(fpPk)
	})
}

func Fuzz_Nosy_MockClientController_QueryFinalityProviderVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var blockHeight uint64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryFinalityProviderVotingPower(fpPk, blockHeight)
	})
}

func Fuzz_Nosy_MockClientController_QueryLastCommittedPublicRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryLastCommittedPublicRand(fpPk, count)
	})
}

func Fuzz_Nosy_MockClientController_QueryLatestFinalizedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.QueryLatestFinalizedBlocks(count)
	})
}

func Fuzz_Nosy_MockClientController_RegisterFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var pop []byte
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var commission *math.LegacyDec
		fill_err = tp.Fill(&commission)
		if fill_err != nil {
			return
		}
		var description []byte
		fill_err = tp.Fill(&description)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil || commission == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.RegisterFinalityProvider(fpPk, pop, commission, description)
	})
}

func Fuzz_Nosy_MockClientController_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.Start()
	})
}

func Fuzz_Nosy_MockClientController_SubmitBatchFinalitySigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var blocks []*types.BlockInfo
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		var pubRandList []*btcec.FieldVal
		fill_err = tp.Fill(&pubRandList)
		if fill_err != nil {
			return
		}
		var proofList [][]byte
		fill_err = tp.Fill(&proofList)
		if fill_err != nil {
			return
		}
		var sigs []*btcec.ModNScalar
		fill_err = tp.Fill(&sigs)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.SubmitBatchFinalitySigs(fpPk, blocks, pubRandList, proofList, sigs)
	})
}

func Fuzz_Nosy_MockClientController_SubmitFinalitySig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var block *types.BlockInfo
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var pubRand *btcec.FieldVal
		fill_err = tp.Fill(&pubRand)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var sig *btcec.ModNScalar
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil || block == nil || pubRand == nil || sig == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.SubmitFinalitySig(fpPk, block, pubRand, proof, sig)
	})
}

func Fuzz_Nosy_MockClientController_UnjailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpPk == nil {
			return
		}

		m := NewMockClientController(ctrl)
		m.UnjailFinalityProvider(fpPk)
	})
}

func Fuzz_Nosy_MockClientControllerMockRecorder_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockClientControllerMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.Close()
	})
}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_CommitPubRandList__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_EditFinalityProvider__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockClientControllerMockRecorder_QueryActivatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockClientControllerMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.QueryActivatedHeight()
	})
}

func Fuzz_Nosy_MockClientControllerMockRecorder_QueryBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockClientControllerMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.QueryBestBlock()
	})
}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryBlock__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryBlocks__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockClientControllerMockRecorder_QueryFinalityActivationBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockClientControllerMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.QueryFinalityActivationBlockHeight()
	})
}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryFinalityProvider__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryFinalityProviderHighestVotedHeight__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryFinalityProviderSlashedOrJailed__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryFinalityProviderVotingPower__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryLastCommittedPublicRand__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_QueryLatestFinalizedBlocks__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_RegisterFinalityProvider__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockClientControllerMockRecorder_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockClientControllerMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.Start()
	})
}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_SubmitBatchFinalitySigs__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_SubmitFinalitySig__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockClientControllerMockRecorder_UnjailFinalityProvider__ because parameters include func, chan, or unsupported interface: interface{}
