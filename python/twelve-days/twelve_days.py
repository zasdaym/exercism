poems = [
    ("first", "a Partridge in a Pear Tree."),
    ("second", "two Turtle Doves, "),
    ("third", "three French Hens, "),
    ("fourth", "four Calling Birds, "),
    ("fifth", "five Gold Rings, "),
    ("sixth", "six Geese-a-Laying, "),
    ("seventh", "seven Swans-a-Swimming, "),
    ("eighth", "eight Maids-a-Milking, "),
    ("ninth", "nine Ladies Dancing, "),
    ("tenth", "ten Lords-a-Leaping, "),
    ("eleventh", "eleven Pipers Piping, "),
    ("twelfth", "twelve Drummers Drumming, ")
]

def recite(start_verse, end_verse):
    return [get_verse(num) for num in range(start_verse, end_verse + 1)]

def get_verse(num):
    verses = []
    if num == 1:
        verses = poems[0][1]
    else:
        verses = [poem[1] for poem in poems[num-1:0:-1]]
        verses.append(f"and {poems[0][1]}")

    opening = f"On the {poems[num - 1][0]} day of Christmas my true love gave to me: "
    return opening + ''.join(verses)

print(recite(12, 12))
