import random
import string


def random_password(length: int = 12) -> str:
    alphabet = string.ascii_letters + string.digits
    return "".join(random.choice(alphabet) for _ in range(length))


if __name__ == "__main__":
    print(random_password())
