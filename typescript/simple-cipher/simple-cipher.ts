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
    static isValid(_inputKey: string): boolean {
        return true
    }

    static generate(): string {
        const KEY_LENGTH: number = 100
        return Key.generateRandomString(KEY_LENGTH)
    }

    static generateRandomString(length: number): string {
        return [...Array(length)].map(this.generateRandomCharacter).join("")
    }

    static generateRandomCharacter(): string  {
        const alphabet = "abcdefghijklmnopqrstuvwxyz"
        return alphabet.charAt(Math.floor(Math.random() * alphabet.length))
    }

}

export default SimpleCipher
