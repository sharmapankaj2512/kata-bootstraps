class Thing:
    MAPPING = {
        'I': 1,
        'V': 5,
        'X': 10
    }

    def __init__(self, roman):
        self.roman = roman

    def to_arabic(self):
        return self.MAPPING.get(self.roman)
        # if self.roman == "V":
        #     return 5
        # elif self.roman == "X":
        #     return 10
