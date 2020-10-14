from thing import Thing


def test_correct_input_single_digits():
    thing = Thing("I")
    assert thing.to_arabic() == 1

    thing = Thing("V")
    assert thing.to_arabic() == 5

    thing = Thing("X")
    assert thing.to_arabic() == 10

    thing = Thing("L")
    assert thing.to_arabic() == 50

    thing = Thing("C")
    assert thing.to_arabic() == 100

    thing = Thing("M")
    assert thing.to_arabic() == 1000


# def test_fail():
#     thing = Thing("Albert")
#     assert "Wrong!" == thing.return_hello_name()
