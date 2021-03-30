def abbreviate(text: str):
    result = ""
    words = text.replace("_", " ").replace("-", " ").split(" ")
    for word in words:
        for char in word:
            if not char.isalpha:
                continue
            result += char.upper()
            break
    return result