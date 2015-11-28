from datetime import timedelta


def add_gigasecond(date):
    return date + timedelta(seconds=10**9)
