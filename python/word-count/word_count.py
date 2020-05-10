import re

def count_words(sentence):
    words = re.findall(r'[0-9]+|[a-z]+\'?[a-z]+|[a-z]+', sentence.lower())
    counts = {}
    for word in words:
        if word in counts:
            counts[word] += 1
        else:
            counts[word] = 1
    return counts
