
export default class Words {
    public count(sentence: string): Map<string | number, number> {
        return this.removeFormatting(sentence).split(" ").map((word) => word.trim()).filter((word) => word !== "").reduce((counts: Map<string | number, number>, word: string) => {
            const currentCountForWord = counts.get(word)
            const newCountForWord = currentCountForWord ? currentCountForWord + 1 : 1
            counts.set(word, newCountForWord)
            return counts
        }, new Map<string | number, number>())
    }

    private removeFormatting(sentence: string): string {
        return sentence.toLowerCase().replace("\n", " ").replace("\t", " ")
    }
}
