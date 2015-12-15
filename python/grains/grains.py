def on_square(x):
    return 2 ** (x - 1)


def total_after(x):
    return 1 if x == 1 else on_square(x) + total_after(x - 1)
