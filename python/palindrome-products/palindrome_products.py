def largest(min_factor, max_factor):
    return palindrome_product(min_factor, max_factor)

def smallest(min_factor, max_factor):
    return palindrome_product(min_factor, max_factor, largest=False)

def palindrome_product(min_factor, max_factor, largest=True):
    if min_factor > max_factor:
        raise ValueError("min_factor must be lower than max_factor")

    args = (min_factor ** 2, max_factor ** 2 + 1) if largest is False else (max_factor ** 2, min_factor ** 2 - 1, -1)

    for i in range(*args):
        nums = str(i)
        if nums == nums[::-1] and any(min_factor <= i // j <= max_factor for j in range(min_factor, max_factor + 1) if i % j == 0):
            return i, ((x, i // x) for x in range(min_factor, max_factor + 1) if i % x == 0 and min_factor <= x <= i // x <= max_factor)

    return None, []
