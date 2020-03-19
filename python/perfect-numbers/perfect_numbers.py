def classify(number):
    if number < 1:
        raise ValueError("Number must be equal or greater than 1")
    if number == 1:
        return "deficient"
    result = 0
    div = 2
    factors = {1}
    while number != div:
        if number % div == 0:
            factors.add(div)
        div += 1
   
    result = sum(factors)
    if result == number:
        return "perfect"
    elif result > number:
        return "abundant"
    else:
        return "deficient"
