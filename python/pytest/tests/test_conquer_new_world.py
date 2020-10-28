from app.conquer_new_world import conquer_new_world

def test_planet_success():
    conquer_new_world("")
    assert 1 == 1

def test_planet_failed():
    assert 1 == 2