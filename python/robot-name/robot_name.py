import random
import string

letters = string.ascii_uppercase
digits = string.digits

class Robot:
    def __init__(self):
        self.reset()
    def reset(self):
        random.seed()
        chars = "".join(random.sample(letters, 2))
        numbers = "".join(random.sample(digits, 3))
        self.name = chars + numbers

