package config

import (
	"testing"

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

func Fuzz_Nosy_DBConfig_DBConfigToBoltBackendConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		db := DefaultDBConfigWithHomePath(homePath)
		db.DBConfigToBoltBackendConfig()
	})
}

func Fuzz_Nosy_DBConfig_GetDBBackend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		db := DefaultDBConfigWithHomePath(homePath)
		db.GetDBBackend()
	})
}

func Fuzz_Nosy_CfgFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		CfgFile(homePath)
	})
}

func Fuzz_Nosy_DataDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		DataDir(homePath)
	})
}

func Fuzz_Nosy_LogDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		LogDir(homePath)
	})
}

func Fuzz_Nosy_LogFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		LogFile(homePath)
	})
}
