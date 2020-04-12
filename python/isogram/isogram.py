def is_isogram(string):
    s = ''.join(x.lower() for x in string if x.isalpha())
    return len(s) == len(set(s))
