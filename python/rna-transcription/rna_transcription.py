def to_rna(dna_strand):
    codes = {
            "C" : "G",
            "G" : "C",
            "T" : "A",
            "A" : "U",
            }
    
    s = list(dna_strand)
    result = ""
    return result.join(list(map(lambda x: codes[x], s)))
