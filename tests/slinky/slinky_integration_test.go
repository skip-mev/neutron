package slinky_test

import (
	sdkmath "cosmossdk.io/math"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/gov"
	feemarketmodule "github.com/skip-mev/feemarket/x/feemarket"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/skip-mev/slinky/tests/integration"
	marketmapmodule "github.com/skip-mev/slinky/x/marketmap"
	"github.com/skip-mev/slinky/x/oracle"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("neutron", "neutronpub")
	cfg.Seal()
}

var (
	minGasPrice = sdkmath.LegacyNewDec(10)
	gasPrice    = sdkmath.LegacyNewDec(1000000)

	image = ibc.DockerImage{
		Repository: "neutron-node",
		Version:    "latest",
		UidGid:     "1025:1025",
	}

	numValidators = 4
	numFullNodes  = 0
	noHostMount   = false
	gasAdjustment = 1.5

	oracleImage = ibc.DockerImage{
		Repository: "ghcr.io/skip-mev/slinky-sidecar",
		Version:    "latest",
		UidGid:     "1000:1000",
	}
	encodingConfig = testutil.MakeTestEncodingConfig(
		bank.AppModuleBasic{},
		oracle.AppModuleBasic{},
		gov.AppModuleBasic{},
		auth.AppModuleBasic{},
		marketmapmodule.AppModuleBasic{},
		feemarketmodule.AppModuleBasic{},
	)

	defaultGenesisKV = []cosmos.GenesisKV{
		{
			Key:   "consensus.params.abci.vote_extensions_enable_height",
			Value: "2",
		},
		{
			Key:   "consensus.params.block.max_gas",
			Value: "1000000000",
		},
		{
			Key: "app_state.feemarket.params",
			Value: feemarkettypes.Params{
				Alpha:                  feemarkettypes.DefaultAlpha,
				Beta:                   feemarkettypes.DefaultBeta,
				Theta:                  feemarkettypes.DefaultTheta,
				Delta:                  feemarkettypes.DefaultDelta,
				MinBaseGasPrice:        minGasPrice,
				MinLearningRate:        feemarkettypes.DefaultMinLearningRate,
				MaxLearningRate:        feemarkettypes.DefaultMaxLearningRate,
				TargetBlockUtilization: feemarkettypes.DefaultTargetBlockUtilization / 4,
				MaxBlockUtilization:    feemarkettypes.DefaultMaxBlockUtilization,
				Window:                 feemarkettypes.DefaultWindow,
				FeeDenom:               feemarkettypes.DefaultFeeDenom,
				// DISABLE
				Enabled:        false,
				DistributeFees: false,
			},
		},
		{
			Key: "app_state.feemarket.state",
			Value: feemarkettypes.State{
				BaseGasPrice: gasPrice,
				LearningRate: feemarkettypes.DefaultMaxLearningRate,
				Window:       make([]uint64, feemarkettypes.DefaultWindow),
				Index:        0,
			},
		},
	}

	denom = "untrn"
	spec  = &interchaintest.ChainSpec{
		ChainName:     "slinky",
		Name:          "slinky",
		NumValidators: &numValidators,
		NumFullNodes:  &numFullNodes,
		Version:       "latest",
		NoHostMount:   &noHostMount,
		ChainConfig: ibc.ChainConfig{
			EncodingConfig: &encodingConfig,
			Images: []ibc.DockerImage{
				image,
			},
			Type:           "cosmos",
			Name:           "slinky",
			Denom:          denom,
			ChainID:        "chain-id-0",
			Bin:            "neutrond",
			Bech32Prefix:   "neutron",
			CoinType:       "118",
			GasAdjustment:  gasAdjustment,
			GasPrices:      fmt.Sprintf("0%s", denom),
			TrustingPeriod: "48h",
			NoHostMount:    noHostMount,
			ModifyGenesis:  cosmos.ModifyGenesis(defaultGenesisKV),
			SkipGenTx:      true,
		},
	}
)

func TestSlinkyOracleIntegration(t *testing.T) {
	baseSuite := integration.NewSlinkyIntegrationSuite(
		spec,
		oracleImage,
		integration.WithInterchainConstructor(integration.CCVInterchainConstructor),
		integration.WithChainConstructor(integration.CCVChainConstructor),
		integration.WithDenom(denom),
	)

	suite.Run(t, integration.NewSlinkyOracleIntegrationSuite(baseSuite))
}

func TestSlinkyCCVIntegration(t *testing.T) {
	s := integration.NewSlinkyCCVIntegrationSuite(
		spec,
		oracleImage,
		integration.WithInterchainConstructor(integration.CCVInterchainConstructor),
		integration.WithChainConstructor(integration.CCVChainConstructor),
		integration.WithDenom(denom),
	)

	suite.Run(t, s)
}
