package v400_test

import (
	"sort"
	"testing"

	"github.com/skip-mev/connect/v2/cmd/constants/marketmaps"
	marketmaptypes "github.com/skip-mev/connect/v2/x/marketmap/types"

	slinkytypes "github.com/skip-mev/connect/v2/pkg/types"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	comettypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/suite"

	v400 "github.com/neutron-org/neutron/v4/app/upgrades/v4.0.1"
	"github.com/neutron-org/neutron/v4/testutil"
)

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

	oldParams, err := app.ConsensusParamsKeeper.Params(ctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	// it is automatically tracked in upgrade handler, we need to set it manually for tests
	oldParams.Params.Version = &comettypes.VersionParams{App: 0}
	// we need to properly set consensus params for tests or we get a panic
	suite.Require().NoError(app.ConsensusParamsKeeper.ParamsStore.Set(ctx, *oldParams.Params))

	markets := marketmaps.CoreMarketMap.Markets
	suite.Require().NoError(err)

	upgrade := upgradetypes.Plan{
		Name:   v400.UpgradeName,
		Info:   "some text here",
		Height: 100,
	}
	require.NoError(t, app.UpgradeKeeper.ApplyUpgrade(ctx, upgrade))

	params, err := app.MarketMapKeeper.GetParams(ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(params.MarketAuthorities[0], "neutron1hxskfdxpp5hqgtjj6am6nkjefhfzj359x0ar3z")
	suite.Require().Equal(params.MarketAuthorities[1], "neutron1ua63s43u2p4v38pxhcxmps0tj2gudyw2hfeetz")
	suite.Require().Equal(params.Admin, "neutron1hxskfdxpp5hqgtjj6am6nkjefhfzj359x0ar3z")

	// check that the market map was properly set
	mm, err := app.MarketMapKeeper.GetAllMarkets(ctx)
	gotMM := marketmaptypes.MarketMap{Markets: mm}
	suite.Require().NoError(err)
	suite.Require().True(marketmaps.CoreMarketMap.Equal(gotMM))

	numCps, err := app.OracleKeeper.GetNumCurrencyPairs(ctx)
	suite.Require().NoError(err)
	suite.Require().Equal(numCps, uint64(len(markets)))

	// check that all currency pairs have been initialized in the oracle module
	tickers := make([]slinkytypes.CurrencyPair, 0, len(markets))
	for _, market := range markets {
		decimals, err := app.OracleKeeper.GetDecimalsForCurrencyPair(ctx, market.Ticker.CurrencyPair)
		suite.Require().NoError(err)
		suite.Require().Equal(market.Ticker.Decimals, decimals)

		price, err := app.OracleKeeper.GetPriceWithNonceForCurrencyPair(ctx, market.Ticker.CurrencyPair)
		suite.Require().NoError(err)
		suite.Require().Equal(uint64(0), price.Nonce())     // no nonce because no updates yet
		suite.Require().Equal(uint64(0), price.BlockHeight) // no block height because no price written yet

		suite.Require().True(market.Ticker.Enabled)

		tickers = append(tickers, market.Ticker.CurrencyPair)
	}

	// check IDs for inserted currency pairs, sort currency-pairs alphabetically
	sort.Slice(tickers, func(i, j int) bool {
		return tickers[i].String() < tickers[j].String()
	})

	for i, cp := range tickers {
		id, found := app.OracleKeeper.GetIDForCurrencyPair(ctx, cp)
		suite.Require().True(found)
		suite.Require().Equal(uint64(i), id)
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
	// we need to properly set consensus params for tests or we get a panic
	suite.Require().NoError(app.ConsensusParamsKeeper.ParamsStore.Set(ctx, *oldParams.Params))

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

func (suite *UpgradeTestSuite) TestDynamicFeesUpgrade() {
	app := suite.GetNeutronZoneApp(suite.ChainA)
	ctx := suite.ChainA.GetContext()
	t := suite.T()

	upgrade := upgradetypes.Plan{
		Name:   v400.UpgradeName,
		Info:   "some text here",
		Height: 100,
	}
	require.NoError(t, app.UpgradeKeeper.ApplyUpgrade(ctx, upgrade))

	params := app.DynamicFeesKeeper.GetParams(ctx)
	suite.Require().Equal(params.NtrnPrices, v400.NtrnPrices)
}
