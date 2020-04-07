def find_anagrams(word, candidates):
    result = []
    for i in candidates:
        if sorted(word.lower()) == sorted(i.lower()) and word.lower() != i.lower():
            result.append(i)

    return result
