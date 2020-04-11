def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("strands length not same!")

    return sum(strand_a[i] != strand_b[i] for i in range(len(strand_a)))
