scores = {}
letter_scores = [
    (['a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't'], 1),
    (['d', 'g'], 2),
    (['b', 'c', 'm', 'p'], 3),
    (['f', 'h', 'v', 'w', 'y'], 4),
    (['k'], 5),
    (['j', 'x'], 8),
    (['q', 'z'], 10)
]

for ls in letter_scores:
    for l in ls[0]:
        scores[l] = ls[1]

def score(word):
    sum = 0
    for c in word.lower():
        sum += scores[c]
    return sum
