from thing import Thing


def test_correct_input_1():
    thing = Thing("I")

    assert thing.to_arabic() == 1


# def test_fail():
#     thing = Thing("Albert")
#     assert "Wrong!" == thing.return_hello_name()
