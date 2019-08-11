
export default class Words {
    public count(sentence: string): Map<string | number, number> {
        const counts: Map<string | number, number> = new Map()
        sentence.split(" ").forEach((word) => {
            const currentCountForWord = counts.get(word)
            const newCountForWord = currentCountForWord ? currentCountForWord + 1 : 1
            counts.set(word, newCountForWord)
        })
        return counts
    }

}
