
export default class Words {
    public count(sentence: string): Map<string | number, number> {
        return this.getWords(sentence).reduce((counts: Map<string | number, number>, currentWord: string) => {
            const currentCountForWord = counts.get(currentWord)
            const newCountForWord = currentCountForWord ? currentCountForWord + 1 : 1
            counts.set(currentWord, newCountForWord)
            return counts
        }, new Map<string | number, number>())
    }

    private getWords(sentence: string): string[] {
        return this.removeMultipleSpaces(this.removeFormatting(sentence).split(" "))
    }

    private removeMultipleSpaces(words: string[]): string[] {
        return words.map((word) => word.trim()).filter((word) => word !== "")
    }

    private removeFormatting(sentence: string): string {
        return sentence.toLowerCase().replace("\n", " ").replace("\t", " ")
    }
}
