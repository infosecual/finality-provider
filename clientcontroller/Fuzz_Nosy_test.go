package clientcontroller

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/finality-provider/finality-provider/config"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/chaincfg"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.etcd.io/etcd/client/v2"
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

func Fuzz_Nosy_BabylonController_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.Close()
	})
}

func Fuzz_Nosy_BabylonController_CommitPubRandList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil || sig == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.CommitPubRandList(fpPk, startHeight, numPubRand, commitment, sig)
	})
}

func Fuzz_Nosy_BabylonController_CreateBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var delBtcPk *types.BIP340PubKey
		fill_err = tp.Fill(&delBtcPk)
		if fill_err != nil {
			return
		}
		var fpPks []*btcec.PublicKey
		fill_err = tp.Fill(&fpPks)
		if fill_err != nil {
			return
		}
		var pop *types.ProofOfPossessionBTC
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var stakingTime uint32
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingValue int64
		fill_err = tp.Fill(&stakingValue)
		if fill_err != nil {
			return
		}
		var stakingTxInfo *types.TransactionInfo
		fill_err = tp.Fill(&stakingTxInfo)
		if fill_err != nil {
			return
		}
		var slashingTx *types.BTCSlashingTx
		fill_err = tp.Fill(&slashingTx)
		if fill_err != nil {
			return
		}
		var delSlashingSig *types.BIP340Signature
		fill_err = tp.Fill(&delSlashingSig)
		if fill_err != nil {
			return
		}
		var unbondingTx []byte
		fill_err = tp.Fill(&unbondingTx)
		if fill_err != nil {
			return
		}
		var unbondingTime uint32
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var unbondingValue int64
		fill_err = tp.Fill(&unbondingValue)
		if fill_err != nil {
			return
		}
		var unbondingSlashingTx *types.BTCSlashingTx
		fill_err = tp.Fill(&unbondingSlashingTx)
		if fill_err != nil {
			return
		}
		var delUnbondingSlashingSig *types.BIP340Signature
		fill_err = tp.Fill(&delUnbondingSlashingSig)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || delBtcPk == nil || pop == nil || stakingTxInfo == nil || slashingTx == nil || delSlashingSig == nil || unbondingSlashingTx == nil || delUnbondingSlashingSig == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.CreateBTCDelegation(delBtcPk, fpPks, pop, stakingTime, stakingValue, stakingTxInfo, slashingTx, delSlashingSig, unbondingTx, unbondingTime, unbondingValue, unbondingSlashingTx, delUnbondingSlashingSig)
	})
}

func Fuzz_Nosy_BabylonController_EditFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var rate *math.LegacyDec
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		var description []byte
		fill_err = tp.Fill(&description)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil || rate == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.EditFinalityProvider(fpPk, rate, description)
	})
}

func Fuzz_Nosy_BabylonController_GetBBNClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.GetBBNClient()
	})
}

func Fuzz_Nosy_BabylonController_GetKeyAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.GetKeyAddress()
	})
}

func Fuzz_Nosy_BabylonController_InsertBtcBlockHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var headers []types.BTCHeaderBytes
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.InsertBtcBlockHeaders(headers)
	})
}

func Fuzz_Nosy_BabylonController_InsertSpvProofs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var submitter string
		fill_err = tp.Fill(&submitter)
		if fill_err != nil {
			return
		}
		var proofs []*types.BTCSpvProof
		fill_err = tp.Fill(&proofs)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.InsertSpvProofs(submitter, proofs)
	})
}

func Fuzz_Nosy_BabylonController_NodeTxIndexEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.NodeTxIndexEnabled()
	})
}

func Fuzz_Nosy_BabylonController_QueryActivatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryActivatedHeight()
	})
}

func Fuzz_Nosy_BabylonController_QueryActiveDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryActiveDelegations(limit)
	})
}

func Fuzz_Nosy_BabylonController_QueryBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryBestBlock()
	})
}

func Fuzz_Nosy_BabylonController_QueryBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryBlock(height)
	})
}

func Fuzz_Nosy_BabylonController_QueryBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryBlocks(startHeight, endHeight, limit)
	})
}

func Fuzz_Nosy_BabylonController_QueryBtcLightClientTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryBtcLightClientTip()
	})
}

func Fuzz_Nosy_BabylonController_QueryCurrentEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryCurrentEpoch()
	})
}

func Fuzz_Nosy_BabylonController_QueryFinalityActivationBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryFinalityActivationBlockHeight()
	})
}

func Fuzz_Nosy_BabylonController_QueryFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryFinalityProvider(fpPk)
	})
}

func Fuzz_Nosy_BabylonController_QueryFinalityProviderHighestVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryFinalityProviderHighestVotedHeight(fpPk)
	})
}

func Fuzz_Nosy_BabylonController_QueryFinalityProviderSlashedOrJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryFinalityProviderSlashedOrJailed(fpPk)
	})
}

func Fuzz_Nosy_BabylonController_QueryFinalityProviderVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryFinalityProviderVotingPower(fpPk, blockHeight)
	})
}

func Fuzz_Nosy_BabylonController_QueryFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryFinalityProviders()
	})
}

func Fuzz_Nosy_BabylonController_QueryLastCommittedPublicRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryLastCommittedPublicRand(fpPk, count)
	})
}

func Fuzz_Nosy_BabylonController_QueryLatestFinalizedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryLatestFinalizedBlocks(count)
	})
}

func Fuzz_Nosy_BabylonController_QueryPendingDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryPendingDelegations(limit)
	})
}

func Fuzz_Nosy_BabylonController_QueryStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryStakingParams()
	})
}

func Fuzz_Nosy_BabylonController_QueryVotesAtHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.QueryVotesAtHeight(height)
	})
}

func Fuzz_Nosy_BabylonController_RegisterFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil || commission == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.RegisterFinalityProvider(fpPk, pop, commission, description)
	})
}

func Fuzz_Nosy_BabylonController_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.Start()
	})
}

func Fuzz_Nosy_BabylonController_SubmitBatchFinalitySigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.SubmitBatchFinalitySigs(fpPk, blocks, pubRandList, proofList, sigs)
	})
}

func Fuzz_Nosy_BabylonController_SubmitCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var covPk *btcec.PublicKey
		fill_err = tp.Fill(&covPk)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var slashingSigs [][]byte
		fill_err = tp.Fill(&slashingSigs)
		if fill_err != nil {
			return
		}
		var unbondingSig *schnorr.Signature
		fill_err = tp.Fill(&unbondingSig)
		if fill_err != nil {
			return
		}
		var unbondingSlashingSigs [][]byte
		fill_err = tp.Fill(&unbondingSlashingSigs)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || covPk == nil || unbondingSig == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.SubmitCovenantSigs(covPk, stakingTxHash, slashingSigs, unbondingSig, unbondingSlashingSigs)
	})
}

func Fuzz_Nosy_BabylonController_SubmitFinalitySig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil || block == nil || pubRand == nil || sig == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.SubmitFinalitySig(fpPk, block, pubRand, proof, sig)
	})
}

func Fuzz_Nosy_BabylonController_UnjailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil || fpPk == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.UnjailFinalityProvider(fpPk)
	})
}

func Fuzz_Nosy_BabylonController_mustGetTxSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.mustGetTxSigner()
	})
}

func Fuzz_Nosy_BabylonController_queryCometBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.queryCometBestBlock()
	})
}

func Fuzz_Nosy_BabylonController_queryDelegationsWithStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var status types.BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.queryDelegationsWithStatus(status, limit)
	})
}

func Fuzz_Nosy_BabylonController_queryLatestBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bbnClient *client.Client
		fill_err = tp.Fill(&bbnClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var btcParams *chaincfg.Params
		fill_err = tp.Fill(&btcParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var startKey []byte
		fill_err = tp.Fill(&startKey)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var status types.QueriedBlockStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var reverse bool
		fill_err = tp.Fill(&reverse)
		if fill_err != nil {
			return
		}
		if bbnClient == nil || cfg == nil || btcParams == nil || logger == nil {
			return
		}

		bc := NewBabylonController(bbnClient, cfg, btcParams, logger)
		bc.queryLatestBlocks(startKey, count, status, reverse)
	})
}

// skipping Fuzz_Nosy_BabylonController_reliablySendMsg__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_BabylonController_reliablySendMsgs__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

func Fuzz_Nosy_ClientController_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.Close()
	})
}

func Fuzz_Nosy_ClientController_CommitPubRandList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil || sig == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.CommitPubRandList(fpPk, startHeight, numPubRand, commitment, sig)
	})
}

func Fuzz_Nosy_ClientController_EditFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil || commission == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.EditFinalityProvider(fpPk, commission, description)
	})
}

func Fuzz_Nosy_ClientController_QueryActivatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryActivatedHeight()
	})
}

func Fuzz_Nosy_ClientController_QueryBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryBestBlock()
	})
}

func Fuzz_Nosy_ClientController_QueryBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryBlock(height)
	})
}

func Fuzz_Nosy_ClientController_QueryBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryBlocks(startHeight, endHeight, limit)
	})
}

func Fuzz_Nosy_ClientController_QueryFinalityActivationBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryFinalityActivationBlockHeight()
	})
}

func Fuzz_Nosy_ClientController_QueryFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryFinalityProvider(fpPk)
	})
}

func Fuzz_Nosy_ClientController_QueryFinalityProviderHighestVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryFinalityProviderHighestVotedHeight(fpPk)
	})
}

func Fuzz_Nosy_ClientController_QueryFinalityProviderSlashedOrJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryFinalityProviderSlashedOrJailed(fpPk)
	})
}

func Fuzz_Nosy_ClientController_QueryFinalityProviderVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryFinalityProviderVotingPower(fpPk, blockHeight)
	})
}

func Fuzz_Nosy_ClientController_QueryLastCommittedPublicRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryLastCommittedPublicRand(fpPk, count)
	})
}

func Fuzz_Nosy_ClientController_QueryLatestFinalizedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.QueryLatestFinalizedBlocks(count)
	})
}

func Fuzz_Nosy_ClientController_RegisterFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil || commission == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.RegisterFinalityProvider(fpPk, pop, commission, description)
	})
}

func Fuzz_Nosy_ClientController_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.Start()
	})
}

func Fuzz_Nosy_ClientController_SubmitBatchFinalitySigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.SubmitBatchFinalitySigs(fpPk, blocks, pubRandList, proofList, sigs)
	})
}

func Fuzz_Nosy_ClientController_SubmitFinalitySig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
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
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil || block == nil || pubRand == nil || sig == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.SubmitFinalitySig(fpPk, block, pubRand, proof, sig)
	})
}

func Fuzz_Nosy_ClientController_UnjailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainType string
		fill_err = tp.Fill(&chainType)
		if fill_err != nil {
			return
		}
		var bbnConfig *config.BBNConfig
		fill_err = tp.Fill(&bbnConfig)
		if fill_err != nil {
			return
		}
		var netParams *chaincfg.Params
		fill_err = tp.Fill(&netParams)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var fpPk *btcec.PublicKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if bbnConfig == nil || netParams == nil || logger == nil || fpPk == nil {
			return
		}

		_x1, err := NewClientController(chainType, bbnConfig, netParams, logger)
		if err != nil {
			return
		}
		_x1.UnjailFinalityProvider(fpPk)
	})
}

func Fuzz_Nosy_ExpectedError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e ExpectedError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

// skipping Fuzz_Nosy_ExpectedError_Is__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_ExpectedError_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e ExpectedError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Unwrap()
	})
}

// skipping Fuzz_Nosy_IsExpected__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_IsUnrecoverable__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_getContextWithCancel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		getContextWithCancel(timeout)
	})
}
