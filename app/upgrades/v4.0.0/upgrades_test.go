package v400_test

import (
	"testing"

	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"

	slinkyutils "github.com/neutron-org/neutron/v4/utils/slinky"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	comettypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/suite"

	_ "embed"

	v400 "github.com/neutron-org/neutron/v4/app/upgrades/v4.0.0"
	"github.com/neutron-org/neutron/v4/testutil"
)

//go:embed markets.json
var marketsJSON []byte

type UpgradeTestSuite struct {
	testutil.IBCConnectionTestSuite
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.IBCConnectionTestSuite.SetupTest()
}

func (suite *UpgradeTestSuite) TestOracleUpgrade() {
	app := suite.GetNeutronZoneApp(suite.ChainA)
	ctx := suite.ChainA.GetContext()
	t := suite.T()

	markets, err := slinkyutils.ReadMarketsFromFile(marketsJSON)
	suite.Require().NoError(err)
	marketMap := slinkyutils.ToMarketMap(markets)

	upgrade := upgradetypes.Plan{
		Name:   v400.UpgradeName,
		Info:   "some text here",
		Height: 100,
	}
	require.NoError(t, app.UpgradeKeeper.ApplyUpgrade(ctx, upgrade))

	params, err := app.MarketMapKeeper.GetParams(ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(params.MarketAuthorities[0], "neutron1hxskfdxpp5hqgtjj6am6nkjefhfzj359x0ar3z")
	suite.Require().Equal(params.Admin, "neutron1hxskfdxpp5hqgtjj6am6nkjefhfzj359x0ar3z")

	// check that the market map was properly set
	mm, err := app.MarketMapKeeper.GetAllMarkets(ctx)
	gotMM := marketmaptypes.MarketMap{Markets: mm}
	suite.Require().NoError(err)
	suite.Require().True(marketMap.Equal(gotMM))

	numCps, err := app.OracleKeeper.GetNumCurrencyPairs(ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(numCps, uint64(len(markets)))

	// check that all currency pairs have been initialized in the oracle module
	for _, market := range markets {
		decimals, err := app.OracleKeeper.GetDecimalsForCurrencyPair(ctx, market.Ticker.CurrencyPair)
		suite.Require().NoError(err)
		suite.Require().Equal(market.Ticker.Decimals, decimals)

		price, err := app.OracleKeeper.GetPriceWithNonceForCurrencyPair(ctx, market.Ticker.CurrencyPair)
		suite.Require().NoError(err)
		suite.Require().Equal(uint64(0), price.Nonce())     // no nonce because no updates yet
		suite.Require().Equal(uint64(0), price.BlockHeight) // no block height because no price written yet

		suite.Require().True(market.Ticker.Enabled)
	}
}

func (suite *UpgradeTestSuite) TestEnableVoteExtensionsUpgrade() {
	app := suite.GetNeutronZoneApp(suite.ChainA)
	ctx := suite.ChainA.GetContext()
	t := suite.T()

	oldParams, err := app.ConsensusParamsKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)

	// VoteExtensionsEnableHeight must be updated after the upgrade on upgrade height
	// but the rest of params must be the same
	oldParams.Params.Abci = &comettypes.ABCIParams{VoteExtensionsEnableHeight: ctx.BlockHeight() + 4}
	// it is automatically tracked in upgrade handler, we need to set it manually for tests
	oldParams.Params.Version = &comettypes.VersionParams{App: 0}

	upgrade := upgradetypes.Plan{
		Name:   v400.UpgradeName,
		Info:   "some text here",
		Height: ctx.BlockHeight(),
	}
	require.NoError(t, app.UpgradeKeeper.ApplyUpgrade(ctx, upgrade))

	newParams, err := app.ConsensusParamsKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)

	suite.Require().Equal(oldParams, newParams)
}
