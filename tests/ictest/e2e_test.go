package ictest_test

import (
	"encoding/json"
	"fmt"
	"github.com/icza/dyno"
	"log"
	"strconv"
	"strings"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/suite"

	"github.com/neutron-org/neutron/v4/app"
	"github.com/neutron-org/neutron/v4/tests/ictest"
)

var (
	minBaseFee = sdkmath.LegacyNewDec(10)
	baseFee    = sdkmath.LegacyNewDec(1000000)

	// config params
	numValidators = 3
	numFullNodes  = 1
	denom         = "stake"

	image = ibc.DockerImage{
		Repository: "neutron-e2e",
		Version:    "latest",
		UidGid:     "1000:1000",
	}
	encodingConfig = &testutil.TestEncodingConfig{
		InterfaceRegistry: app.MakeEncodingConfig().InterfaceRegistry,
		Codec:             app.MakeEncodingConfig().Marshaler,
		TxConfig:          app.MakeEncodingConfig().TxConfig,
		Amino:             app.MakeEncodingConfig().Amino,
	}
	noHostMount   = false
	gasAdjustment = 10.0

	genesisKV = []cosmos.GenesisKV{
		{
			Key: "app_state.feemarket.params",
			Value: feemarkettypes.Params{
				Alpha:                  feemarkettypes.DefaultAlpha,
				Beta:                   feemarkettypes.DefaultBeta,
				Theta:                  feemarkettypes.DefaultTheta,
				Delta:                  feemarkettypes.DefaultDelta,
				MinBaseFee:             minBaseFee,
				MinLearningRate:        feemarkettypes.DefaultMinLearningRate,
				MaxLearningRate:        feemarkettypes.DefaultMaxLearningRate,
				TargetBlockUtilization: feemarkettypes.DefaultTargetBlockUtilization / 4,
				MaxBlockUtilization:    feemarkettypes.DefaultMaxBlockUtilization,
				Window:                 feemarkettypes.DefaultWindow,
				FeeDenom:               feemarkettypes.DefaultFeeDenom,
				Enabled:                true,
				DistributeFees:         false,
			},
		},
		{
			Key: "app_state.feemarket.state",
			Value: feemarkettypes.State{
				BaseFee:      baseFee,
				LearningRate: feemarkettypes.DefaultMaxLearningRate,
				Window:       make([]uint64, feemarkettypes.DefaultWindow),
				Index:        0,
			},
		},
	}

	// interchain specification
	spec = &interchaintest.ChainSpec{
		ChainName:     "neutron",
		Name:          "neutron",
		NumValidators: &numValidators,
		NumFullNodes:  &numFullNodes,
		Version:       "latest",
		NoHostMount:   &noHostMount,
		ChainConfig: ibc.ChainConfig{
			EncodingConfig: encodingConfig,
			Images: []ibc.DockerImage{
				image,
			},
			Type:           "cosmos",
			Name:           "neutron",
			Denom:          denom,
			ChainID:        "chain-id-0",
			Bin:            "neutrond",
			Bech32Prefix:   "neutron",
			SkipGenTx:      false,
			CoinType:       "118",
			GasAdjustment:  gasAdjustment,
			GasPrices:      fmt.Sprintf("0%s", denom),
			TrustingPeriod: "48h",
			NoHostMount:    noHostMount,
			ModifyGenesis:  ModifyGenesis(genesisKV),
		},
	}
)

func TestE2ETestSuite(t *testing.T) {
	s := ictest.NewE2ETestSuiteFromSpec(spec)
	suite.Run(t, s)
}

func ModifyGenesis(genesisKV []cosmos.GenesisKV) func(ibc.ChainConfig, []byte) ([]byte, error) {
	return func(chainConfig ibc.ChainConfig, genbz []byte) ([]byte, error) {
		g := make(map[string]interface{})
		if err := json.Unmarshal(genbz, &g); err != nil {
			return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
		}

		for idx, values := range genesisKV {
			splitPath := strings.Split(values.Key, ".")

			path := make([]interface{}, len(splitPath))
			for i, component := range splitPath {
				if v, err := strconv.Atoi(component); err == nil {
					path[i] = v
				} else {
					path[i] = component
				}
			}

			if err := dyno.Set(g, values.Value, path...); err != nil {
				return nil, fmt.Errorf("failed to set value (index:%d) in genesis json: %w", idx, err)
			}
		}

		log.Fatal(g)

		out, err := json.Marshal(g)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal genesis bytes to json: %w", err)
		}
		return out, nil
	}
}
