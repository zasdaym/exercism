def is_valid(isbn):
    splitted = list("".join(isbn.split("-")))
    if len(splitted) != 10 or not "".join(splitted[:-1]).isdigit():
        return False

    if not splitted[-1].isdigit() and splitted[-1] != "X":
        return False

    sum = 0
    for i in range(9):
        sum += int(splitted[i]) * (10-i)

    if splitted[-1] == "X":
        sum += 10
    else:
        sum += int(splitted[-1])

    return sum % 11 == 0
