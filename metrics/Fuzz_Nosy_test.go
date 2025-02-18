package metrics

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/finality-provider/finality-provider/proto"
	"github.com/babylonlabs-io/finality-provider/finality-provider/store"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.uber.org/zap"
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

func Fuzz_Nosy_EotsMetrics_IncrementEotsFpTotalEotsSignCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		em := NewEotsMetrics()
		em.IncrementEotsFpTotalEotsSignCounter(fpBtcPkHex)
	})
}

func Fuzz_Nosy_EotsMetrics_IncrementEotsFpTotalGeneratedRandomnessCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		em := NewEotsMetrics()
		em.IncrementEotsFpTotalGeneratedRandomnessCounter(fpBtcPkHex)
	})
}

func Fuzz_Nosy_EotsMetrics_IncrementEotsFpTotalSchnorrSignCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		em := NewEotsMetrics()
		em.IncrementEotsFpTotalSchnorrSignCounter(fpBtcPkHex)
	})
}

func Fuzz_Nosy_EotsMetrics_SetEotsFpLastEotsSignHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, height float64) {
		em := NewEotsMetrics()
		em.SetEotsFpLastEotsSignHeight(fpBtcPkHex, height)
	})
}

func Fuzz_Nosy_EotsMetrics_SetEotsFpLastGeneratedRandomnessHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, height float64) {
		em := NewEotsMetrics()
		em.SetEotsFpLastGeneratedRandomnessHeight(fpBtcPkHex, height)
	})
}

func Fuzz_Nosy_FpMetrics_AddToFpTotalCommittedRandomness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, num float64) {
		fm := NewFpMetrics()
		fm.AddToFpTotalCommittedRandomness(fpBtcPkHex, num)
	})
}

func Fuzz_Nosy_FpMetrics_AddToFpTotalVotedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, num float64) {
		fm := NewFpMetrics()
		fm.AddToFpTotalVotedBlocks(fpBtcPkHex, num)
	})
}

func Fuzz_Nosy_FpMetrics_IncrementFpTotalBlocksWithoutVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		fm := NewFpMetrics()
		fm.IncrementFpTotalBlocksWithoutVotingPower(fpBtcPkHex)
	})
}

func Fuzz_Nosy_FpMetrics_IncrementFpTotalFailedRandomness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		fm := NewFpMetrics()
		fm.IncrementFpTotalFailedRandomness(fpBtcPkHex)
	})
}

func Fuzz_Nosy_FpMetrics_IncrementFpTotalFailedVotes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		fm := NewFpMetrics()
		fm.IncrementFpTotalFailedVotes(fpBtcPkHex)
	})
}

func Fuzz_Nosy_FpMetrics_IncrementFpTotalVotedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		fm := NewFpMetrics()
		fm.IncrementFpTotalVotedBlocks(fpBtcPkHex)
	})
}

func Fuzz_Nosy_FpMetrics_RecordBabylonTipHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, height uint64) {
		fm := NewFpMetrics()
		fm.RecordBabylonTipHeight(height)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpLastCommittedRandomnessHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, height uint64) {
		fm := NewFpMetrics()
		fm.RecordFpLastCommittedRandomnessHeight(fpBtcPkHex, height)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpLastProcessedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, height uint64) {
		fm := NewFpMetrics()
		fm.RecordFpLastProcessedHeight(fpBtcPkHex, height)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpLastVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, height uint64) {
		fm := NewFpMetrics()
		fm.RecordFpLastVotedHeight(fpBtcPkHex, height)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpRandomnessTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		fm := NewFpMetrics()
		fm.RecordFpRandomnessTime(fpBtcPkHex)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpSecondsSinceLastRandomness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, seconds float64) {
		fm := NewFpMetrics()
		fm.RecordFpSecondsSinceLastRandomness(fpBtcPkHex, seconds)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpSecondsSinceLastVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string, seconds float64) {
		fm := NewFpMetrics()
		fm.RecordFpSecondsSinceLastVote(fpBtcPkHex, seconds)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBtcPkHex string
		fill_err = tp.Fill(&fpBtcPkHex)
		if fill_err != nil {
			return
		}
		var status proto.FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		fm := NewFpMetrics()
		fm.RecordFpStatus(fpBtcPkHex, status)
	})
}

func Fuzz_Nosy_FpMetrics_RecordFpVoteTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fpBtcPkHex string) {
		fm := NewFpMetrics()
		fm.RecordFpVoteTime(fpBtcPkHex)
	})
}

func Fuzz_Nosy_FpMetrics_RecordLastPolledHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, height uint64) {
		fm := NewFpMetrics()
		fm.RecordLastPolledHeight(height)
	})
}

func Fuzz_Nosy_FpMetrics_RecordPollerStartingHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, height uint64) {
		fm := NewFpMetrics()
		fm.RecordPollerStartingHeight(height)
	})
}

func Fuzz_Nosy_FpMetrics_UpdateFpMetrics__(f *testing.F) {
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
		if fp == nil {
			return
		}

		fm := NewFpMetrics()
		fm.UpdateFpMetrics(fp)
	})
}

func Fuzz_Nosy_Server_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if logger == nil {
			return
		}

		s := Start(addr, logger)
		s.Stop(ctx)
	})
}
