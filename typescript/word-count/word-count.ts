
export default class Words {
    public count(sentence: string): Map<string | number, number> {
        const counts: Map<string | number, number> = new Map()
        this.removeFormatting(sentence.toLowerCase()).split(" ").map((word) => word.trim()).filter((word) => word !== "").forEach((word) => {
            const currentCountForWord = counts.get(word)
            const newCountForWord = currentCountForWord ? currentCountForWord + 1 : 1
            counts.set(word, newCountForWord)
        })
        return counts
    }

    private removeFormatting(sentence: string): string {
        return this.replaceNewlinesWithSpaces(this.replaceTabsWithSpaces(sentence))
    }
    private replaceNewlinesWithSpaces(sentence: string): string {
        return sentence.replace("\n", " ")
    }

    private replaceTabsWithSpaces(sentence: string): string {
        return sentence.replace("\t", " ")
    }

}
