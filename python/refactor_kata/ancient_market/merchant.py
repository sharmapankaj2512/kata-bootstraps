class Merchant:
    def __init__(
            self,
            does_it_has_assasins,
            did_he_burn_our_market_in_past,
            did_he_wear_scalps_of_his_victims,
            did_his_goods_ever_kill_anyone,
            scale_of_usual_bribe,
            did_he_offer_free_drinks_to_clerks,
            did_he_trade_on_competittors_marketplace,
            how_many_seasons_do_we_work_together,
    ):
        # danger factor
        self.does_it_has_assasins = does_it_has_assasins
        self.did_he_burn_our_market_in_past = did_he_burn_our_market_in_past
        self.did_he_wear_scalps_of_his_victims = did_he_wear_scalps_of_his_victims
        self.did_his_goods_ever_kill_anyone = did_his_goods_ever_kill_anyone
        # attractiveness factor
        self.scale_of_usual_bribe = scale_of_usual_bribe
        self.did_he_offer_free_drinks_to_clerks = did_he_offer_free_drinks_to_clerks
        self.did_he_trade_on_competittors_marketplace = did_he_trade_on_competittors_marketplace
        self.how_many_seasons_do_we_work_together = how_many_seasons_do_we_work_together