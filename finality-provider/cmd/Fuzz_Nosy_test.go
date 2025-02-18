package cmd

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
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

func Fuzz_Nosy_PersistClientCtx__(f *testing.F) {
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

		PersistClientCtx(ctx)
	})
}

// skipping Fuzz_Nosy_RunEWithClientCtx__ because parameters include func, chan, or unsupported interface: func(ctx github.com/cosmos/cosmos-sdk/client.Context, cmd *github.com/spf13/cobra.Command, args []string) error
