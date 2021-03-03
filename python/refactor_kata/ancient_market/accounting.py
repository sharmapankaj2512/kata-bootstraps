def compute_fair_price(src_amount, src_currency_name, src_currency_unit,
                       dst_currency_name, dst_currency_unit):
    exchange_rates = [
        # SRC               DST
        ('Cypraea', 'unit', 'Groat_grain', 'tod', 0.9),
        ('Groat_grain', 'tod', 'Cypraea', 'unit', 111),
        ('Salt', 'pood', 'Groat_grain', 'tod', 3.4),
        ('Horse', 'unit', 'Cypraea', 'unit', 100000),
        ('bear_tooth', 'unit', 'Groat_grain', 'tod', 1.0),
        ('Fur', 'UNIT', 'Cypraea', 'unit', 500),
        ('Rai_stones', 'tod', 'Groat_grain', 'tod', 0.025),
        ('slave', 'unit', 'Cypraea', 'unit', 150000),
    ]

    weights = {
        'pood': {
            'tod': 0.775
        },
        'tod': {
            'pood': 1.289
        }
    }

    return (
            src_amount *
            [x for x in exchange_rates if x[0] == src_currency_name and x[2] == dst_currency_name]
            [0][-1]
    )


def compute_market_fee(deal_details):
    """
        Given deal details determine appropriate size of fee to charge from merchant's profit
        take into account merchant characteristics
    :param deal_details:
    :return:
    """
    fair_price = compute_fair_price(
        deal_details.src.amount, deal_details.src.product_name, deal_details.src.product_unit,
        deal_details.dst.product_name, deal_details.dst.product_unit
    )

    merchants_profit = abs(deal_details.dst.amount - fair_price)

    merchant_fee_factor = calculate_merchant_fee(deal_details)

    print(merchant_fee_factor)

    if merchant_fee_factor <= 0:
        return 0

    return merchants_profit * merchant_fee_factor


def calculate_merchant_fee(deal_details):
    merchant_fee_factor = 1.0
    if deal_details.merchant.did_he_burn_our_market_in_past and deal_details.merchant.did_his_goods_ever_kill_anyone:
        merchant_fee_factor = 0.25
    elif (
            not deal_details.merchant.did_he_burn_our_market_in_past and not deal_details.merchant.did_his_goods_ever_kill_anyone
            and (
                    deal_details.merchant.does_it_has_assasins or
                    deal_details.merchant.did_his_goods_ever_kill_anyone
            )
    ):
        merchant_fee_factor = 0.75
    elif (
            deal_details.merchant.did_he_burn_our_market_in_past or
            deal_details.merchant.did_his_goods_ever_kill_anyone or
            deal_details.merchant.does_it_has_assasins or
            deal_details.merchant.did_his_goods_ever_kill_anyone
    ):
        merchant_fee_factor = 0.75
    if deal_details.merchant.did_he_offer_free_drinks_to_clerks:
        merchant_fee_factor -= 0.05
    if deal_details.merchant.did_he_trade_on_competittors_marketplace:
        merchant_fee_factor += 0.1
    if 0 < deal_details.merchant.scale_of_usual_bribe < 10:
        merchant_fee_factor -= 0.05 * deal_details.merchant.scale_of_usual_bribe
    else:
        raise Exception("Check your papers!!!")
    if deal_details.merchant.how_many_seasons_do_we_work_together > 0:
        merchant_fee_factor -= 0.01 * deal_details.merchant.how_many_seasons_do_we_work_together
    return merchant_fee_factor
