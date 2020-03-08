from textwrap import wrap

def proteins(strand):
    result = []
    proteins_dict = {
            "AUG" : "Methionine",
            "UUU" : "Phenylalanine",
            "UUC" : "Phenylalanine",
            "UUA" : "Leucine",
            "UUG" : "Leucine",
            "UCU" : "Serine",
            "UCC" : "Serine",
            "UCA" : "Serine",
            "UCG" : "Serine",
            "UAU" : "Tyrosine",
            "UAC" : "Tyrosine",
            "UGU" : "Cysteine",
            "UGC" : "Cysteine",
            "UGG" : "Tryptophan",
            }

    codes = wrap(strand, 3)
    for code in codes:
        if code in proteins_dict:
            result.append(proteins_dict[code])
        else:
            break
    return result
