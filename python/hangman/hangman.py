# Game status categories
# Change the values as you see fit
STATUS_WIN = "win"
STATUS_LOSE = "lose"
STATUS_ONGOING = "ongoing"


class Hangman:
    def __init__(self, word):
        self.remaining_guesses = 9
        self.status = STATUS_ONGOING
        self.word = word
        self.word_letters = set(word)

    def guess(self, char):
        if self.status != STATUS_ONGOING:
            raise ValueError("Game over!")

        if char in self.word_letters:
            self.word_letters.remove(char)
        else:
            self.remaining_guesses -= 1
        
        if not self.word_letters:
            self.status = STATUS_WIN
        
        if self.remaining_guesses < 0:
            self.status = STATUS_LOSE

    def get_masked_word(self):
        return "".join("_" if c in self.word_letters else c for c in self.word)

    def get_status(self):
        return self.status
