def is_armstrong_number(number):
    numbers = list(map(int, str(number)))
    length = len(numbers)
    sum = 0
    for i in numbers:
        sum += i ** length

    return sum == number
