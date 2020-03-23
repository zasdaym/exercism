def is_valid(isbn):
    new_isbn = list("".join(isbn.split("-")))

    if len(new_isbn) != 10:
        return False

    if "X" == new_isbn[-1]:
        new_isbn[-1] = "10"

    if not "".join(new_isbn).isdigit():
        return False

    sum = 0
    multiplier = 10
    for i in new_isbn:
        sum += int(i) * multiplier
        multiplier -= 1 

    return sum % 11 == 0
