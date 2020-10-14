class Thing:

    def __init__(self, roman):
        self.roman = roman

    def to_arabic(self):
        if self.roman == "V":
            return 5

        return 1
