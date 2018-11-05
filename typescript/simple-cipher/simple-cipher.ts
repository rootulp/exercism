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
        const shiftDistance = this.getShiftDistance(index)
        const encryptedValue = this.valueOf(character) + shiftDistance
        return this.convertValueToCharacter(encryptedValue)
    }

    private decodeCharacter = (character: string, index: number): string => {
        const shiftDistance = this.getShiftDistance(index)
        const decryptedValue = this.valueOf(character) - shiftDistance
        return this.convertValueToCharacter(decryptedValue)
    }

    private getShiftDistance = (index: number): number => {
        const encryptionCharacter = this.getEncryptionCharacter(index)
        return this.valueOf(encryptionCharacter)
    }

    private getEncryptionCharacter = (index: number): string => {
        return this.key[index % this.key.length]
    }

    private convertValueToCharacter = (value: number): string => {
        return SimpleCipher.alphabet[mod(value,  SimpleCipher.alphabet.length)]
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

// % in JavaScript is the remainder operator (rather than modulo)
// See: https://stackoverflow.com/questions/4467539/javascript-modulo-gives-a-negative-result-for-negative-numbers
function mod(n: number, m: number) {
    return ((n % m) + m) % m
}

export default SimpleCipher
