import random

QUOTES = [
    "Simple is better than complex.",
    "Readability counts.",
    "Errors should never pass silently.",
    "Now is better than never.",
]


def random_quote() -> str:
    return random.choice(QUOTES)


if __name__ == "__main__":
    print(random_quote())
