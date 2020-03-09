def convert(number):
    factors = {
            3 : "Pling",
            5 : "Plang",
            7 : "Plong",
            }

    result = ""

    for factor in factors:
        if number % factor == 0:
            result += factors[factor]

    if result == "":
        result = str(number)

    return result
