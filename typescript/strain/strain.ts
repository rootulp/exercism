export function keep<T>(collection: T[], predicate: (element: T)  => boolean ): T[] {
    // no op
    const kept: T[] = []
    collection.forEach((element) => {
        if (predicate(element)) {
            kept.push(element)
        }
    })
    return kept
}

export function discard<T>(collection: T[], predicate: (element: T)  => boolean ): T[] {
    // no op
    const discarded: T[] = []
    collection.forEach((element) => {
        if (!predicate(element)) {
            discarded.push(element)
        }
    })
    return discarded

}
