def saddle_points(matrix):
    # make sure every row have same length
    if len(set(map(len, matrix))) > 1:
        raise ValueError("invalid input matrix")

    results = []
    for i, row in enumerate(matrix):
        for j, number in enumerate(row):
            if number == max(row) and number == min([row[j] for row in matrix]):
                results.append({"row": i + 1, "column": j + 1})

    return results
