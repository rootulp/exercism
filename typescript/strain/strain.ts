export function keep<T>(collection: T[], predicate: (element: T)  => boolean ): T[] {
    const kept: T[] = []
    collection.forEach((element) => {
        if (predicate(element)) {
            kept.push(element)
        }
    })
    return kept
}

export function discard<T>(collection: T[], predicate: (element: T)  => boolean ): T[] {
    const discarded: T[] = []
    collection.forEach((element) => {
        if (!predicate(element)) {
            discarded.push(element)
        }
    })
    return discarded
}
