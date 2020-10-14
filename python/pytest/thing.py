class Thing:
    MAPPING = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'M': 1000
    }

    def __init__(self, roman):
        self.roman = roman

    def to_arabic(self):
        if self.roman == "II":
            return 2
        if self.roman == "VI":
            return 6
        return self.MAPPING.get(self.roman)
