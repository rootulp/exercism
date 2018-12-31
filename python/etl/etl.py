def transform(old):
    return {value.lower(): score for score, values in
            old.items() for value in values}
