package daemon

import (
	"io"
	"testing"

	"github.com/babylonlabs-io/finality-provider/eotsmanager"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
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

// skipping Fuzz_Nosy_SignCosmosAdr36__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/crypto/keyring.Keyring

func Fuzz_Nosy_ValidBabySignEots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, babyPk string, babyAddr string, eotsPkHex string, babySigOverEotsPk string) {
		ValidBabySignEots(babyPk, babyAddr, eotsPkHex, babySigOverEotsPk)
	})
}

func Fuzz_Nosy_ValidEotsSignBaby__(f *testing.F) {
	f.Fuzz(func(t *testing.T, eotsPk string, babyAddr string, eotsSigOverBabyAddr string) {
		ValidEotsSignBaby(eotsPk, babyAddr, eotsSigOverBabyAddr)
	})
}

func Fuzz_Nosy_ValidPopExport__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pop PoPExport
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}

		ValidPopExport(pop)
	})
}

func Fuzz_Nosy_babyFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		babyFlags(cmd)
	})
}

func Fuzz_Nosy_babyKeyring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babyHomePath string
		fill_err = tp.Fill(&babyHomePath)
		if fill_err != nil {
			return
		}
		var babyKeyName string
		fill_err = tp.Fill(&babyKeyName)
		if fill_err != nil {
			return
		}
		var babyKeyringBackend string
		fill_err = tp.Fill(&babyKeyringBackend)
		if fill_err != nil {
			return
		}
		var userInput io.Reader
		fill_err = tp.Fill(&userInput)
		if fill_err != nil {
			return
		}

		babyKeyring(babyHomePath, babyKeyName, babyKeyringBackend, userInput)
	})
}

func Fuzz_Nosy_cmdCloseEots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		var eotsManager *eotsmanager.LocalEOTSManager
		fill_err = tp.Fill(&eotsManager)
		if fill_err != nil {
			return
		}
		if cmd == nil || eotsManager == nil {
			return
		}

		cmdCloseEots(cmd, eotsManager)
	})
}

func Fuzz_Nosy_eotsFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		eotsFlags(cmd)
	})
}

func Fuzz_Nosy_eotsSignMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eotsManager *eotsmanager.LocalEOTSManager
		fill_err = tp.Fill(&eotsManager)
		if fill_err != nil {
			return
		}
		var keyName string
		fill_err = tp.Fill(&keyName)
		if fill_err != nil {
			return
		}
		var fpPkStr string
		fill_err = tp.Fill(&fpPkStr)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var hashOfMsgToSign []byte
		fill_err = tp.Fill(&hashOfMsgToSign)
		if fill_err != nil {
			return
		}
		if eotsManager == nil {
			return
		}

		eotsSignMsg(eotsManager, keyName, fpPkStr, passphrase, hashOfMsgToSign)
	})
}

func Fuzz_Nosy_getAllEOTSKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		getAllEOTSKeys(cmd)
	})
}

func Fuzz_Nosy_getCleanPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		var flag string
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		getCleanPath(cmd, flag)
	})
}

func Fuzz_Nosy_getHomePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		getHomePath(cmd)
	})
}

func Fuzz_Nosy_getInterpretedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		getInterpretedMessage(cmd)
	})
}

func Fuzz_Nosy_readMnemonicFromFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		readMnemonicFromFile(filePath)
	})
}

func Fuzz_Nosy_saveKeyOnPostRun__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		var commandName string
		fill_err = tp.Fill(&commandName)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		saveKeyOnPostRun(cmd, commandName)
	})
}
