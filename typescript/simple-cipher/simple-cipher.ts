class SimpleCipher {
    public key: string
    static readonly alphabet = "abcdefghijklmnopqrstuvwxyz"

    constructor(inputKey ?: string) {
        if (inputKey === undefined) {
            this.key = Key.generate()
        } else if (Key.isValid(inputKey)) {
            this.key = inputKey
        } else {
            throw "Bad key"
        }
    }
    public encode = (data: string): string => {
        return data.split("").map(this.encodeCharacter).join("")
    }

    public decode = (encodedData: string): string => {
        return encodedData.split("").map(this.decodeCharacter).join("")
    }

    private encodeCharacter = (character: string, index: number): string => {
        const dataValue = this.valueOf(character)
        const shiftDistance = this.getShiftDistance(index)
        const encryptedValue = dataValue + shiftDistance
        const encodedCharacter = SimpleCipher.alphabet[encryptedValue % SimpleCipher.alphabet.length]
        return encodedCharacter
    }

    private decodeCharacter = (character: string, index: number): string => {
        const dataValue = this.valueOf(character)
        const shiftDistance = this.getShiftDistance(index)
        const decryptedValue = dataValue - shiftDistance
        const decodedCharacter = SimpleCipher.alphabet[decryptedValue % SimpleCipher.alphabet.length]
        return decodedCharacter
    }

    private getShiftDistance = (index: number): number => {
        const encryptionCharacter = this.getEncryptionCharacter(index)
        return this.valueOf(encryptionCharacter)
    }

    private getEncryptionCharacter = (index: number): string => {
        return this.key[index % this.key.length]
    }

    private valueOf = (character: string): number => {
        return SimpleCipher.alphabet.indexOf(character)

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
