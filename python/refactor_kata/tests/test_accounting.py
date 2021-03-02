import pytest

from fixtures import (
    Furs_10, Cypraea_100500, BearTooth_3, RaiStone_Pood_123,
    GroatGrain_Tod_10, GroatGrain_100, Cypraea_10, Slave_1,
    evil_merchant, good_merchant, neutral_merchant
)
from ancient_market.accounting import compute_market_fee, TradeDetails


def test_direct_trade_market_fee_evil_merchant():
    trade = TradeDetails(GroatGrain_100, Cypraea_100500, evil_merchant)
    assert pytest.approx(compute_market_fee(trade), 0.01) == 8939.99


def test_direct_trade_market_fee_good_merchant():
    trade = TradeDetails(GroatGrain_100, Cypraea_100500, good_merchant)
    assert pytest.approx(compute_market_fee(trade), 0.01) == 37547.99


def test_direct_trade_market_fee_neutral_merchant():
    trade = TradeDetails(GroatGrain_100, Cypraea_100500, neutral_merchant)
    assert pytest.approx(compute_market_fee(trade), 0.01) == 37547.99


def test_direct_trade_furs_neutral_merchant():
    trade = TradeDetails(Furs_10, Cypraea_10, neutral_merchant)
    assert pytest.approx(compute_market_fee(trade), 0.01) == 2095.79


def test_direct_trade_slave_neutral_merchant():
    trade = TradeDetails(Slave_1, Cypraea_100500, neutral_merchant)
    assert pytest.approx(compute_market_fee(trade), 0.01) == 20789.99


@pytest.mark.skip(reason="Non-direct currencies conversions is not supported yet")
def test_non_direct_trade_market_fee_neutral_merchant():
    trade = TradeDetails(BearTooth_3, Furs_10, neutral_merchant)

    """
    Furs_100 = Goods('Fur', 100, 'unit') => 50000 Cypraea
    BearTooth_3 = Goods('bear_tooth', 3, 'unit') => 3.0 tod Groat_grain = 333 Cypraea
    Groat_grain 1.0 tod = 111 Cypraea

    I.e. fair price for 3 bear tooth - is 333 cypraea

    taxable amount: 50000-333 = 49667 
    neutral merchant fee: 0.41999999999999993

    """

    assert pytest.approx(compute_market_fee(trade), 0.01) == 20860.139


@pytest.mark.skip(reason="conversions of measurement units not supported yet")
def test_non_same_unit_trade_market_fee_good_merchant():
    trade = TradeDetails(RaiStone_Pood_123, GroatGrain_Tod_10, good_merchant)
    """
    RaiStone_Pood_123 = 123 * 1.289 tod
    tod = 1.289 pood
    GroatGrain_Tod_10
    
    I.e. fair price for 123 poods of RaiStone - is 158.547 tods of GroatGrain
    
    taxable amount: 158.547-10 = 148.547 
    good merchant fee: 0.41999999999999993
    
    """

    assert pytest.approx(compute_market_fee(trade), 0.01) == 62.389
