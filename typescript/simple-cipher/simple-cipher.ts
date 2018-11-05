class SimpleCipher {
    public key: string

    constructor(inputKey ?: string) {
        if (inputKey === undefined) {
            this.key = Key.generate()
        } else if (Key.isValid(inputKey)) {
            this.key = inputKey
        } else {
            throw "Bad key"
        }
    }
    public encode(data: string): string {
        return data
    }

    public decode(encodedData: string): string {
        return encodedData
    }

}

class Key {
    static readonly alphabet = "abcdefghijklmnopqrstuvwxyz"

    static isValid(inputKey: string): boolean {
        return Key.containsOnlyLowercaseCharacters(inputKey) && Key.isNotEmpty(inputKey)
    }

    static generate(): string {
        const KEY_LENGTH: number = 100
        return Key.generateRandomString(KEY_LENGTH)
    }

    static generateRandomString(length: number): string {
        return [...Array(length)].map(this.generateRandomCharacter).join("")
    }

    static generateRandomCharacter(): string  {
        return Key.alphabet.charAt(Math.floor(Math.random() * Key.alphabet.length))
    }

    static containsOnlyLowercaseCharacters(inputKey: string): boolean {
        return inputKey.split("").every(Key.isLowercaseLetter)
    }

    static isLowercaseLetter(character: string): boolean {
        return Key.alphabet.includes(character)
    }

    static isNotEmpty(inputKey: string): boolean {
        return inputKey !== ""
    }

}

export default SimpleCipher
