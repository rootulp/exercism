def sieve(limit):
    primes = []
    unmarked = list(range(2, limit + 1))
    while len(unmarked) > 0:
        prime = unmarked.pop(0)
        for multiplier in range(2, limit // prime + 1):
            if prime * multiplier in unmarked:
                unmarked.remove(prime * multiplier)
        primes.append(prime)
    return primes
