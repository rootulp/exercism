class SimpleCipher {

    static isValidKey(_inputKey: string): boolean {
        return true
    }

    static generateRandomKey(): string {
        return "foo"
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
