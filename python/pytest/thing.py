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
        result = 0
        for entry in self.roman:
            n = self.MAPPING.get(entry)
            result += n
        return result
