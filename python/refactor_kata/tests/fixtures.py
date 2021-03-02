from ancient_market.accounting import compute_market_fee, Merchant, Goods, TradeDetails


evil_merchant = Merchant(
    does_it_has_assasins=True,
    did_he_burn_our_market_in_past=True,
    did_he_wear_scalps_of_his_victims=True,
    did_his_goods_ever_kill_anyone=True,
    scale_of_usual_bribe=3,
    did_he_offer_free_drinks_to_clerks=False,
    did_he_trade_on_competittors_marketplace=True,
    how_many_seasons_do_we_work_together=10
)


neutral_merchant = Merchant(
    does_it_has_assasins=True,
    did_he_burn_our_market_in_past=False,
    did_he_wear_scalps_of_his_victims=False,
    did_his_goods_ever_kill_anyone=True,
    scale_of_usual_bribe=5,
    did_he_offer_free_drinks_to_clerks=True,
    did_he_trade_on_competittors_marketplace=False,
    how_many_seasons_do_we_work_together=3
)


good_merchant = Merchant(
    does_it_has_assasins=True,
    did_he_burn_our_market_in_past=False,
    did_he_wear_scalps_of_his_victims=False,
    did_his_goods_ever_kill_anyone=True,
    scale_of_usual_bribe=5,
    did_he_offer_free_drinks_to_clerks=True,
    did_he_trade_on_competittors_marketplace=False,
    how_many_seasons_do_we_work_together=3
)

GroatGrain_100 = Goods('Groat_grain', 100, 'tod')
Furs_10 = Goods('Fur', 10, 'unit')
Cypraea_10 = Goods('Cypraea', 10, 'unit')
Cypraea_100500 = Goods('Cypraea', 100500, 'unit')
BearTooth_3 = Goods('bear_tooth', 3, 'unit')
RaiStone_Pood_123 = Goods('Rai_stones', 123, 'pood')
GroatGrain_Tod_10 = Goods('Groat_grain', 10, 'tod')
Slave_1 = Goods('slave', 1, 'unit')
