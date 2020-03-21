def slices(series, length):
    if length > len(series) or length < 1:
        raise ValueError("bad input")

    result = []
    for start in range(0, len(series)-length+1):
        result.append(series[start:start+length])
    return result
