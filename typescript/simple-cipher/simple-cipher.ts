class SimpleCipher {

    static isValidKey(_inputKey: string): boolean {
        return true
    }

    static generateRandomKey(): string {
        const KEY_LENGTH: number = 100
        return SimpleCipher.randomString(KEY_LENGTH)
    }

    static randomString(length: number): string {
        return [...Array(length)].map(this.randomCharacter).join("")
    }

    static randomCharacter(): string  {
        const alphabet = "abcdefghijklmnopqrstuvwxyz"
        return alphabet.charAt(Math.floor(Math.random() * alphabet.length))
    }

    public key: string

    constructor(inputKey ?: string) {
        if (inputKey === undefined) {
            this.key = SimpleCipher.generateRandomKey()
        } else if (SimpleCipher.isValidKey(inputKey)) {
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

export default SimpleCipher
